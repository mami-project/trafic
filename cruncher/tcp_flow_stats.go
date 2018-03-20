package cruncher

type TCPFlowStats struct {
	Start struct {
		Connected []struct {
			Socket     int    `json:"socket"`
			LocalHost  string `json:"local_host"`
			LocalPort  int    `json:"local_port"`
			RemoteHost string `json:"remote_host"`
			RemotePort int    `json:"remote_port"`
		} `json:"connected"`
		Version    string `json:"version"`
		SystemInfo string `json:"system_info"`
		Timestamp  struct {
			Time     string `json:"time"`
			Timesecs int    `json:"timesecs"`
		} `json:"timestamp"`
		ConnectingTo struct {
			Host string `json:"host"`
			Port int    `json:"port"`
		} `json:"connecting_to"`
		Cookie        string `json:"cookie"`
		TCPMssDefault int    `json:"tcp_mss_default"`
		SockBufsize   int    `json:"sock_bufsize"`
		SndbufActual  int    `json:"sndbuf_actual"`
		RcvbufActual  int    `json:"rcvbuf_actual"`
		TestStart     struct {
			Protocol   string `json:"protocol"`
			NumStreams int    `json:"num_streams"`
			Blksize    int    `json:"blksize"`
			Omit       int    `json:"omit"`
			Duration   int    `json:"duration"`
			Bytes      int    `json:"bytes"`
			Blocks     int    `json:"blocks"`
			Reverse    int    `json:"reverse"`
			Tos        int    `json:"tos"`
		} `json:"test_start"`
	} `json:"start"`
	Intervals []struct {
		Streams []struct {
			Socket        int     `json:"socket"`
			Start         float64 `json:"start"`
			End           float64 `json:"end"`
			Seconds       float64 `json:"seconds"`
			Bytes         int     `json:"bytes"`
			BitsPerSecond float64 `json:"bits_per_second"`
			Retransmits   int     `json:"retransmits"`
			SndCwnd       int     `json:"snd_cwnd"`
			Rtt           int     `json:"rtt"`
			Rttvar        int     `json:"rttvar"`
			Pmtu          int     `json:"pmtu"`
			Omitted       bool    `json:"omitted"`
		} `json:"streams"`
		Sum struct {
			Start         float64 `json:"start"`
			End           float64 `json:"end"`
			Seconds       float64 `json:"seconds"`
			Bytes         int     `json:"bytes"`
			BitsPerSecond float64 `json:"bits_per_second"`
			Retransmits   int     `json:"retransmits"`
			Omitted       bool    `json:"omitted"`
		} `json:"sum"`
	} `json:"intervals"`
	End struct {
		Streams []struct {
			Sender struct {
				Socket        int     `json:"socket"`
				Start         float64 `json:"start"`
				End           float64 `json:"end"`
				Seconds       float64 `json:"seconds"`
				Bytes         int64   `json:"bytes"`
				BitsPerSecond float64 `json:"bits_per_second"`
				Retransmits   int     `json:"retransmits"`
				MaxSndCwnd    int     `json:"max_snd_cwnd"`
				MaxRtt        int     `json:"max_rtt"`
				MinRtt        int     `json:"min_rtt"`
				MeanRtt       int     `json:"mean_rtt"`
			} `json:"sender"`
			Receiver struct {
				Socket        int     `json:"socket"`
				Start         float64 `json:"start"`
				End           float64 `json:"end"`
				Seconds       float64 `json:"seconds"`
				Bytes         int64   `json:"bytes"`
				BitsPerSecond float64 `json:"bits_per_second"`
			} `json:"receiver"`
		} `json:"streams"`
		SumSent struct {
			Start         float64 `json:"start"`
			End           float64 `json:"end"`
			Seconds       float64 `json:"seconds"`
			Bytes         int64   `json:"bytes"`
			BitsPerSecond float64 `json:"bits_per_second"`
			Retransmits   int     `json:"retransmits"`
		} `json:"sum_sent"`
		SumReceived struct {
			Start         float64 `json:"start"`
			End           float64 `json:"end"`
			Seconds       float64 `json:"seconds"`
			Bytes         int64   `json:"bytes"`
			BitsPerSecond float64 `json:"bits_per_second"`
		} `json:"sum_received"`
		CPUUtilizationPercent struct {
			HostTotal    float64 `json:"host_total"`
			HostUser     float64 `json:"host_user"`
			HostSystem   float64 `json:"host_system"`
			RemoteTotal  float64 `json:"remote_total"`
			RemoteUser   float64 `json:"remote_user"`
			RemoteSystem float64 `json:"remote_system"`
		} `json:"cpu_utilization_percent"`
		SenderTCPCongestion   string `json:"sender_tcp_congestion"`
		ReceiverTCPCongestion string `json:"receiver_tcp_congestion"`
	} `json:"end"`
	Title            string `json:"title"`
	ServerOutputJSON struct {
		Start struct {
			Connected []struct {
				Socket     int    `json:"socket"`
				LocalHost  string `json:"local_host"`
				LocalPort  int    `json:"local_port"`
				RemoteHost string `json:"remote_host"`
				RemotePort int    `json:"remote_port"`
			} `json:"connected"`
			Version      string `json:"version"`
			SystemInfo   string `json:"system_info"`
			SockBufsize  int    `json:"sock_bufsize"`
			SndbufActual int    `json:"sndbuf_actual"`
			RcvbufActual int    `json:"rcvbuf_actual"`
			Timestamp    struct {
				Time     string `json:"time"`
				Timesecs int    `json:"timesecs"`
			} `json:"timestamp"`
			AcceptedConnection struct {
				Host string `json:"host"`
				Port int    `json:"port"`
			} `json:"accepted_connection"`
			Cookie        string `json:"cookie"`
			TCPMssDefault int    `json:"tcp_mss_default"`
			TestStart     struct {
				Protocol   string `json:"protocol"`
				NumStreams int    `json:"num_streams"`
				Blksize    int    `json:"blksize"`
				Omit       int    `json:"omit"`
				Duration   int    `json:"duration"`
				Bytes      int    `json:"bytes"`
				Blocks     int    `json:"blocks"`
				Reverse    int    `json:"reverse"`
				Tos        int    `json:"tos"`
			} `json:"test_start"`
		} `json:"start"`
		Intervals []struct {
			Streams []struct {
				Socket        int     `json:"socket"`
				Start         float64 `json:"start"`
				End           float64 `json:"end"`
				Seconds       float64 `json:"seconds"`
				Bytes         int     `json:"bytes"`
				BitsPerSecond float64 `json:"bits_per_second"`
				Omitted       bool    `json:"omitted"`
			} `json:"streams"`
			Sum struct {
				Start         float64 `json:"start"`
				End           float64 `json:"end"`
				Seconds       float64 `json:"seconds"`
				Bytes         int     `json:"bytes"`
				BitsPerSecond float64 `json:"bits_per_second"`
				Omitted       bool    `json:"omitted"`
			} `json:"sum"`
		} `json:"intervals"`
		End struct {
			Streams []struct {
				Sender struct {
					Socket        int     `json:"socket"`
					Start         float64 `json:"start"`
					End           float64 `json:"end"`
					Seconds       float64 `json:"seconds"`
					Bytes         int     `json:"bytes"`
					BitsPerSecond int     `json:"bits_per_second"`
				} `json:"sender"`
				Receiver struct {
					Socket        int     `json:"socket"`
					Start         float64 `json:"start"`
					End           float64 `json:"end"`
					Seconds       float64 `json:"seconds"`
					Bytes         int64   `json:"bytes"`
					BitsPerSecond float64 `json:"bits_per_second"`
				} `json:"receiver"`
			} `json:"streams"`
			SumSent struct {
				Start         float64 `json:"start"`
				End           float64 `json:"end"`
				Seconds       float64 `json:"seconds"`
				Bytes         int     `json:"bytes"`
				BitsPerSecond int     `json:"bits_per_second"`
			} `json:"sum_sent"`
			SumReceived struct {
				Start         float64 `json:"start"`
				End           float64 `json:"end"`
				Seconds       float64 `json:"seconds"`
				Bytes         int64   `json:"bytes"`
				BitsPerSecond float64 `json:"bits_per_second"`
			} `json:"sum_received"`
			CPUUtilizationPercent struct {
				HostTotal    float64 `json:"host_total"`
				HostUser     float64 `json:"host_user"`
				HostSystem   float64 `json:"host_system"`
				RemoteTotal  float64 `json:"remote_total"`
				RemoteUser   float64 `json:"remote_user"`
				RemoteSystem float64 `json:"remote_system"`
			} `json:"cpu_utilization_percent"`
			SenderTCPCongestion   string `json:"sender_tcp_congestion"`
			ReceiverTCPCongestion string `json:"receiver_tcp_congestion"`
		} `json:"end"`
	} `json:"server_output_json"`
}
