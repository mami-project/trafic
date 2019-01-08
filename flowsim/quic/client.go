package quic

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"time"

	"crypto/tls"
	"math/rand"

	quic "github.com/lucas-clemente/quic-go"
	common "github.com/mami-project/trafic/flowsim/common"
)

func Client(ip string, port int, iter int, interval int, bunch int, dscp int) error {

	addr := net.JoinHostPort(ip, strconv.Itoa(port))
	updAddr, err := net.ResolveUDPAddr("udp", addr)
	if common.FatalError(err) != nil {
		return err
	}

	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: 0})
	if common.FatalError(err) != nil {
		return err
	}

	err = common.SetUdpTos(udpConn, dscp)
	if common.FatalError(err) != nil {
		return err
	}

	session, err := quic.Dial(udpConn, updAddr, addr, &tls.Config{InsecureSkipVerify: true}, nil)
	if common.FatalError(err) != nil {
		return err
	}
	defer session.Close()

	fmt.Printf("Opened session for %s\n", addr)
	buf := make([]byte, bunch)
	stream, err := session.OpenStreamSync()
	if common.FatalError(err) != nil {
		return err
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	initWait := r.Intn(interval*50) / 100.0
	time.Sleep(time.Duration(initWait) * time.Second)

	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()
	mkTransfer(stream, buf, 1, iter, time.Now())
	currIter := 2

	if iter > 1 {
		done := make(chan bool, 1)
		for {
			select {
			case t := <-ticker.C:
				mkTransfer(stream, buf, currIter, iter, t)
				currIter++
				if currIter > iter {
					close(done)
				}
			case <-done:
				fmt.Printf("\nFinished...\n\n")
				return nil
			}
		}
	}
	return nil
}

func mkTransfer(stream quic.Stream, buf []byte, current int, iter int, t time.Time) error {

	message := fmt.Sprintf("GET %d/%d %d\n", current, iter, len(buf))
	fmt.Printf("Client: (%v) Sending > %s", t, message)

	_, err := stream.Write([]byte(message))
	if common.FatalError(err) != nil {
		return err
	}

	n, err := io.ReadFull(stream, buf)
	common.FatalError(err)
	fmt.Printf("Client: Got %d bytes back\n", n)
	return nil
}
