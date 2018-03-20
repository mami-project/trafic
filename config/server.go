package config

// ServerConfig implements the Configurer interface for an iperf3 client
type ServerConfig struct {
	AuthUsersFile  string `yaml:"authorized-users-path"`
	Daemon         bool   `yaml:"daemon"`
	OneOff         bool   `yaml:"one-off"`
	PidFile        string `yaml:"pidfile"`
	RSAPrivKeyFile string `yaml:"rsa-private-key-path"`

	CommonConfig `yaml:",inline"`
}

// The following CLI arguments are not exposed:
// - daemon		suppress

func (cfg *ServerConfig) ToArgs() ([]string, error) {
	args := []string{"--server"}

	args = AppendKeyVal(args, "--authorized-users-path", cfg.AuthUsersFile)
	//	args = AppendKey(args, "--daemon", cfg.Daemon)
	args = AppendKey(args, "--one-off", cfg.OneOff)
	args = AppendKeyVal(args, "--pidfile", cfg.PidFile)
	args = AppendKeyVal(args, "--rsa-private-key-path", cfg.RSAPrivKeyFile)

	return cfg.CommonConfig.ToArgs(args)
}
