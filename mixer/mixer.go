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

type MixerMap struct {
	m map[string]Mixer
}

func NewMixerMap() *MixerMap {
	m := MixerMap{make(map[string]Mixer)}

	m.MixerRegister(NewRealtimeAudio())
	m.MixerRegister(NewRealtimeVideo())
	m.MixerRegister(NewScavenger())
	m.MixerRegister(NewGreedy())
	m.MixerRegister(NewABRVideo())
	m.MixerRegister(NewWebPage())

	return &m
}

func (m *MixerMap) MixerRegister(mixer Mixer) {
	id := mixer.Name()

	log.Printf("adding %s to the mixer map", id)

	m.m[id] = mixer
}

func (m *MixerMap) LookupMixer(id string) (*Mixer, error) {
	log.Printf("looking up flow mixer for %s", id)

	mixer, ok := m.m[id]
	if !ok {
		return nil, fmt.Errorf("no mixer registered for %s", id)
	}

	return &mixer, nil
}
