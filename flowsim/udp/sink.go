package udp

import (
	"fmt"
	"net"
	"strconv"
	common "github.com/mami-project/trafic/flowsim/common"
)

//
// Statistics are encapulated in the Stats structure
// Handled in stats.go
//
func Sink(ip string, port int, multi bool, verbose bool) {
    destAddrStr := net.JoinHostPort(ip,strconv.Itoa(port))
        if verbose {
            fmt.Printf("Starting UDP sink at %s\n", destAddrStr)
        }


    ServerAddr,err := net.ResolveUDPAddr("udp",destAddrStr)
    common.CheckError(err)
	Conn, err := net.ListenUDP("udp", ServerAddr)
	common.CheckError(err)
	defer Conn.Close()

	buf      := make([]byte, 64 * 1024)
	stats := make(map[string]*Stats)

	for {
		n,fromUDP,err := Conn.ReadFromUDP(buf)
		tStamp := MakeTimestamp()

		src := []byte(net.IP.To16(fromUDP.IP))
		src = append(src, (byte)(fromUDP.Port & 0xff))
		srcs := string(append(src, (byte)((fromUDP.Port >> 8) & 0xff)))
		// srcs := fmt.Sprintf("%v", fromUDP)

		_, ok := stats[srcs]
		if ok == false {
			if verbose {
				fmt.Printf("Creating stats for %s\n",fmt.Sprintf("%v",fromUDP))
			}
			stats[srcs] = &Stats{0,0,0,0,0,0,0}
		}
		if verbose {
			fmt.Printf("stats: %v\n",stats)
		}
		if common.CheckError(err) != nil {
			continue
		}
		info := DecodePacket(buf[0:n])
		//
		// Just in case we lose the last packet
		// We send a packet with pktId = -1
		//
		if (info.pktId == -1) {
			PrintStats(fmt.Sprintf("%v",fromUDP), stats[srcs],  "us")
			if multi {
				continue
			}
			break
		}
		udelay := tStamp - info.tStamp
		stats[srcs] = AddSample(stats[srcs], int(udelay), int(info.pktId))
		if verbose {
			fmt.Printf("Delay was: %d us\n", udelay)
		}
		//
		// TODO: define how to handle reordered packets after the last packet
		//
		if (info.pktId == info.total) {
			_,_,err := Conn.ReadFromUDP(buf) // discard last resort packet
			common.CheckError(err)
			PrintStats(fmt.Sprintf("%v",fromUDP), stats[srcs],  "us")
			if multi {
				continue
			}
			break
		}
	}
	// last = 0
}
