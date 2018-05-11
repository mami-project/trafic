package mixer

import (
	"path"
	"time"
)

var defaultGreedyTmpl string = `
{{/*
  Models a greedy (not application limited) TCP flow that wants to download a
  25MB file every 30s

  Configuration keys:
  - label: a label added to the final report
  - server: name or address of the server side of the flow
  - port: the UDP port to use when connecting to the server
  - report_interval: flow measures sampling timer
*/}}

label: &l {{ .Label }}

port: &p {{ .Port }}

client:
  at:
{{- range .At}}
    - {{ . -}}s
{{end}}
  config:
    server-address: {{ .Server }}
    server-port: *p
    title: *l
    report-interval-s: {{ .ReportInterval }}
    bytes: 25M

server:
  at:
    - 0s
  config:
    server-port: *p
`

type Greedy struct{}

func NewGreedy() Mixer {
	return &Greedy{}
}

func (Greedy) WriteConf(baseDir string, g GlobalDesc, c FlowDesc) error {
	burstSize := float64(25000000 * 8) // bytes: 25M
	burstPeriod, _ := time.ParseDuration("30s")

	return writeBursting(
		path.Join(baseDir, "greedy-tcp"),
		defaultGreedyTmpl,
		g, c,
		burstSize,
		burstPeriod,
	)
}

func (Greedy) Name() string {
	return "greedy"
}
