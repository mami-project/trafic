package mixer

import "log"

func MixerInit() *MixerMap {
	log.Printf("initialise mixer map")

	mm := NewMixerMap()

	mm.MixerRegister(NewRealtimeAudio())
	mm.MixerRegister(NewRealtimeVideo())
	mm.MixerRegister(NewScavenger())
	mm.MixerRegister(NewGreedy())

	return mm
}
