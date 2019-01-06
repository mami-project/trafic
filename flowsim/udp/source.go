package udp

import (
	"crypto/rand"
	"fmt"
	"net"
	"time"

	common "github.com/mami-project/trafic/flowsim/common"
	"strconv"
)

func Source(ip string, port int, localip string, duration int, pps int, psize int, tos int, verbose bool) {
	fmt.Printf("Starting server at %s:%d for %d secs at %d pps for %d byte packets (TOS: %02x)\n",
		ip, port, duration, pps, psize, tos)

	var maxpackets int64
	maxpackets = int64(duration * pps)

	destAddrStr := net.JoinHostPort(ip, strconv.Itoa(port))
	srcAddrStr := net.JoinHostPort(localip, "0")
	fmt.Println("To   ", destAddrStr)
	fmt.Println("From ", srcAddrStr)

	ServerAddr, err := net.ResolveUDPAddr("udp", destAddrStr)
	common.CheckError(err)
	LocalAddr, err := net.ResolveUDPAddr("udp", srcAddrStr)
	common.CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	common.CheckError(err)

	err = common.SetUdpTos(Conn, tos)
	common.CheckError(err)

	fmt.Printf("Starting to send to %v\n", ServerAddr)
	defer Conn.Close()

	// Initialise packet structure
	var packet myStruct
	_, err = rand.Read(packet.padding[:])
	common.CheckError(err)
	packet.total = maxpackets
	packet.pktId = 1

	done := make(chan bool, 1) // signal end of task, kill ticker

	ticker := time.NewTicker(time.Duration(1000000/pps) * time.Microsecond)
	defer ticker.Stop()
	for {
		select {
		case t := <-ticker.C:
			packet.tStamp = toTimestamp(t) // refresh packet timestamp
			_, err := Conn.Write(EncodePacket(packet, psize))
			common.CheckError(err)

			if verbose {
				fmt.Printf("Sent %4d of %4d at %v\n", packet.pktId, maxpackets, t)
			}

			packet.pktId++ // prepare for next packet
			if packet.pktId > maxpackets {
				packet.tStamp = toTimestamp(t) // not strictly needed, but anyhow
				packet.pktId = -1              // signal end of measurement
				_, err = Conn.Write(EncodePacket(packet, psize))
				close(done)
			}
		case <-done:
			return
		}
	}
}
