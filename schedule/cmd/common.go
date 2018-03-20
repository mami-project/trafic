package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/mami-project/trafic/config"
	"github.com/mami-project/trafic/cruncher"
	"github.com/mami-project/trafic/runner"
	"github.com/spf13/viper"
)

type StatusReport struct {
	Label string
	Role  runner.Role
	Error error
}

type Runners []runner.Runner
type FlowConfigs []config.FlowConfig

type RunnerStats struct {
	out string
}

const (
	DefaultSchedFrequency = (100 * time.Millisecond)
)

// RunnersMap maps a flow label (which MUST be unique in a given traffic mix)
// to the associated iperf3 runnable instance on one side of the flow
type RunnersMap map[string]runner.Runner

var (
	R RunnersMap
	M sync.Mutex
)

func NewLogger(tag string) (*log.Logger, error) {
	return log.New(os.Stderr, tag, log.LstdFlags|log.LUTC|log.Lshortfile), nil
}

func loadFlows(dir string) (FlowConfigs, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("cannot read dir: %s: %v", dir, err)
		return nil, err
	}

	var flows FlowConfigs

	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		fqn := filepath.Join(dir, file.Name())

		flow, err := config.NewFlowConfigFromFile(fqn)
		if err != nil {
			log.Printf("cannot parse %s: %v", fqn, err)
			return nil, err
		}

		flows = append(flows, *flow)
	}

	return flows, nil
}

func sortRunnersByDeadline(runners Runners) {
	sort.Slice(runners, func(i, j int) bool {
		return runners[i].At < runners[j].At
	})
}

func configurerForRole(role runner.Role, flow config.FlowConfig) config.Configurer {
	if role == runner.RoleClient {
		return &flow.Client.Config
	}

	return &flow.Server.Config
}

func atForRole(role runner.Role, flow config.FlowConfig) []time.Duration {
	if role == runner.RoleClient {
		return flow.Client.At
	}

	return flow.Server.At
}

func run(role runner.Role) {
	log, err := NewLogger(fmt.Sprintf("[%s] ", viper.GetString("log.tag")))
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	flows, err := loadFlows(viper.GetString("flows.dir"))
	if err != nil {
		log.Fatalf("cannot load flows: %v", err)
	}

	runners, err := setupRunners(flows, log, role)
	if err != nil {
		log.Fatalf("cannot load flows: %v", err)
	}

	// XXX this queue is consumed by Telegraf, but what if there is no
	// consumer?  The runners will not be .  We should probably do something
	// different here...
	stats := make(chan RunnerStats, 1024)
	go httpStats(viper.GetString("http.stats"), stats)

	done := make(chan StatusReport)
	go sched(role, runners, log, done, stats)

	if err = wait(done); err != nil {
		log.Printf("waiting: %v", err)
	}
}

func httpStats(addrport string, runnerStats chan RunnerStats) {
	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		var m string
		select {
		case rs := <-runnerStats:
			m = rs.out
		default:
			m = "{}"
		}

		w.Write([]byte(m))
	})

	err := http.ListenAndServe(addrport, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func setupRunners(flows FlowConfigs, log *log.Logger, role runner.Role) (Runners, error) {
	var runners Runners

	for _, flow := range flows {
		for _, at := range atForRole(role, flow) {
			cfg := configurerForRole(role, flow)
			cruncher := cruncher.NewTelegrafCruncher()
			runner, err := runner.NewRunner(role, log, at, flow.Label, cfg, cruncher)
			if err != nil {
				log.Printf("cannot create %s %s: %v", role, flow.Label, err)
				return nil, err
			}

			runners = append(runners, *runner)
		}
	}

	return runners, nil
}

func schedFreq(f string) (time.Duration, error) {
	if f == "" {
		return DefaultSchedFrequency, nil
	}

	return time.ParseDuration(f)
}

func sched(role runner.Role, runners Runners, log *log.Logger, done chan StatusReport, stats chan RunnerStats) {
	tickFreq, err := schedFreq(viper.GetString("schedule.tick"))
	if err != nil {
		log.Fatalf("cannot set the scheduler tick frequency: %v", err)
	}

	R = make(RunnersMap)

	// runners MUST be ordered by deadline for the scheduler to work
	sortRunnersByDeadline(runners)

	// runner          [0]    [1][2][3]         [4]   [5]
	//   |--------------O------O--O--O-----------O-----O---->
	// ticks   ^     ^     ^     ^     ^     ^     ^     ^
	//
	// Every 100ms, walk the list of runners that have not been scheduled yet
	// and pull out those whose deadline is expired.
	epoch := time.Now()
	last := 0
	finished := false

	for now := range time.Tick(tickFreq) {
		for i := last; i < len(runners); i++ {
			runner := runners[i]

			if time.Since(epoch) > runner.At {
				log.Printf("%v -> deadline elapsed for %s", now, runner.Label)

				err := runner.Start()
				if err != nil {
					log.Fatalf("cannot start %s %s: %v", role, runner.Label, err)
				}

				M.Lock()
				R[runner.Label] = runner
				M.Unlock()

				// Start a watchdog for this iperf3 instance
				go watchdog(runner, runner.Label, done, stats)

				if i == len(runners)-1 {
					// All runners have been scheduled, job done
					finished = true
					break
				}
			} else {
				// Set a mark for where we need to restart and break out to the
				// ticker
				last = i
				break
			}
		}

		if finished {
			break
		}
	}
}

func watchdog(r runner.Runner, label string, done chan StatusReport, stats chan RunnerStats) {
	out, err := r.Wait()
	if err != nil {
		log.Printf("reaping %s %s: %v", r.Role, label, err)
	}

	M.Lock()
	delete(R, label)
	M.Unlock()

	// send stats for this run to the HTTP endpoint
	stats <- RunnerStats{string(out)}

	// send a status report for this run to the waiter
	done <- StatusReport{label, r.Role, err}
}

func wait(done chan StatusReport) error {
	// XXX if server and not one-off, there's no way this loop can break
	// XXX we need another (interactive - keyboard, signals) source of
	// XXX events
	for {
		select {
		case s := <-done:
			// If one has failed, flag the test as invalid and bail out
			if s.Error != nil {
				tearDownRunners()
				return fmt.Errorf("%v %s failed: %v", s.Role, s.Label, s.Error)
			}

			log.Printf("%v %s finished ok", s.Role, s.Label)

			M.Lock()
			left := len(R)
			M.Unlock()

			if left == 0 {
				log.Printf("all currently active %v(s) finished ok", s.Role)
				//return nil
				break
			}

			log.Printf("%d %v(s) to go", left, s.Role)
		}
	}
}

func tearDownRunners() {
	M.Lock()
	defer M.Unlock()

	for k, v := range R {
		log.Printf("killing %s", k)
		v.Kill()
	}
}
