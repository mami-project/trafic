package mixer

import (
	"errors"
	"fmt"
)

type Ratio struct {
	Val float64
}

func (r *Ratio) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v string

	err := unmarshal(&v)
	if err != nil {
		return errors.New("unable to unmarshal to Ratio")
	}

	var percent float64
	_, err = fmt.Sscanf(v, "%f%%", &percent)
	if err != nil {
		return fmt.Errorf("unable to unmarshal %s to Ratio: %v", v, err)
	}

	r.Val = percent / 100

	return nil
}
