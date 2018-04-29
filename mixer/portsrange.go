package mixer

import (
	"errors"
	"fmt"
)

type PortsRange struct {
	First uint16
	Last  uint16
}

func (pr *PortsRange) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v string

	err := unmarshal(&v)
	if err != nil {
		return errors.New("unable to unmarshal to PortsRange")
	}

	_, err = fmt.Sscanf(v, "%d-%d", &pr.First, &pr.Last)
	if err != nil {
		return fmt.Errorf("unable to unmarshal %s to PortsRange: %v", v, err)
	}

	return nil
}
