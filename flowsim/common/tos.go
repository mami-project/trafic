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
	// fmt.Printf("While setting TOS to %d on %v: %v\n", tos, f, err)
	//   return err
	// }


  if ipv6 {
	err := syscall.SetsockoptInt(int(f.Fd()), syscall.IPPROTO_IPV6, syscall.IPV6_TCLASS, tos)
	if err != nil {
		fmt.Printf("While setting TCLASS to %d: %v\n", tos, err)
	}
	  return err
  }

	err := syscall.SetsockoptInt(int(f.Fd()), syscall.IPPROTO_IP, syscall.IP_TOS, tos)
	if err != nil {
		fmt.Printf("While setting TOS to %d: %v\n", tos, err)
	}

  return err
}
