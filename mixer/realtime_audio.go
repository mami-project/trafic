package mixer

import "path"

var defaultRealtimeAudioTmpl string = `
{{/*
  Models one direction of a real-time audio:
  64Kbps, 126 bytes of RTP + media payload

  Configuration keys:
  - label: a label added to the final report
  - server: name or address of the server side of the flow
  - port: the UDP port to use when connecting to the server
  - instances: how many instances of the flow to run in parallel
  - time: how long the flow needs to live for
  - report_interval: flow measures sampling timer
*/}}

label: &l {{ .Label }}

port: &p {{ .Port }}

instances: &i {{ .Instances }}

client:
  at:
    - 0s
  config:
    server-address: {{ .Server }}
    server-port: *p
    time-s: {{ .Time }}
    udp: true
    length: 126
    target-bitrate: 64K
    title: *l
    report-interval-s: {{ .ReportInterval }}
    parallel: *i

server:
  at:
    - 0s
  config:
    server-port: *p
    report-interval-s: {{ .ReportInterval }}
`

type RealtimeAudio struct{}

func NewRealtimeAudio() Mixer {
	return &RealtimeAudio{}
}

func (RealtimeAudio) WriteConf(baseDir string, g GlobalDesc, c FlowDesc) error {
	// target-bitrate: 64K
	return writeFixedBitrate(
		path.Join(baseDir, "realtime-audio.yaml"),
		defaultRealtimeAudioTmpl,
		g, c, 64000,
	)
}

func (RealtimeAudio) Name() string {
	return "realtime-audio"
}
