package flow

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"regexp"
	"io"
	"log"
	"os"
	"errors"
	"syscall"
)

func matcher(cmd string) (string, string, string, error) {
	expr := regexp.MustCompile(`GET (\d+)/(\d+) (\d+)`)
	parsed := expr.FindStringSubmatch(cmd)
	if len(parsed) == 4 {
        return parsed[1], parsed[2], parsed[3], nil
	}
	return "", "", "", errors.New(fmt.Sprintf("Unexpected request %s",cmd))
}

func setTos(tcpConn *net.TCPConn, tos int) (error) {
	f, err := tcpConn.File()

    if err != nil {
		fmt.Printf("While setting TOS to %d on %v: %v\n", tos, f, err)
        return err
    }
	//
	// TODO
	//
    err = syscall.SetsockoptInt(int(f.Fd()), syscall.IPPROTO_IP, syscall.IP_TOS, tos)
    if err != nil {
		fmt.Printf("While setting TOS to %d: %v\n", tos, err)
    }
	return err
}

func closeFdSocket (conn *net.TCPConn) (error){
	f, err := conn.File()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	f.Close()
	conn.Close()
	fmt.Println("Closed fd and socket")
	return nil
}



func handleConn (conn *net.TCPConn) {
	var run, total, bunch string

	//	defer conn.Close()
	defer closeFdSocket(conn)
	zero, err := os.Open("/dev/zero")
	defer zero.Close()
	if err != nil {
		log.Fatal(err)
	}
	for {
		// will listen for message to process ending in newline (\n)
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}
		// output message received
		fmt.Print("Message Received:", string(message))

		// Checked in the client
		run, total, bunch, err = matcher(strings.ToUpper(string(message)))
		if err != nil {
			log.Fatal(err)
			continue
		}
		// fmt.Println(run, total, bunch)
		run_iter,   _ := strconv.Atoi(run)
		total_iter, _ := strconv.Atoi(total)
		bunch_len,  _ := strconv.Atoi(bunch)

		// conn.Write([]byte(fmt.Sprintf("run %d of %d... should send %d bytes\n",run_iter, total_iter, bunch_len)))

		testBunch := make([]byte, bunch_len)
		numRead, err := io.ReadFull(zero, testBunch)

		// fmt.Printf("Read %d bytes from /dev/zero\n",len(testBunch))
		if err != nil {
			log.Fatal(err)
			continue
		}
		fmt.Printf("Sending %d bytes...\n",numRead)
		conn.Write(testBunch)
		if run_iter == total_iter {
			// fmt.Println("This should kill this TCP server thread")
			break
		}
	}

	fmt.Println("Connection closed...")
}

func Server(ip string, port int, single bool, tos int) {

	listenAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Printf("Error resolving %s:%d (%v)\n", ip, port, err)
		return
	}

	ln, err := net.ListenTCP("tcp", listenAddr)
	if err != nil {
		fmt.Printf("Error binding server to %s\n", listenAddr)
		return
	}
	fmt.Printf("Listening at %s\n",listenAddr)
	for {
		// accept connection on port
		conn, err := ln.AcceptTCP()
		if err != nil {
			fmt.Printf("Error accepting connection\n")
			continue
		}
		err = setTos (conn, tos)
		if err != nil {
			continue
		}
		if single {
			handleConn(conn)
			break
		} else {
			go handleConn(conn)
		}
	}
}
