package quic

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"time"
	"math/rand"
	quic "github.com/lucas-clemente/quic-go"
)

func Client(ip string, port int, iter int, interval int, bunch int) error {

	addr := fmt.Sprintf("%s:%d", ip, port)
	session, err := quic.DialAddr(addr, &tls.Config{InsecureSkipVerify: true}, nil)
	if err != nil {
		return err
	}
	defer session.Close()
	fmt.Printf("Opened session for %s\n", addr)
	buf := make([]byte, bunch)
	stream, err := session.OpenStreamSync()
	if err != nil {
		return err
	}
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
	initWait := r.Intn(interval * 50) / 100.0
	time.Sleep(time.Duration(initWait) * time.Second)

	currIter := 1
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()
	mkTransfer (stream, buf, currIter, iter, time.Now())

	if iter > 1 {
		done := make(chan bool,1)
		for {
			select {
			case t := <-ticker.C:
				mkTransfer (stream, buf, currIter, iter, t)
				currIter ++
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

func mkTransfer (stream quic.Stream, buf []byte, current int, iter int,t time.Time) error {

	message := fmt.Sprintf("GET %d/%d %d\n", current, iter, len(buf))
	fmt.Printf("Client: (%v) Sending > %s", t, message)

	_, err := stream.Write([]byte(message))
	if err != nil {
		return err
	}

	n, err := io.ReadFull(stream, buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Client: Got %d bytes back\n", n)
	return nil
}
