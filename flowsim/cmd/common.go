package cmd

// A couple of commodity functions to mimic the iperf3 cli
import (
	"errors"
	"fmt"
)

const (
	kilo float64 = 1024.0
	mega float64 = kilo * kilo
	giga float64 = mega * kilo
	tera float64 = giga * kilo
)

// Convert a string to a float

var conv = map[rune]float64{
	'\000': 1,
	'k':    kilo, 'K': kilo,
	'm': mega, 'M': mega,
	'g': giga, 'G': giga,
	't': tera, 'T': tera,
}

func utoi(s string) (int, error) {
	v, e := utof(s)
	return int(v), e
}

/*
 utof(string) (float64, error)

 Try convert the input string to a float (convert kmgtKMGT abbrevs)
 Accepts plain number
*/
func utof(s string) (float64, error) {
	var val float64
	var unit rune

	fmt.Sscanf(s, "%f%c", &val, &unit)

	//	fmt.Printf("In utof(%s), unit='%c', val=%f\n", s, unit, val)
	// Check and ignore unknown unit multiplier
	mult, ok := conv[unit]
	if ok {
		return mult * val, nil
	}

	return val, errors.New(fmt.Sprintf("Unknown multiplier '%s'", string(unit)))
}

// A dictionary to map DSCP labels to values
// CAVEAT: Multiply by 4 to get the full TOS byte value

var dscpDict = map[string]int{
	"CS0":  0,
	"CS1":  8,
	"CS2":  16,
	"CS3":  24,
	"CS4":  32,
	"CS5":  40,
	"CS6":  48,
	"CS7":  56,
	"EF":   46,
	"AF11": 10,
	"AF12": 12,
	"AF13": 14,
	"AF21": 18,
	"AF22": 20,
	"AF23": 22,
	"AF31": 26,
	"AF32": 28,
	"AF33": 30,
	"AF41": 34,
	"AF42": 36,
	"AF43": 38,
}

func Dscp(s string) (int, error) {
	if val, ok := dscpDict[s]; ok {
		return val, nil
	} else {
		_, err := fmt.Sscanf(s, "%d", &val)
		if err != nil {
			return 0, errors.New("Unknown DSCP ID ")
		}
		if val < 0 || val > 64 {
			return 0, errors.New("Value out of range [0,64)")
		}
		return val, nil
	}
}
