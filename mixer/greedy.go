package mixer

import (
	"fmt"
	"path"
)

var defaultGreedyTmpl string = `
{{/*
  Models a greedy (not application limited) TCP flow that wants to download a
  25MB file at regular intervals

  Configuration keys:
  - label: a label added to the final report
  - server: name or address of the server side of the flow
  - port: the UDP port to use when connecting to the server
  - report_interval: flow measures sampling timer
*/}}

label: &l {{ .label }}

port: &p {{ .port }}

client:
  at:
{{range .}}
    - {{.at -}}
{{end}}
  config:
    server-address: {{ .port }}
    server-port: *p
    title: *l
    report-interval-s: {{ .report_interval }}
    bytes: 25M
    reverse: true

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
	outFile := path.Join(baseDir, "greedy-tcp.yaml")

	return fmt.Errorf("TODO dump greedy to %s", outFile)
}

func (Greedy) Name() string {
	return "greedy"
}
