package cruncher

import (
	"encoding/json"
)

type Cruncher interface {
	// A Cruncher takes an iperf3 JSON report (either UDP or TCP) and
	// transforms it into a format that can be natively consumed by a certain
	// storage/processing engine (e.g., Telegraf, InfluxDB, CSV, etc.)
	CrunchUDP(UDPFlowStats) ([]byte, error)
	CrunchTCP(TCPFlowStats) ([]byte, error)
}

func Crunch(c Cruncher, j []byte) ([]byte, error) {
	var tcpFlowStats TCPFlowStats
	var udpFlowStats UDPFlowStats

	err := json.Unmarshal(j, &tcpFlowStats)
	if err == nil && tcpFlowStats.Start.TestStart.Protocol == "TCP" {
		return c.CrunchTCP(tcpFlowStats)
	} else if err = json.Unmarshal(j, &udpFlowStats); err == nil {
		return c.CrunchUDP(udpFlowStats)
	} else {
		return nil, err
	}
}
