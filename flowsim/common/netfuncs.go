package common

import (
	"fmt"
	// 	"os"
	// 	"syscall"
	"errors"
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

func FirstIP(host string, ipv6 bool) (string, error) {

	ips, err := net.LookupIP(host)
	if err == nil {
		for _, ip := range ips {
			if ip.To4() == nil {
				if ipv6 {
					return ip.String(), nil
				}
			} else {
				if ipv6 == false {
					return ip.String(), nil
				}
			}
		}
		if ipv6 {
			return "", errors.New(fmt.Sprintf("Couldn't find IPv6 address for %s\n", host))
		}
		return "", errors.New(fmt.Sprintf("Couldn't find IPv4 address for %s\n", host))
	}
	return "", err
}
