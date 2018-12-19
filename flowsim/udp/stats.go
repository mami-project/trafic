package udp

import (
	"fmt"
	"net"
	// "encoding/json"
)

type Stats struct {
	lastdelay  int
	mdelay     int
	mjitter    int
	samples    int

	lastsample int
	loss       int
}

// func NewStats() *Stats {
// 	var newStats Stats

// 	newStats.samples = 0
// 	newStats.mdelay = 0
// 	newStats.mjitter = 0

// 	newStats.lastsample = 0
// 	newStats.loss = 0

// 	return &newStats
// }

func AddSample(stats *Stats, delay int, nsample int) {
	diff := nsample - stats.lastsample
	if diff != 1  {
		stats.loss += diff - 1
		// fmt.Printf(" nsample: %d , last %d\n",nsample,stats.lastsample)
	}
	stats.lastsample = nsample

	stats.mdelay += delay
	if stats.samples > 1 {
		if stats.lastdelay > delay {
			stats.mjitter += stats.lastdelay - delay
		} else {
			stats.mjitter += delay - stats.lastdelay
		}
	}
	stats.samples ++
	stats.lastdelay = delay
}

func PrintStats(addr *net.UDPAddr, stats *Stats, unit string) {
	fmt.Printf(" { \"%v\" : {\n", addr)
	fmt.Printf("   \"Delay\" :  \"%6.2f %s\",\n", float64(stats.mdelay) / float64(stats.samples), unit)
	fmt.Printf("   \"Jitter\" : \"%6.2f %s\",\n", float64(stats.mjitter) / float64(stats.samples - 1), unit)
	fmt.Printf("   \"Loss\" : \"%d\",\n", stats.loss)
	fmt.Printf("   \"Samples\" : \"%d\"\n", stats.samples)
	fmt.Printf("\n   }\n  }\n")
}
