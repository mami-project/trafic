package cruncher

import (
	"encoding/json"
)

type Cruncher interface {
	// A Cruncher can take an iperf3 JSON report (either UDP or TCP) and
	// transform it into a format that works for a certain consumer (e.g.,
	// Telegraf, InfluxDB, CSV, etc.)
	CrunchUDP(UDPFlowStats) ([]byte, error)
	CrunchTCP(TCPFlowStats) ([]byte, error)
}

// Telegraf-friendly representation
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
