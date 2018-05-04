package mixer

import "log"

func MixerInit() *MixerMap {
	log.Printf("initialising mixer map")

	mm := NewMixerMap()

	mm.MixerRegister(NewRealtimeAudio())
	mm.MixerRegister(NewRealtimeVideo())
	mm.MixerRegister(NewScavenger())
	mm.MixerRegister(NewGreedy())
	mm.MixerRegister(NewABRVideo())
	mm.MixerRegister(NewWebPage())

	return mm
}
