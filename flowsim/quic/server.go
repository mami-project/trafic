package quic
import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"math/big"

	"fmt"
	"bufio"
	"regexp"
	"strconv"
	"errors"
	//	"strings"
	//	"log"

	quic "github.com/lucas-clemente/quic-go"
)

// const addr = "localhost:4242"

// We start a server echoing data on the first stream the client opens,
// then connect with a client, send the message, and wait for its receipt.
// func main() {
// 	err := Server("localhost", 4242)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// Start a server that echos all data on the first stream opened by the client
func Server(ip string, port int, single bool) error {
	addr := fmt.Sprintf("%s:%d", ip, port)
	listener, err := quic.ListenAddr(addr, generateTLSConfig(), nil)
	if err != nil {
		return err
	}

	for {
		sess, err := listener.Accept()
		if err != nil {
			return err
		}
		if single {
			quicHandler(sess)
			return nil
		}
		go quicHandler(sess)
	}
}

func quicHandler(sess quic.Session) error {

	fmt.Println("Entering quicHandler")
	stream, err := sess.AcceptStream()
	if err != nil {
		panic(err)
		return err
	}

	reader := bufio.NewReader(stream)
	for {
		// fmt.Println("In server loop")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		wbuf, end, err := parseCmd(cmd)
		if err != nil {
			return err
		}
		_,err = stream.Write(wbuf)
		if err != nil {
			return err
		}
		if end {
			break
		// } else {
		// 	fmt.Println("Continuing")
		}
		// Echo through the loggingWriter
		// _, err = io.Copy(loggingWriter{stream}, stream)
	}
	return err
}

// From flowim TCP
func matcher(cmd string) (string, string, string, error) {
	expr := regexp.MustCompile(`GET (\d+)/(\d+) (\d+)`)
	parsed := expr.FindStringSubmatch(cmd)
	if len(parsed) == 4 {
        return parsed[1], parsed[2], parsed[3], nil
	}
	return "", "", "", errors.New(fmt.Sprintf("Unexpected request %s",cmd))
}

// A wrapper for io.Writer that also logs the message.
//type loggingWriter struct{ io.Writer }

//func (w loggingWriter) Write(b []byte) (int, error) {
//	strb := string(b)
func parseCmd(strb string) ([]byte, bool, error) {
	stop := false
	fmt.Printf("Server: Got %s", strb)
	iter, total, bunchStr, err := matcher(strb)
	if err == nil {
		bunch, _ := strconv.Atoi(bunchStr) // ignore error, wouldn't have parsed the command
		nb := make([]byte,bunch,bunch)
		if iter == total {
			stop = true
		}
		// fmt.Printf("Sending %v\n", nb)
		fmt.Printf("Sending %d bytes\n", len(nb))
		return nb, stop, err
	}
	return nil, false, err
}

// Setup a bare-bones TLS config for the server
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{Certificates: []tls.Certificate{tlsCert}}
}
