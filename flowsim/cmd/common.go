package cmd

// A couple of commodity functions to mimic the iperf3 cli
import (
	"fmt"
	"errors"
	"regexp"
	"strconv"
	// "log"
)

const (
        kilo float64 = 1024.0
        mega float64 = kilo * kilo
        giga float64 = mega * kilo
        tera float64 = giga * kilo
)

// Convert a string to a float

var conv = map[string] float64 {
	"" : 1.0,
	"k": kilo, // "K": kilo,
	/* "m": mega, */ "M": mega,
	/* "g": giga, */ "G": giga,
	/* "t": tera, */ "T": tera,
}


func utoi(s string) (int, error) {
	v, e := utof(s)
	return int(v), e
}

/*
 utof(string) (float64, error)

 Try convert the input string to a float (convert kmgtKMGT abbrevs)
 */
func utof (s string) (float64, error) {

	expr := regexp.MustCompile(`^([0-9]+([.][0-9]+)?)([kMGT]?)$`)
	parsed := expr.FindStringSubmatch(s)
	if len(parsed) == 4 {
		var val float64
		var mult float64
		var unit string

		val, err := strconv.ParseFloat(parsed[1], 64)
		if err != nil {
			// log.Println("ParseFloat failed!")
			return -1, err
		}
		unit = parsed[3]
		mult, ok := conv[unit]

		if ok {
			return mult * val, nil
		}
		// fmt.Printf("conv[%s] doesn't exist",unit)
		return -1, errors.New(fmt.Sprintf ("Unknown multiplier '%s'", unit))
	}
	return -1, errors.New(fmt.Sprintf("Invalid size '%s'",s))
}
