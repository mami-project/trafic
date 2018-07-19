package cmd

// A couple of commodity functions to mimic the iperf3 cli
import (
	"fmt"
)

// Convert a string to a float

func iperf3_atof (s string) float64 {
	var val float64
	var unit rune
	conv := make(map[rune] float64)
	conv['k'] = 1024.0;
	conv['K'] = conv['k']
	conv['m'] = 1024.0 * conv['k']
	conv['M'] = conv['m']
	conv['g'] = 1024.0 * conv['m']
	conv['G'] = conv['g']
	conv['t'] = 1024.0 * conv['g']
	conv['T'] = conv['t']

	fmt.Sscanf(s, "%f%c", &val, &unit)

	fmt.Printf("%s -> %f %c\n", s, val, unit)
	mult, ok := conv[unit]
	if !ok {
		mult = 1.0
	}
	return mult * val
}

func test (s string) {
	t := iperf3_atof(s)
	fmt.Printf("%s = %f bytes\n\n", s, t)
}

// func main() {
// 	test("1.8M")
// 	test("1887436")
// 	test("3G")
// 	test("1.0T")
// }
