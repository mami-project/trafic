package cmd

// A couple of commodity functions to mimic the iperf3 cli
import (
	"fmt"
	"errors"
)

const (
        kilo float64 = 1024.0
        mega float64 = kilo * kilo
        giga float64 = mega * kilo
        tera float64 = giga * kilo
)

// Convert a string to a float

var conv = map[rune] float64 {
			'k': kilo,
			'K': kilo,
			'm': mega,
			'M': mega,
			'g': giga,
			'G': giga,
			't': tera,
			'T': tera,
}


func utoi(s string) (int, error) {
	v, e := utof(s)
	return int(v), e
}

/*
 utof(string) (float64, error)

 Try convert the input string to a float (convert kmgt abbrev)
 */
func utof (s string) (float64, error) {
	var val float64
	var unit rune

	fmt.Sscanf(s, "%f%c", &val, &unit)

	// Check and ignore unknown unit multiplier
	if mult, ok := conv[unit]; ok {
		return mult * val, nil
	}
	return val, errors.New(fmt.Sprintf ("Unknown multiplier '%c'", unit))
}
