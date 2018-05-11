package mixer

import (
	"path"
	"time"
)

var defaultABRVideoTmpl string = `
{{/*
  Models ...

  Configuration keys:
  - label: a label added to the final report
  - server: name or address of the server side of the flow
  - report_interval: flow measures sampling timer
  - port: <generated automatically>
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
    bytes: 1.8M
    report-interval-s: {{ .ReportInterval }}

server:
  at:
    - 0s
  config:
    server-port: *p
`

type ABRVideo struct{}

func NewABRVideo() Mixer {
	return &ABRVideo{}
}

func (ABRVideo) WriteConf(baseDir string, g GlobalDesc, c FlowDesc) error {
	burstSize := float64(1800000 * 8)           // bytes: 1.8M
	burstPeriod, _ := time.ParseDuration("10s") // 10s segments

	return writeBursting(
		path.Join(baseDir, "abr-video"),
		defaultABRVideoTmpl,
		g, c,
		burstSize,
		burstPeriod,
	)
}

func (ABRVideo) Name() string {
	return "abr-video"
}
