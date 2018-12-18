package udp

import (
	"fmt"
)

func Source(ip string, port int, duration int, pps int, psize int, tos int) {
	fmt.Printf("Starting server at %s:%d for %d secs at %d pps for %d byte packets (TOS: %02x)\n",
		ip, port, duration, pps, psize, tos)
}
