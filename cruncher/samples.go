package cruncher

// TCPFlowSample contains a set of measures relative to a TCP flow, sampled
// in the [.Start, .End] period.
type TCPFlowSample struct {
	ID             string  `json:"id"`
	Timestamp      string  `json:"timestamp"`
	SampleDuration float64 `json:"sample-duration-s"`
	Bytes          int64   `json:"bytes"`
	Bps            float64 `json:"bps"`
	Retransmits    int     `json:"retransmits"`
	SndCwnd        int     `json:"snd-cwnd"`
	RttMs          float64 `json:"rtt-ms"`
	RttVar         float64 `json:"rtt-var"`
	Pmtu           int     `json:"pmtu"`
}

// UDPFlowSample contains a set of measures relative to a UDP flow, sampled
// in the [.Start, .End] period.
type UDPFlowSample struct {
	ID             string  `json:"id"`
	Timestamp      string  `json:"timestamp"`
	SampleDuration float64 `json:"sample-duration-s"`
	Bytes          int64   `json:"bytes"`
	Bps            float64 `json:"bps"`
	JitterMs       float64 `json:"jitter-ms"`
	LostPackets    int     `json:"lost-packets"`
	LostPercent    float64 `json:"lost-percent"`
	Packets        int     `json:"packets"`
}
