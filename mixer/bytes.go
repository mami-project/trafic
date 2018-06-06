package mixer

import (
	"errors"
	"fmt"

	"github.com/alecthomas/units"
)

type Bytes struct {
	Val units.MetricBytes
}

func (b *Bytes) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v string

	err := unmarshal(&v)
	if err != nil {
		return errors.New("unable to unmarshal to Bytes")
	}

	b.Val, err = units.ParseMetricBytes(v)
	if err != nil {
		return fmt.Errorf("unable to unmarshal %s to Bytes: %v", v, err)
	}

	return nil
}
