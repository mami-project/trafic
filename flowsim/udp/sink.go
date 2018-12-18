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

func ServeOneConnection(ServerConn *net.UDPConn, verbose bool) {
	defer ServerConn.Close()

	buf := make([]byte, 64 * 1024)

	var info myStruct
	tStamp := MakeTimestamp()
	var last int64
	last = 0
	loss := 0
	ldelay   := int64(0)
	mdelay   := int64(0)
	mpackets := int64(0)
	mjitter  := int64(0)
	for {
		n,addr,err := ServerConn.ReadFromUDP(buf)
		tStamp = MakeTimestamp()
		if verbose {
			fmt.Println("Received ",n, "bytes from ",addr)
		}
		if err != nil {
			fmt.Println("Error: ",err)
			continue
		}
		info = DecodePacket(buf[0:n])
		udelay := tStamp - info.tStamp
		if verbose {
			fmt.Printf("Delay was: %d us\n", udelay)
		}
		mdelay += udelay
		mpackets ++
		if (ldelay != 0) {
			delta := ldelay - udelay
			if (delta > 0) {
				mjitter += delta
			} else {
				mjitter -= delta
			}
		}
		ldelay = udelay
		if (info.pktId != last + 1) {
			fmt.Printf("Packet lost! (%d) >(%d)\n",last,info.pktId)
			loss++
		}
		last = info.pktId
		if (info.pktId == info.total) {
			n,addr,err = ServerConn.ReadFromUDP(buf) // discard last resort
			break
		}
		//
		// Just in case we lost the last packet!
		//
		if (info.pktId == -1) {
			break
		}
	}
	fmt.Printf("Packet loss: %d/%d\n",loss,info.total)
	fmt.Printf("Mean delay:  %5d us\n",mdelay/mpackets)
	if mpackets > 1 {
		fmt.Printf("Mean jitter: %5d us\n",mjitter/(mpackets - 1))
	}
	last = 0
}
