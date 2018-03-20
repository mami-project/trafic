package cruncher

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
)

type TelegrafCruncher struct{}

func NewTelegrafCruncher() Cruncher {
	return &TelegrafCruncher{}
}

// Crunch takes an iperf3 JSON report and transforms it into a
// Telegraf-friendly representation
func (t *TelegrafCruncher) Crunch(j []byte) ([]byte, error) {
	var tcpFlowStats TCPFlowStats
	var udpFlowStats UDPFlowStats

	err := json.Unmarshal(j, &tcpFlowStats)
	if err == nil && tcpFlowStats.Start.TestStart.Protocol == "TCP" {
		return crunchTCP(tcpFlowStats)
	} else if err = json.Unmarshal(j, &udpFlowStats); err == nil {
		return crunchUDP(udpFlowStats)
	} else {
		return nil, err
	}
}

func formatFlowID(title string, cookie string, start int, sd int) string {
	// Use the auto-generated cookie if title has not been explicitly set
	if title == "" {
		title = cookie
	}

	flowID := fmt.Sprintf("%s-%d-%d", title, start, sd)

	return flowID
}

func crunchTCP(tcpFlowStats TCPFlowStats) ([]byte, error) {
	// TCPFlowSample contains a set of measures relative to a TCP flow, sampled
	// in the [.Start, .End] period.
	type TCPFlowSample struct {
		ID             string  `json:"id"`
		Timestamp      string  `json:"timestamp"`
		SampleDuration float64 `json:"sample-duration-s"`
		Bytes          int     `json:"bytes"`
		Bps            float64 `json:"bps"`
		Retransmits    int     `json:"retransmits"`
		SndCwnd        int     `json:"snd-cwnd"`
		RttMs          float64 `json:"rtt-ms"`
		RttVar         int     `json:"rtt-var"`
		Pmtu           int     `json:"pmtu"`
	}

	var tcpFlowSamples []TCPFlowSample
	var tcpFlowSample TCPFlowSample

	start := tcpFlowStats.Start.Timestamp.Timesecs

	for _, interval := range tcpFlowStats.Intervals {
		for _, stream := range interval.Streams {
			flowID := formatFlowID(tcpFlowStats.Title, tcpFlowStats.Start.Cookie, start, stream.Socket)

			timestamp := formatTimestamp(start, stream.Start)

			tcpFlowSample = TCPFlowSample{
				ID:             flowID,
				Timestamp:      timestamp,
				SampleDuration: stream.End - stream.Start,
				Bytes:          stream.Bytes,
				Bps:            stream.BitsPerSecond,
				Retransmits:    stream.Retransmits,
				SndCwnd:        stream.SndCwnd,
				RttMs:          float64(stream.Rtt) / 1000,
				RttVar:         stream.Rttvar,
				Pmtu:           stream.Pmtu,
			}
		}

		tcpFlowSamples = append(tcpFlowSamples, tcpFlowSample)
	}

	return json.Marshal(tcpFlowSamples)
}

func formatTimestamp(flowStart int, sampleStart float64) string {
	tsec, tnsec := math.Modf(float64(flowStart) + sampleStart)

	return time.Unix(int64(tsec), int64(tnsec*(1e9))).String()
}

func crunchUDP(udpFlowStats UDPFlowStats) ([]byte, error) {
	// UDPFlowSample contains a set of measures relative to a UDP flow, sampled
	// in the [.Start, .End] period.
	type UDPFlowSample struct {
		ID             string  `json:"id"`
		Timestamp      string  `json:"timestamp"`
		SampleDuration float64 `json:"sample-duration-s"`
		Bytes          int     `json:"bytes"`
		Bps            float64 `json:"bps"`
		JitterMs       float64 `json:"jitter-ms"`
		LostPackets    int     `json:"lost-packets"`
		LostPercent    float64 `json:"lost-percent"`
		Packets        int     `json:"packets"`
	}

	var udpFlowSamples []UDPFlowSample
	var udpFlowSample UDPFlowSample

	start := udpFlowStats.ServerOutputJSON.Start.Timestamp.Timesecs

	for _, interval := range udpFlowStats.ServerOutputJSON.Intervals {
		for _, stream := range interval.Streams {
			flowID := formatFlowID(udpFlowStats.Title, udpFlowStats.Start.Cookie, start, stream.Socket)

			timestamp := formatTimestamp(start, stream.Start)

			udpFlowSample = UDPFlowSample{
				ID:             flowID,
				Timestamp:      timestamp,
				SampleDuration: stream.End - stream.Start,
				Bytes:          stream.Bytes,
				Bps:            stream.BitsPerSecond,
				JitterMs:       stream.JitterMs,
				LostPackets:    stream.LostPackets,
				LostPercent:    stream.LostPercent,
				Packets:        stream.Packets,
			}
		}

		udpFlowSamples = append(udpFlowSamples, udpFlowSample)
	}

	return json.Marshal(udpFlowSamples)
}
