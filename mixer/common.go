package mixer

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type BurstProps struct {
	Label          string
	Port           uint16
	Server         string
	ReportInterval string
	At             []int
}

type FixedBitrateProps struct {
	Label          string
	Port           uint16
	Server         string
	ReportInterval string
	Time           string
	Instances      uint
}

func writeFixedBitrate(outFile string, defaultTmpl string, g GlobalDesc,
	c FlowDesc, flowBitrate float64) error {

	tmpl, err := setupTemplate(c, defaultTmpl)
	if err != nil {
		return err
	}

	props, err := makeFixedBitrateProps(g, c, flowBitrate)
	if err != nil {
		return err
	}

	out, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("cannot create %s: %v", outFile, err)
	}
	defer out.Close()

	err = tmpl.Execute(out, props)
	if err != nil {
		return fmt.Errorf("cannot execute template: %v", err)
	}

	return nil
}

func flowQuotaBps(total Bytes, percent Ratio) float64 {
	return float64(total.Val) * 8 * percent.Val
}

func evalInstancesFixedBitrate(total Bytes, percent Ratio, flowBitrate float64) uint {
	return uint(flowQuotaBps(total, percent) / flowBitrate)
}

func makeFixedBitrateProps(g GlobalDesc, c FlowDesc, flowBitrate float64) (*FixedBitrateProps, error) {
	instances := uint(flowQuotaBps(g.TotalBandwidth, c.PercentBandwidth) / flowBitrate)

	return &FixedBitrateProps{
		Port:           c.PortsRange.First,
		Instances:      instances,
		Time:           fmt.Sprintf("%fs", g.TotalTime.Seconds()),
		Server:         c.Props["server"],
		ReportInterval: fmt.Sprintf("%fs", g.ReportInterval.Seconds()),
		Label:          c.Props["label"],
	}, nil
}

func writeBursting(outFile string, defaultTmpl string, g GlobalDesc,
	c FlowDesc, burstSize float64, burstInterval time.Duration) error {

	numBursts := g.TotalTime.Seconds() *
		flowQuotaBps(g.TotalBandwidth, c.PercentBandwidth) /
		burstSize

	numIntervalsPerClient := g.TotalTime.Seconds() /
		burstInterval.Seconds()

	numClients := numBursts / numIntervalsPerClient

	clientSchedule := makeBurstScheduler(
		int(numClients),
		int(numIntervalsPerClient),
		burstInterval,
	)

	return writeBurstingClients(outFile, defaultTmpl, g, c, clientSchedule)
}

func writeBurstingClients(outFile string, defaultTmpl string, g GlobalDesc,
	c FlowDesc, clientSchedule [][]int) error {
	tmpl, err := setupTemplate(c, defaultTmpl)
	if err != nil {
		return err
	}

	for clientId := range clientSchedule {
		fn := fmt.Sprintf("%s-%d.yaml", outFile, clientId)

		out, err := os.Create(fn)
		if err != nil {
			return fmt.Errorf("cannot create %s: %v", fn, err)
		}
		defer out.Close()

		props, err := makeBurstProps(g, c, clientSchedule[clientId], clientId)
		if err != nil {
			return err
		}

		err = tmpl.Execute(out, props)
		if err != nil {
			return fmt.Errorf("cannot execute template: %v", err)
		}
	}

	return nil
}

func makeBurstProps(g GlobalDesc, c FlowDesc, clientSchedule []int, clientId int) (*BurstProps, error) {
	flowPort := c.PortsRange.First + uint16(clientId)
	if flowPort > c.PortsRange.Last {
		return nil, fmt.Errorf("ports exhausted: can't go past %u", c.PortsRange.Last)
	}

	return &BurstProps{
		At:             clientSchedule,
		Label:          c.Props["label"],
		Port:           flowPort,
		Server:         c.Props["server"],
		ReportInterval: fmt.Sprintf("%fs", g.ReportInterval.Seconds()),
	}, nil
}

func makeBurstScheduler(numClients int, numIntervalsPerClient int,
	burstInterval time.Duration) [][]int {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// one "burst scheduler" per client
	burstScheduler := make([][]int, numClients)
	for c := range burstScheduler {
		// pick a random offset in [0, burstInterval]
		perClientOffset := r.Intn(int(burstInterval.Seconds()))

		// make room for the bursts' timings
		burstScheduler[c] = make([]int, numIntervalsPerClient)
		for i := range burstScheduler[c] {
			burstScheduler[c][i] = perClientOffset + (i * int(burstInterval.Seconds()))
		}
	}

	return burstScheduler
}
