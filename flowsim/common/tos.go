package common

import (
	"fmt"
	// "net"
	"os"
	"syscall"
)


func SetTos(f *os.File, tos int,ipv6 bool) (error) {

	// This needs to go in the caller code
	// f, err := Conn.File()
	// if err != nil {
	//   fmt.Printf("While setting TOS to %d on %v: %v\n", tos, f, err)
	//   return err
	// }

	proto  := syscall.IPPROTO_IP
	call   := syscall.IP_TOS
	errmsg := "TOS"

	if ipv6 {
		proto  = syscall.IPPROTO_IPV6
		call   = syscall.IPV6_TCLASS
		errmsg = "TCLASS"
	}

	err := syscall.SetsockoptInt(int(f.Fd()), proto, call, tos)
	if err != nil {
		fmt.Printf("While setting %s to %d: %v\n", errmsg, tos, err)
	}
	return err
}
