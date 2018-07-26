package config

import (
	"errors"
)

// ClientConfig implements the Configurer interface for a flowsim client
type FlowClientConfig struct {
	ServerAddr       string `yaml:"server"`
	BufSize          string `yaml:"length"`
	DSCP             string `yaml:"dscp"`
	ClientPort       uint16 `yaml:"cport"`
	FlowBytes        string `yaml:"bytes"`
	FlowIterations   uint64 `yaml:"iterations"`
	FlowInterval     uint64 `yaml:"interval"`
}

// The following CLI arguments are not exposed:
// - get-server-output  force (true)
// - reverse            force (true)

func (cfg *FlowClientConfig) ToArgs() ([]string, error) {
	args := []string{"flowsim", "client"}

	// Make sure this results in a "sensible" flowsim invocation
	// we want to identify errors as early as possible (i.e., before
	// the command is actually scheduled.)

	if cfg.ServerAddr == "" {
		return nil, errors.New("missing mandatory server address")
	}
	args = AppendKeyVal(args, "-I", cfg.ServerAddr)
	args = AppendKeyVal(args, "-T", cfg.DSCP)
	args = AppendKeyVal(args, "-t", cfg.FlowInterval)
	args = AppendKeyVal(args, "-p", cfg.ClientPort)
	args = AppendKeyVal(args, "-N", cfg.FlowBytes)
	args = AppendKeyVal(args, "-n", cfg.FlowIterations)

	// return cfg.CommonConfig.ToArgs(args)
	return args, nil
}
