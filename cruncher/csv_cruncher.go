package cruncher

import (
	"bytes"
	"fmt"
)

type CSVCruncher struct {
}

func NewCSVCruncher() Cruncher {
	return &CSVCruncher{}
}

func (c *CSVCruncher) CrunchTCP(tcpFlowStats TCPFlowStats) ([]byte, error) {
	var lines bytes.Buffer

	// CSV header
	lines.WriteString("Timestamp,FlowID,FlowType,ToS,PMTU,Bytes,BitsPerSecond,Retransmissions,SenderCWND,RTTms,RTTvar\n")

	start := tcpFlowStats.ServerOutputJSON.Start.Timestamp.Timesecs

	for _, interval := range tcpFlowStats.ServerOutputJSON.Intervals {
		for _, stream := range interval.Streams {
			flowID := formatFlowID(tcpFlowStats.Title, tcpFlowStats.Start.Cookie, start, stream.Socket)

			line := fmt.Sprintf("%f,%s,tcp,0x%02x,%d,%d,%f,%d,%d,%f,%d\n",
				float64(start)+stream.Start,
				flowID,
				tcpFlowStats.Start.TestStart.Tos,
				stream.Pmtu,
				stream.Bytes,
				stream.BitsPerSecond,
				stream.Retransmits,
				stream.SndCwnd,
				float64(stream.Rtt)/1000,
				stream.Rttvar)

			// Don't bother about the bytes written, just check the return status
			_, err := lines.WriteString(line)
			if err != nil {
				return nil, err
			}
		}
	}

	return lines.Bytes(), nil
}

func (c *CSVCruncher) CrunchUDP(udpFlowStats UDPFlowStats) ([]byte, error) {
	var lines bytes.Buffer

	// CSV header
	lines.WriteString("Timestamp,FlowID,FlowType,ToS,Bytes,BitsPerSecond,Jitterms,Packets,LostPackets,LostPercent\n")

	start := udpFlowStats.Start.Timestamp.Timesecs

	for _, interval := range udpFlowStats.Intervals {
		for _, stream := range interval.Streams {
			flowID := formatFlowID(udpFlowStats.Title, udpFlowStats.Start.Cookie, start, stream.Socket)

			line := fmt.Sprintf("%f,%s,udp,0x%02x,%d,%f,%f,%d,%d,%f\n",
				float64(start)+stream.Start,
				flowID,
				udpFlowStats.Start.TestStart.Tos,
				stream.Bytes,
				stream.BitsPerSecond,
				stream.JitterMs,
				stream.Packets,
				stream.LostPackets,
				stream.LostPercent)

			_, err := lines.WriteString(line)
			if err != nil {
				return nil, err
			}
		}
	}

	return lines.Bytes(), nil
}
