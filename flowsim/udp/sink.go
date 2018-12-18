package udp

import (
	"fmt"
)

func Sink(ip string, port int) {
	fmt.Printf("Starting UDP sink at %s:%d\n", ip, port)
}
