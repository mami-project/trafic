package config

// A Configurer implements the conversion from an external description format
// (e.g., yaml) to a bunch of command line arguments suitable for iperf3,
// either client or server mode.
type Configurer interface {
	ToArgs() ([]string, error)
}
