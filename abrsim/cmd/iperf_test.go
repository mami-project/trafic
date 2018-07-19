package cmd

// A couple of commodity functions to mimic the iperf3 cli
import (
	"testing"
)

func TestConversionNoUnit(t *testing.T) {
	s := "1025"
	n := int(iperf3_atof(s))
	if n != 1025 {
		t.Errorf("Expected %s=1025 and not %d", s, n)
	}
}

func TestConversionK(t *testing.T) {
	s := "1K"
	n := int(iperf3_atof(s))
	if n != 1024 {
		t.Errorf("Expected %s=1024 and not %d", s,n)
	}
}

func TestConversionM(t *testing.T) {
	s := "1M"
	n := int(iperf3_atof(s))
	if n != 1024 * 1024 {
		t.Errorf("Expected %s=1024*1024 and not %d", s,n)
	}
}

func TestConversionG(t *testing.T) {
	s := "1g"
	n := int(iperf3_atof(s))
	if n != 1024 * 1024 * 1024 {
		t.Errorf("Expected %s=1024*1024*1024 and not %d", s,n)
	}
}
