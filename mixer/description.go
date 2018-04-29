package mixer

import (
	"io/ioutil"
	"log"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type FlowDesc struct {
	Kind             string            `yaml:"kind"`
	PercentBandwidth Ratio             `yaml:"percent-bandwidth"`
	PortsRange       PortsRange        `yaml:"ports-range"`
	Template         string            `yaml:"template"`
	Props            map[string]string `yaml:"props"`
}

type GlobalDesc struct {
	TotalBandwidth Bytes         `yaml:"total-bandwidth"`
	TotalTime      time.Duration `yaml:"total-time"`
	ReportInterval time.Duration `yaml:"report-interval"`
}

type Description struct {
	Global GlobalDesc `yaml:",inline"`
	Flows  []FlowDesc `yaml:"flows"`
}

func NewDescriptionFromYaml(buf []byte) (*Description, error) {
	var t Description

	err := yaml.Unmarshal(buf, &t)
	if err != nil {
		return nil, err
	}

	log.Printf("description loaded")

	return &t, nil
}

func NewDescriptionFromFile(path string) (*Description, error) {
	log.Printf("loading description from %s", path)

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return NewDescriptionFromYaml(buf)
}
