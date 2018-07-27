package cmd

// A couple of commodity functions to mimic the iperf3 cli
import (
	"fmt"
)

const (
        kilo float64 = 1024.0
        mega float64 = kilo * kilo
        giga float64 = mega * kilo
        tera float64 = giga * kilo
)

// Convert a string to a float

var conv map[rune] float64

func iperf3_atoi(s string) int {
		return int(iperf3_atof(s))
}

func iperf3_atof (s string) float64 {
	var val float64
	var unit rune

	// Make map only once
	_, ok := conv['k']
	if !ok {
		// fmt.Println("Initialising conv map")
		conv = map[rune] float64 {
			'k': kilo,
			'K': kilo,
			'm': mega,
			'M': mega,
			'g': giga,
			'G': giga,
			't': tera,
			'T': tera,
		}
	}
	fmt.Sscanf(s, "%f%c", &val, &unit)

	// Check and ignore unknown unit multiplier
	mult, ok := conv[unit]
	if !ok {
		mult = 1.0
	}
	return mult * val
}
