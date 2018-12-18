package udp

import (
	"fmt"
	"net"
	"time"
	"syscall"
)

func Source(ip string, port int, duration int, pps int, psize int, tos int, verbose bool) {
	fmt.Printf("Starting server at %s:%d for %d secs at %d pps for %d byte packets (TOS: %02x)\n",
		ip, port, duration, pps, psize, tos)

	var maxpackets int64
	maxpackets = int64(duration * pps)

    ServerAddr,err := net.ResolveUDPAddr("udp",fmt.Sprintf("%s:%d",ip,port))
    CheckError(err)

    LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
    CheckError(err)

    Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
    CheckError(err)

	err = setTos (Conn, tos)
	CheckError(err)

	fmt.Println("Starting to send to ",ServerAddr)
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

func setTos(Conn *net.UDPConn, tos int) (error) {
	f, err := Conn.File()

    if err != nil {
		fmt.Printf("While setting TOS to %d on %v: %v\n", tos, f, err)
        return err
    }

    err = syscall.SetsockoptInt(int(f.Fd()), syscall.IPPROTO_IP, syscall.IP_TOS, tos)
    if err != nil {
		fmt.Printf("While setting TOS to %d: %v\n", tos, err)
    }
	return err
}
