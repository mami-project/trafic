package cmd

// A couple of commodity functions to mimic the iperf3 cli
import (
	"fmt"
)

// Convert a string to a float

var conv map[rune] float64

func iperf3_atof (s string) float64 {
	var val float64
	var unit rune

	// Make map only once
	_, ok := conv['k']
	if !ok {
		// fmt.Println("Initialising conv map")
		conv = map[rune] float64 {
			'k': 1024.0,
			'K': 1024.0,
			'm': 1024.0 * 1024.0,
			'M': 1024.0 * 1024.0,
			'g': 1024.0 * 1024.0 * 1024.0,
			'G': 1024.0 * 1024.0 * 1024.0,
			't': 1024.0 * 1024.0 * 1024.0 * 1024,
			'T': 1024.0 * 1024.0 * 1024.0 * 1024,
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
