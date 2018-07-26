package flow

import (
	"fmt"
	"net"
	"syscall"
	"errors"
)

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

// A dictionary to map DSCP labels to values
// CAVEAT: Multiply by 4 to get the full TOS byte value

var dscpDict = map[string]int{
	"CS0":  0,
	"CS1":  8,
	"CS2":  16,
	"CS3":  24,
	"CS4":  32,
	"CS5":  40,
	"CS6":  48,
	"CS7":  56,
	"EF":   46,
	"AF11": 10,
	"AF12": 12,
	"AF13": 14,
	"AF21": 18,
	"AF22": 20,
	"AF23": 22,
	"AF31": 26,
	"AF32": 28,
	"AF33": 30,
	"AF41": 34,
	"AF42": 36,
	"AF43": 38,
}

func Dscp(s string) (int, error) {
	var val int
	if val, ok := dscpDict[s]; ok {
		return val, nil
	}
	_, err := fmt.Sscanf(s, "%d", &val)
	if err != nil {
		return 0, errors.New("Unknown DSCP ID ")
	}
	if val < 0 || val > 64 {
		return 0, errors.New("Value out of range [0,64)")
	}
	return val, nil
}
