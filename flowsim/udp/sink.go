package udp

import (
	"fmt"
	"net"
)

//
// Statistics are encapulated in the Stats structure
// Handled in stats.go
//
func Sink(ip string, port int,verbose bool) {
	fmt.Printf("Starting UDP sink at %s:%d\n", ip, port)

    ServerAddr,err := net.ResolveUDPAddr("udp",fmt.Sprintf("%s:%d", ip, port))
    CheckError(err)
	Conn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer Conn.Close()

	buf      := make([]byte, 64 * 1024)
	stats    := &Stats { 0,0,0,0,0,0,0 }

	for {
		n,fromUDP,err := Conn.ReadFromUDP(buf)
		tStamp := MakeTimestamp()
		if verbose {
			fmt.Println("Received ",n, "bytes from ",fromUDP)
		}
		if err != nil {
			fmt.Println("Error: ",err)
			continue
		}
		info := DecodePacket(buf[0:n])
		//
		// Just in case we lose the last packet
		// We send a packet with pktId = -1
		//
		if (info.pktId == -1) {
			PrintStats(fromUDP, stats, "us")
			break
		}
		udelay := tStamp - info.tStamp
		AddSample(stats, int(udelay), int(info.pktId))
		if verbose {
			fmt.Printf("Delay was: %d us\n", udelay)
		}
		//
		// TODO: define how to handle reordered packets after the last packet
		//
		if (info.pktId == info.total) {
			_,_,err := Conn.ReadFromUDP(buf) // discard last resort packet
			if err != nil {
				fmt.Printf("Error: %v\n",err)
			}
			PrintStats(fromUDP, stats, "us")
			break
		}
	}
	// last = 0
}
