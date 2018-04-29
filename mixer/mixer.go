package mixer

import (
	"fmt"
	"log"
)

type Mixer interface {
	// WriteConf writes one or more configuration files for the given flow
	WriteConf(string, GlobalDesc, FlowDesc) error

	// Name is used by a Mixer to provide its (unique) name
	Name() string
}

type MixerMap = map[string]Mixer

var mgm MixerMap

func init() {
	log.Printf("initialise mixer map")

	mgm = make(MixerMap)
}

func MixerRegister(mixer Mixer) {
	id := mixer.Name()

	log.Printf("adding %s to the mixer map", id)

	mgm[id] = mixer
}

func LookupMixer(id string) (*Mixer, error) {
	log.Printf("looking up flow mixer for %s", id)

	mixer, ok := mgm[id]
	if !ok {
		return nil, fmt.Errorf("no mixer registered for %s", id)
	}

	return &mixer, nil
}
