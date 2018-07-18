package abr

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
	// "flag"
)

func matcher(cmd string) (string, string, string) {
	expr := regexp.MustCompile(`GET (\d+)/(\d+) (\d+)`)
	parsed := expr.FindStringSubmatch(cmd)
	if len(parsed) == 4 {
        return parsed[1], parsed[2], parsed[3]
	}
	return "", "", ""
}

func handleConn (conn net.Conn) {
	var run, total, bunch string

	defer conn.Close()
	zero, err := os.Open("/dev/zero")
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
		run, total, bunch = matcher(strings.ToUpper(string(message)))

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
		}
		fmt.Printf("Sending %d bytes...\n",numRead)
		conn.Write(testBunch)
		if run_iter == total_iter {
			fmt.Println("This should kill this TCP server thread")
			break
		}
	}
	fmt.Println("Connection closed...")
}

func Server(ip string, port int, single bool) {

	listenAddr := fmt.Sprintf("%s:%d", ip, port)
	fmt.Printf("Listening at %s\n",listenAddr)
	// listen on all interfaces
	ln, _ := net.Listen("tcp", listenAddr)

	for {
		// accept connection on port
		conn, err := ln.Accept()
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

// func main() {
// 	ipPtr := flag.String("ip", "127.0.0.1", "The IP address to bind to")
// 	portPtr := flag.Int("port", 8081, "The port to use")
// 	flag.Parse()

// 	abrserver(*ipPtr, *portPtr)
// }
