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
	n := iperf3_atoi(s)
	if n != int(kilo) {
		t.Errorf("Expected %s=%d and not %d", s,int(kilo),n)
	}
}

func TestConversionM(t *testing.T) {
	s := "1M"
	n := int(iperf3_atof(s))
	if n != int(mega) {
		t.Errorf("Expected %s=%d and not %d", s,int(mega), n)
	}
}

func TestConversionG(t *testing.T) {
	s := "1g"
	n := int(iperf3_atof(s))
	if n != int(giga) {
		t.Errorf("Expected %s=%d and not %d", s,int(giga), n)
	}
}
