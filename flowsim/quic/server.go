package quic

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"math/big"

	"bufio"
	"errors"
	"fmt"
	"net"
	"regexp"
	"strconv"

	quic "github.com/lucas-clemente/quic-go"
	common "github.com/mami-project/trafic/flowsim/common"
)

// Start a server that echos all data on the first stream opened by the client
func Server(ip string, port int, single bool) error {

	addr := net.JoinHostPort(ip, strconv.Itoa(port))

	listener, err := quic.ListenAddr(addr, generateTLSConfig(), nil)
	if common.CheckError(err) != nil {
		return err
	}

	for {
		sess, err := listener.Accept()
		if common.CheckError(err) != nil {
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

	// fmt.Println("Entering quicHandler")
	stream, err := sess.AcceptStream()
	if common.CheckError(err) != nil {
		return err
	}

	reader := bufio.NewReader(stream)
	for {
		// fmt.Println("In server loop")
		cmd, err := reader.ReadString('\n')
		if common.CheckError(err) != nil {
			return err
		}
		wbuf, end, err := parseCmd(cmd)
		if common.CheckError(err) != nil {
			return err
		}
		_, err = stream.Write(wbuf)
		if common.CheckError(err) != nil {
			return err
		}
		if end {
			break
		}
	}
	return err
}

// From flowsim TCP
func matcher(cmd string) (string, string, string, error) {
	expr := regexp.MustCompile(`GET (\d+)/(\d+) (\d+)`)
	parsed := expr.FindStringSubmatch(cmd)
	if len(parsed) == 4 {
		return parsed[1], parsed[2], parsed[3], nil
	}
	return "", "", "", errors.New(fmt.Sprintf("Unexpected request %s", cmd))
}

/*
* Purpuse: parse get Command from client
*         and generate a buffer with random bytes
* Return: byte buffer to send or nil on error
*         boolean: true id last bunch
*         error or nil if all went well
*
* Uses crypto/rand, which is already imported for key handling
 */
func parseCmd(strb string) ([]byte, bool, error) {
	fmt.Printf("Server: Got %s", strb)
	iter, total, bunchStr, err := matcher(strb)
	if err == nil {
		bunch, _ := strconv.Atoi(bunchStr) // ignore error, wouldn't have parsed the command
		nb := make([]byte, bunch, bunch)
		_, err := rand.Read(nb)
		if err != nil {
			fmt.Println("ERROR while filling random buffer: ", err)
			return nil, iter == total, err
		}
		fmt.Printf("Sending %d bytes\n", len(nb))
		return nb, iter == total, err
	}
	return nil, false, err
}

// Setup a bare-bones TLS config for the server
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if common.CheckError(err) != nil {
		return nil
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if common.CheckError(err) != nil {
		return nil
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if common.CheckError(err) != nil {
		return nil
	}
	return &tls.Config{Certificates: []tls.Certificate{tlsCert}}
}
