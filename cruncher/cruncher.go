package cruncher

type Cruncher interface {
	// Crunch takes an iperf3 JSON report and transforms it into a format that
	// works for a certain consumer (e.g., Telegraf / InfluxDB)
	Crunch([]byte) ([]byte, error)
}
