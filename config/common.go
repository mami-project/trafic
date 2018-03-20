package config

// CommonConfig implements the Configurer interface for configuration keys that
// are common to iperf3 clients and servers
type CommonConfig struct {
	BindInterface  string  `yaml:"bind"`
	Debug          bool    `yaml:"debug"`
	ExchangedFile  string  `yaml:"file"`
	ForceFlush     bool    `yaml:"forceflush"`
	JSON           bool    `yaml:"json"`
	LogFile        string  `yaml:"logfile"`
	ReportFormat   string  `yaml:"format"`
	ReportInterval float32 `yaml:"report-interval-s"`
	ServerPort     uint16  `yaml:"server-port"`
	Verbose        bool    `yaml:"verbose"`
}

// The following CLI arguments are not exposed:
// - debug		suppressed
// - json		forced (true)
// - logfile	suppressed
// - format		suppressed
// - verbose	suppressed

func (cfg *CommonConfig) ToArgs(args []string) ([]string, error) {
	cargs := []string{"--json"}

	cargs = AppendKeyVal(cargs, "--bind", cfg.BindInterface)
	//	args = AppendKey(cargs, "--debug", cfg.Debug)
	cargs = AppendKeyVal(cargs, "--file", cfg.ExchangedFile)
	cargs = AppendKey(cargs, "--forceflush", cfg.ForceFlush)
	//	args = AppendKey(cargs, "--json", cfg.JSON)
	//	args = AppendKeyVal(cargs, "--logfile", cfg.LogFile)
	//	args = AppendKeyVal(cargs, "--format", cfg.ReportFormat)
	cargs = AppendKeyVal(cargs, "--interval", cfg.ReportInterval)
	cargs = AppendKeyVal(cargs, "--port", cfg.ServerPort)
	//	args = AppendKey(cargs, "--verbose", cfg.Verbose)

	return append(args, cargs...), nil
}
