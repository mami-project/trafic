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
 * Check whether the IP destination is IPv4 or IPv6
 * and set the UDP family to 'udp4' or 'udp6'
 */
func UdpFamily(ip string) (string, error) {
	ipAddr, err := net.ResolveIPAddr("ip", ip)
	if err == nil {
		if ipAddr.IP.To4() == nil {
			return "udp6", nil
		}
		return "udp4", nil
	}
	return "", err
}
