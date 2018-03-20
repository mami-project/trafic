package config

import (
	"errors"
)

// ClientConfig implements the Configurer interface for an iperf3 client
type ClientConfig struct {
	BufSize          string `yaml:"length"`
	ClientPort       uint16 `yaml:"cport"`
	ConnectTimeout   uint64 `yaml:"connect-timeout-ms"`
	DSCP             string `yaml:"dscp"`
	DisableNagle     bool   `yaml:"no-delay"`
	FlowBytes        string `yaml:"bytes"`
	FlowDuration     uint64 `yaml:"time-s"`
	FlowPackets      string `yaml:"blockcount"`
	GetServerOutput  bool   `yaml:"get-server-output"`
	MSS              uint   `yaml:"set-mss"`
	OmitLeadingSecs  uint   `yaml:"omit-s"`
	PacingTimer      string `yaml:"pacing-timer-ms"`
	ParallelFlows    uint   `yaml:"parallel"`
	RSAPubKeyFile    string `yaml:"rsa-public-key-path"`
	ReverseDir       bool   `yaml:"reverse"`
	ServerAddr       string `yaml:"server-address"`
	TargetBitrate    string `yaml:"target-bitrate"`
	ToS              uint8  `yaml:"tos"`
	UDP              bool   `yaml:"udp"`
	UDP64BitCounters bool   `yaml:"udp-counters-64bit"`
	Username         string `yaml:"username"`
	V4Only           bool   `yaml:"v4"`
	V6Only           bool   `yaml:"v6"`
	WindowSize       string `yaml:"window"`
	ZeroCopy         bool   `yaml:"zerocopy"`

	CommonConfig `yaml:",inline"`
}

// The following CLI arguments are not exposed:
// - get-server-output	force (true)

func (cfg *ClientConfig) ToArgs() ([]string, error) {
	args := []string{"--client"}

	// Make sure this results in a "sensible" iperf3 invocation
	// we want to identify errors as early as possible (i.e., before
	// the command is actually scheduled.)
	if cfg.ServerAddr == "" {
		return nil, errors.New("missing mandatory server address")
	}
	args = append(args, cfg.ServerAddr)

	if cfg.FlowDuration != 0 && cfg.FlowBytes != "" && cfg.FlowPackets != "" {
		return nil, errors.New("time-s, bytes and blockcount are mutually exclusive")
	}

	args = AppendKeyVal(args, "--length", cfg.BufSize)
	args = AppendKeyVal(args, "--cport", cfg.ClientPort)
	args = AppendKeyVal(args, "--connect-timeout", cfg.ConnectTimeout)
	args = AppendKeyVal(args, "--dscp", cfg.DSCP)
	args = AppendKey(args, "--no-delay", cfg.DisableNagle)
	args = AppendKeyVal(args, "--bytes", cfg.FlowBytes)
	args = AppendKeyVal(args, "--time", cfg.FlowDuration)
	args = AppendKeyVal(args, "--blockcount", cfg.FlowPackets)
	// args = AppendKey(args, "--get-server-output", cfg.GetServerOutput)
	args = AppendKeyVal(args, "--set-mss", cfg.MSS)
	args = AppendKeyVal(args, "--omit", cfg.OmitLeadingSecs)
	args = AppendKeyVal(args, "--pacing-timer", cfg.PacingTimer)
	args = AppendKeyVal(args, "--parallel", cfg.ParallelFlows) // to be determined
	args = AppendKeyVal(args, "--rsa-public-key-path", cfg.RSAPubKeyFile)
	args = AppendKey(args, "--reverse", cfg.ReverseDir)
	args = AppendKeyVal(args, "--bitrate", cfg.TargetBitrate)
	args = AppendKeyVal(args, "--tos", cfg.ToS)
	args = AppendKey(args, "--udp", cfg.UDP)
	args = AppendKey(args, "--udp-counters-64bit", cfg.UDP64BitCounters)
	args = AppendKeyVal(args, "--username", cfg.Username)
	args = AppendKey(args, "--version4", cfg.V4Only)
	args = AppendKey(args, "--version6", cfg.V6Only)
	args = AppendKeyVal(args, "--window", cfg.WindowSize)
	args = AppendKey(args, "--zerocopy", cfg.ZeroCopy)

	return cfg.CommonConfig.ToArgs(args)
}
