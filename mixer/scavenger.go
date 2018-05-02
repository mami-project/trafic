package mixer

import (
	"path"
)

var defaultScavengerTmpl string = `
{{/*
  Models a number of 1Mbps (application limited), long-lived TCP streams
  TODO set LEDBAT as CC (if available)

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
    target-bitrate: 1M
    title: *l
    report-interval-s: {{ .ReportInterval }}
    parallel: *i
    reverse: true

server:
  at:
    - 0s
  config:
    server-port: *p
`

type Scavenger struct{}

func NewScavenger() Mixer {
	return &Scavenger{}
}

func (Scavenger) WriteConf(baseDir string, g GlobalDesc, c FlowDesc) error {
	// target-bitrate: 1M
	return writeFixedBitrate(
		path.Join(baseDir, "scavenger.yaml"),
		defaultScavengerTmpl,
		g, c, 1000000,
	)
}

func (Scavenger) Name() string {
	return "scavenger"
}
