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

// const addr = "localhost:4242"

// We start a server echoing data on the first stream the client opens,
// then connect with a client, send the message, and wait for its receipt.
// func main() {
// 	err := Client("localhost", 4242, 6 , 20, 10)
// 	if err != nil {
// 		panic(err)
// 	}
// }

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
	go mkTransfer (stream, buf, currIter, iter)

	ticker := time.Tick(time.Duration(interval) * time.Second)
	for now := range ticker {
		// read in input from stdin

		currIter ++
		if currIter > iter {
			break
		}
		fmt.Printf("Launching at %v\n", now)
		go mkTransfer (stream, buf, currIter, iter)

	}
	fmt.Printf("\nFinished...\n\n")
	return nil
}

func mkTransfer (stream quic.Stream, buf []byte, current int, iter int) error {

	message := fmt.Sprintf("GET %d/%d %d\n", current, iter, len(buf))
	fmt.Printf("Client: Sending > %s", message)

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
