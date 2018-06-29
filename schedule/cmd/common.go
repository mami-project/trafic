package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
	Runner runner.Runner
	Stats  []byte
	Error  error
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

func loadFlows(dirs string) (FlowConfigs, error) {
	var flows FlowConfigs

	for _, dir := range strings.Split(dirs, ",") {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Printf("cannot read dir: %s: %v", dir, err)
			return nil, err
		}

		for _, file := range files {
			if !strings.HasSuffix(file.Name(), ".yaml") {
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

	flows, err := loadFlows(viper.GetString("flows.dirs"))
	if err != nil {
		log.Fatalf("cannot load flows: %v", err)
	}

	runners, err := setupRunners(flows, log, role)
	if err != nil {
		log.Fatalf("cannot load flows: %v", err)
	}

	stats := make(chan RunnerStats, 1024)
	go statStorer(stats)

	done := make(chan StatusReport)
	go sched(role, runners, log, done, stats)

	if err = wait(done); err != nil {
		log.Printf("waiting: %v", err)
	}
}

func statStorer(runnerStats <-chan RunnerStats) {

	var statFilesEnabled bool = viper.GetBool("stats.enabled")
	if statFilesEnabled {
		err := os.MkdirAll(viper.GetString("stats.dir"), 0755)
		if err != nil {
			log.Fatalf("cannot create stats directory %v", err)
		}
	}

	var influxDBConfigured bool = false

	if viper.GetBool("influxdb.enabled") {
		err := setupInfluxDB()
		if err != nil {
			log.Println("cannot setup InfluxDB endpoint:", err)
		} else {
			influxDBConfigured = true
		}
	}

	for {
		select {
		case rs := <-runnerStats:
			if rs.Error != nil {
				log.Println("cannot collect stats for failed runner", rs.Runner.Label, ":", rs.Error)
				break
			}

			if statFilesEnabled {
				err := saveRunnerStats(rs)
				if err != nil {
					log.Println("cannot save to CSV:", err)
				}
			}

			if influxDBConfigured {
				err := forwardToInfluxDB(rs)
				if err != nil {
					log.Println("cannot forward to InfluxDB:", err)
				}
			}
		}
	}
}

func setupInfluxDB() error {
	// curl -XPOST http://localhost:8086/query --data "q=CREATE DATABASE mydb"
	url, err := url.Parse(viper.GetString("influxdb.endpoint"))
	if err != nil {
		return err
	}
	url.Path = "/query"

	q := fmt.Sprintf("q=CREATE DATABASE \"%s\"", viper.GetString("influxdb.db"))

	resp, err := http.Post(url.String(), "application/x-www-form-urlencoded", bytes.NewBufferString(q))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("request %s failed: %s", resp.Header.Get("Request-Id"), string(body))
	}

	return nil
}

func forwardToInfluxDB(runnerStats RunnerStats) error {
	c := cruncher.NewInfluxDBCruncher(viper.GetString("influxdb.measurements"))

	lp, err := cruncher.Crunch(c, runnerStats.Stats)
	if err != nil {
		return err
	}

	u, err := url.Parse(viper.GetString("influxdb.endpoint"))
	if err != nil {
		return err
	}
	q := url.Values{}
	q.Set("db", viper.GetString("influxdb.db"))
	u.RawQuery = q.Encode()
	u.Path = "/write"

	resp, err := http.Post(u.String(), "application/x-www-form-urlencoded", bytes.NewBuffer(lp))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("request %s failed: %s", resp.Header.Get("Request-Id"), string(body))
	}

	log.Printf("POST %s [%s ...] -> %s (this is OK)", u.String(), viper.GetString("influxdb.measurements"), resp.Status)

	return nil
}

// Save raw stats along with the processed CSV
func saveRunnerStats(runnerStats RunnerStats) error {
	err := ioutil.WriteFile(makeStatsFile(runnerStats, ".json"), runnerStats.Stats, 0644)
	if err != nil {
		return err
	}

	csv, err := cruncher.Crunch(cruncher.NewCSVCruncher(), runnerStats.Stats)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(makeStatsFile(runnerStats, ".csv"), csv, os.ModePerm)
}

func makeStatsFile(runnerStats RunnerStats, ext string) string {
	fn := fmt.Sprintf("%s_%s_%s%s",
		time.Now().Format("20060102150405"),
		runnerStats.Runner.Role,
		runnerStats.Runner.Label,
		ext)

	return filepath.Join(viper.GetString("stats.dir"), fn)
}

func setupRunners(flows FlowConfigs, log *log.Logger, role runner.Role) (Runners, error) {
	var runners Runners

	for _, flow := range flows {
		for _, at := range atForRole(role, flow) {
			cfg := configurerForRole(role, flow)
			runner, err := runner.NewRunner(role, log, at, flow.Label, cfg)
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

func watchdog(r runner.Runner, label string, done chan<- StatusReport, stats chan<- RunnerStats) {
	out, err := r.Wait()
	if err != nil {
		log.Printf("reaping %s %s: %v", r.Role, label, err)
	}

	M.Lock()
	delete(R, label)
	M.Unlock()

	// send stats for this run to the HTTP endpoint
	stats <- RunnerStats{r, out, err}

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
