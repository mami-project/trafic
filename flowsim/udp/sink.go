package udp

import (
	"fmt"
	"net"
)

func Sink(ip string, port int,verbose bool) {
	fmt.Printf("Starting UDP sink at %s:%d\n", ip, port)
    /* Lets prepare a address at any address at port 10001*/
    ServerAddr,err := net.ResolveUDPAddr("udp",fmt.Sprintf("%s:%d", ip, port))
    CheckError(err)

    /* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	ServeOneConnection(ServerConn, verbose)
}

func ServeOneConnection(Conn *net.UDPConn, verbose bool) {
	defer Conn.Close()

	buf := make([]byte, 64 * 1024)

	var info myStruct
	tStamp := MakeTimestamp()
	// var last int64
	// last = 0
	// loss := 0
	// ldelay   := int64(0)
	mdelay   := int64(0)
	mpackets := int64(0)
	// mjitter  := int64(0)
	stats    := NewStats()

	// var fromUDP *net.UDPAddr
	for {
		n,fromUDP,err := Conn.ReadFromUDP(buf)
		tStamp = MakeTimestamp()
		if verbose {
			fmt.Println("Received ",n, "bytes from ",fromUDP)
		}
		if err != nil {
			fmt.Println("Error: ",err)
			continue
		}
		info = DecodePacket(buf[0:n])
		udelay := tStamp - info.tStamp
		AddSample(stats, int(udelay), int(info.pktId))
		if verbose {
			fmt.Printf("Delay was: %d us\n", udelay)
		}
		mdelay += udelay
		mpackets ++

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
