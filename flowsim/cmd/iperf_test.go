package cmd

// A couple of commodity functions to mimic the iperf3 cli
import (
	"testing"
	"fmt"
)

func TestConversionNoUnit(t *testing.T) {
	s := "1025"
	n, _ := utof(s)
	if int(n) != 1025 {
		t.Errorf("Expected %s=1025 and not %d", s, int(n))
	}
}

func TestConversionK(t *testing.T) {
	s := "1k"
	n, _ := utoi(s)
	if n != int(kilo) {
		t.Errorf("Expected %s=%d and not %d", s,int(kilo),n)
	}
}

func TestConversionM(t *testing.T) {
	s := "1M"
	n, _ := utof(s)
	if n != mega {
		t.Errorf("Expected %s=%d and not %d", s,int(mega), int(n))
	}
}

func TestConversionG(t *testing.T) {
	s := "1G"
	n, _ := utof(s)
	if int(n) != int(giga) {
		t.Errorf("Expected %s=%d and not %f", s,int(giga), n)
	}
}

func TestConversionError(t *testing.T) {
	// '1P' passes, '1p' doesnot ???
	for _, c := range "abcdfhijlnoPqrsuvwxyz" {
		s := fmt.Sprintf("1%c", c)
		n, err := utof(s)
		if err == nil {
			t.Errorf("Expected %s would yield an error", s)
		}
		if n != -1.0 {
			t.Errorf("Expected %s would yield a value of 1.0 and not %f", s, n)
		}
	}
}
