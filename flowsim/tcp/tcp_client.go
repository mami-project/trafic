package flow

import (
	"net"
	"fmt"
	"log"
	"io"
	"time"
	"math/rand"
)

func mkTransfer (conn *net.TCPConn, iter int, total int, tsize int) {
	// send to socket
	fmt.Fprintf(conn, fmt.Sprintf("GET %d/%d %d\n", iter, total, tsize))
	// listen for reply
	readBuffer := make([]byte, tsize)
	fmt.Printf("Trying to read %d bytes back...", len(readBuffer))
	readBytes, err := io.ReadFull(conn, readBuffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Effectively read %d bytes\n", readBytes)
}


func Client(host string, port int, iter int, interval int, burst int, tos int) {
	// connect to this socket
	serverAddr := fmt.Sprintf("%s:%d",host,port)
	server, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		fmt.Printf("Error resolving %s: %v\n", serverAddr, err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, server)
	if err != nil {
		fmt.Printf("Error connecting to %s: %v\n", serverAddr, err)
		return
	}
	defer closeFdSocket (conn)
	fmt.Printf("Talking to %s\n",serverAddr)
	err = setTos (conn, tos)
	if err != nil {
		log.Fatal(err)
		return
	}
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
	initWait := r.Intn(interval * 50) / 100.0
	time.Sleep(time.Duration(initWait) * time.Second)

	currIter := 1
	go mkTransfer (conn , currIter, iter, burst)

	ticker := time.Tick(time.Duration(interval) * time.Second)
	for now := range ticker {
		// read in input from stdin

		currIter ++
		if currIter > iter {
			break
		}
		fmt.Printf("Launching at %v\n", now)
		go mkTransfer (conn , currIter, iter, burst)
	}
	fmt.Printf("\nFinished...\n\n")
}
