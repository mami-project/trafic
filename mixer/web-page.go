package mixer

import (
	"path"
	"time"
)

var defaultWebPageTmpl string = `
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
    bytes: 1246K
    report-interval-s: {{ .ReportInterval }}

server:
  at:
    - 0s
  config:
    server-port: *p
`

type WebPage struct{}

func NewWebPage() Mixer {
	return &WebPage{}
}

func (WebPage) WriteConf(baseDir string, g GlobalDesc, c FlowDesc) error {
	burstSize := float64(1246000 * 8)           // bytes: 1246K
	burstPeriod, _ := time.ParseDuration("10s") // burst every 10s

	return writeBursting(
		path.Join(baseDir, "web-page"),
		defaultWebPageTmpl,
		g, c,
		burstSize,
		burstPeriod,
	)
}

func (WebPage) Name() string {
	return "web-page"
}
