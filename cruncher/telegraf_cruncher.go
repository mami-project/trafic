package cruncher

import (
	"encoding/json"
)

type TelegrafCruncher struct{}

func NewTelegrafCruncher() Cruncher {
	return &TelegrafCruncher{}
}

func (c *TelegrafCruncher) CrunchTCP(tcpFlowStats TCPFlowStats) ([]byte, error) {
	var tcpFlowSamples []TCPFlowSample
	var tcpFlowSample TCPFlowSample

	start := tcpFlowStats.ServerOutputJSON.Start.Timestamp.Timesecs

	for _, interval := range tcpFlowStats.ServerOutputJSON.Intervals {
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
	var udpFlowSamples []UDPFlowSample
	var udpFlowSample UDPFlowSample

	start := udpFlowStats.Start.Timestamp.Timesecs

	for _, interval := range udpFlowStats.Intervals {
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
