package udp

import (
	"fmt"
	// TODO, ftmb all is handled with print
	// "encoding/json"
)

type Stats struct {
	lastdelay  int
	mdelay     int
	mjitter    int
	samples    int

	lastsample int
	loss       int

	reorder    int
}

// stats:    structure holding the statistics
// delay:    packet.timestamp - now()
// nsample:  packet.counter

func AddSample(stats *Stats, delay int, nsample int) *Stats {
	fmt.Printf("before: stats = %v\n", stats)
	diff := nsample - stats.lastsample
	if diff >= 0 {
		// if diff == 0 ==> repeated packet??
		if diff > 1  {
			stats.loss += diff - 1
		} // else diff == 1 ==> OK

		stats.lastsample = nsample
	} else {
		stats.loss --
		stats.reorder ++
	}

	stats.mdelay += delay

	// The jitter measure in the presence of reordering is not 100% accurate
	// The accurate way would be to keep a time-ordered vector of delays

	if stats.samples > 1 {
		if stats.lastdelay > delay {
			stats.mjitter += stats.lastdelay - delay
		} else {
			stats.mjitter += delay - stats.lastdelay
		}
	}
	stats.samples ++
	stats.lastdelay = delay
	fmt.Printf("after: stats = %v\n", stats)
	return stats
}

//
// TODO Generate this using the JSON libraries
//
func PrintStats(addr string, stats *Stats, unit string) {
	fmt.Printf(" { \"%s\" : {\n", addr)
	fmt.Printf("   \"Delay\" :  \"%6.2f %s\",\n", float64(stats.mdelay) / float64(stats.samples), unit)
	fmt.Printf("   \"Jitter\" : \"%6.2f %s\",\n", float64(stats.mjitter) / float64(stats.samples - 1), unit)
	fmt.Printf("   \"Loss\" : \"%d\",\n", stats.loss)
	fmt.Printf("   \"Reorder\" : \"%d\",\n", stats.reorder )
	fmt.Printf("   \"Samples\" : \"%d\"\n", stats.samples)
	fmt.Printf("\n   }\n  }\n")
}
