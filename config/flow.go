package config

import (
	"io/ioutil"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type FlowClient struct {
	At     []time.Duration `yaml:"at"` // time.ParseDuration
	Config ClientConfig    `yaml:"config"`
}

type FlowServer struct {
	At     []time.Duration `yaml:"at"` // time.ParseDuration
	Config ServerConfig    `yaml:"config"`
}

type FlowConfig struct {
	Label  string     `yaml:"label"`
	Client FlowClient `yaml:"client"`
	Server FlowServer `yaml:"server"`
}

func NewFlowConfigFromYaml(buf []byte) (*FlowConfig, error) {
	var t FlowConfig

	err := yaml.Unmarshal(buf, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func NewFlowConfigFromFile(path string) (*FlowConfig, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return NewFlowConfigFromYaml(buf)
}
