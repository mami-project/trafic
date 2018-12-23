package udp

import (
	"fmt"
	"net"
	"time"
	// "syscall"
	"strconv"
	common "github.com/mami-project/trafic/flowsim/common"
)

func Source(ip string, port int, localip string,duration int, pps int, psize int, tos int, verbose bool) {
	fmt.Printf("Starting server at %s:%d for %d secs at %d pps for %d byte packets (TOS: %02x)\n",
		ip, port, duration, pps, psize, tos)

	var maxpackets int64
	maxpackets = int64(duration * pps)

	destAddrStr := net.JoinHostPort(ip,strconv.Itoa(port))
	srcAddrStr  := net.JoinHostPort(localip,"0")
	fmt.Println("To   ",destAddrStr)
	fmt.Println("From ",srcAddrStr)

    ServerAddr,err := net.ResolveUDPAddr("udp",destAddrStr)
    CheckError(err)
    LocalAddr, err := net.ResolveUDPAddr("udp", srcAddrStr)
    CheckError(err)

    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    CheckError(err)

	f, err := Conn.File()
	if err != nil {
		CheckError(err)
	}
	err = common.SetTos (f, tos, net.IP.To4(ServerAddr.IP) == nil)
	CheckError(err)

	fmt.Printf("Starting to send to %v\n",ServerAddr)
    defer Conn.Close()
	var msg string
	var packet myStruct

	packet.total = maxpackets
	done := make(chan bool, 1)
	ticker := time.NewTicker(time.Duration(1000000 / pps) * time.Microsecond)
	defer ticker.Stop()
	packet.pktId = 1
	for {
		select {
		case t := <-ticker.C:
			packet.tStamp = toTimestamp(t)
			_,err := Conn.Write(EncodePacket(packet, psize))
			if err != nil {
				fmt.Println(msg, err)
			}
			if verbose {
				fmt.Printf("Sent %4d of %4d at %v\n", packet.pktId, maxpackets, t)
			}
			packet.pktId++
			if (packet.pktId > maxpackets) {
				packet.pktId = -1;
				_,err = Conn.Write(EncodePacket(packet, psize))
				close(done)
			}
		case <-done:
			return
		}
	}
}
