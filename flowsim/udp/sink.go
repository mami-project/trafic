package udp

import (
	"fmt"
	"net"
)

func Sink(ip string, port int,verbose bool) {
	fmt.Printf("Starting UDP sink at %s:%d\n", ip, port)

    ServerAddr,err := net.ResolveUDPAddr("udp",fmt.Sprintf("%s:%d", ip, port))
    CheckError(err)
	Conn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer Conn.Close()

	buf      := make([]byte, 64 * 1024)
	// mdelay   := int64(0)
	// mpackets := int64(0)
	stats    := &Stats { 0,0,0, 0,0,0 }

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
		udelay := tStamp - info.tStamp
		AddSample(stats, int(udelay), int(info.pktId))
		if verbose {
			fmt.Printf("Delay was: %d us\n", udelay)
		}
		// mdelay += udelay
		// mpackets ++

		if (info.pktId == info.total) {
			_,_,err := Conn.ReadFromUDP(buf) // discard last resort
			if err != nil {
				fmt.Printf("Error: %v\n",err)
			}
			PrintStats(fromUDP, stats, "us")
			break
		}
		//
		// Just in case we lost the last packet!
		//
		if (info.pktId == -1) {
			PrintStats(fromUDP, stats, "us")
			break
		}
	}
	// last = 0
}
