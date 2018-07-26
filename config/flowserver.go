package config

import "errors"
// ServerConfig implements the Configurer interface for an iperf3 client
type FlowServerConfig struct {
	ServerAddr     string `yaml:"server"`
	ServerPort     uint16 `yaml:"sport"`
	OneOff         bool   `yaml:"one-off"`
	DSCP           string `yaml:"dscp"`
}

// The following CLI arguments are not exposed:
// - daemon		suppress

func (cfg *FlowServerConfig) ToArgs() ([]string, error) {
	args := []string{"flowsim", "server"}

	//	args = AppendKey(args, "--daemon", cfg.Daemon)

	if cfg.ServerAddr == "" {
		return nil, errors.New("missing mandatory server address")
	}
	args = AppendKeyVal(args, "-I", cfg.ServerAddr)
	args = AppendKeyVal(args, "-p", cfg.ServerPort)
	args = AppendKeyVal(args, "-T", cfg.DSCP)
	args = AppendKey(args, "-1", cfg.OneOff)

	// return cfg.CommonConfig.ToArgs(args)
	return args, nil
}
