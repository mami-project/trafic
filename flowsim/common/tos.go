package common

import (
	//  "fmt"
	// 	"os"
	// 	"syscall"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
	"net"
)

func SetTcpTos(conn *net.TCPConn, dscp int) error {
	err := ipv4.NewConn(conn).SetTOS(dscp)
	if err != nil {
		// common.WarnErrorf(err, "while setting TOS")
		err = ipv6.NewConn(conn).SetTrafficClass(dscp)
	}
	return err
}

func SetUdpTos(conn *net.UDPConn, dscp int) error {
	err := ipv4.NewConn(conn).SetTOS(dscp)
	if err != nil {
		// common.WarnErrorf(err, "while setting TOS")
		err = ipv6.NewConn(conn).SetTrafficClass(dscp)
	}

	return err
}

/*
func SetTos(f *os.File, tos int, ipv6 bool) error {

	proto := syscall.IPPROTO_IP
	call := syscall.IP_TOS
	errmsg := "TOS"

	if ipv6 {
		proto = syscall.IPPROTO_IPV6
		call = syscall.IPV6_TCLASS
		errmsg = "TCLASS"
	}
	// fmt.Printf("Setting %s to 0x%x\n", errmsg, tos)

	err := syscall.SetsockoptInt(int(f.Fd()), proto, call, tos)

	if err != nil {
		WarnErrorf(err, "while setting %s to %02x", errmsg, tos)
	}

	return err
}

*/
