package cruncher

import (
	"encoding/json"
)

type TelegrafCruncher struct{}

func NewTelegrafCruncher() Cruncher {
	return &TelegrafCruncher{}
}

func (c *TelegrafCruncher) CrunchTCP(tcpFlowStats TCPFlowStats) ([]byte, error) {
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

func (c *TelegrafCruncher) CrunchUDP(udpFlowStats UDPFlowStats) ([]byte, error) {
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
