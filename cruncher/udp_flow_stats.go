package cruncher

type UDPFlowStats struct {
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
		Cookie       string `json:"cookie"`
		SockBufsize  int    `json:"sock_bufsize"`
		SndbufActual int    `json:"sndbuf_actual"`
		RcvbufActual int    `json:"rcvbuf_actual"`
		TestStart    struct {
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
			Packets       int     `json:"packets"`
			Omitted       bool    `json:"omitted"`
		} `json:"streams"`
		Sum struct {
			Start         float64 `json:"start"`
			End           float64 `json:"end"`
			Seconds       float64 `json:"seconds"`
			Bytes         int     `json:"bytes"`
			BitsPerSecond float64 `json:"bits_per_second"`
			Packets       int     `json:"packets"`
			Omitted       bool    `json:"omitted"`
		} `json:"sum"`
	} `json:"intervals"`
	End struct {
		Streams []struct {
			UDP struct {
				Socket        int     `json:"socket"`
				Start         float64 `json:"start"`
				End           float64 `json:"end"`
				Seconds       float64 `json:"seconds"`
				Bytes         int     `json:"bytes"`
				BitsPerSecond float64 `json:"bits_per_second"`
				JitterMs      float64 `json:"jitter_ms"`
				LostPackets   int     `json:"lost_packets"`
				Packets       int     `json:"packets"`
				LostPercent   float64 `json:"lost_percent"`
				OutOfOrder    int     `json:"out_of_order"`
			} `json:"udp"`
		} `json:"streams"`
		Sum struct {
			Start         float64 `json:"start"`
			End           float64 `json:"end"`
			Seconds       float64 `json:"seconds"`
			Bytes         int     `json:"bytes"`
			BitsPerSecond float64 `json:"bits_per_second"`
			JitterMs      float64 `json:"jitter_ms"`
			LostPackets   int     `json:"lost_packets"`
			Packets       int     `json:"packets"`
			LostPercent   float64 `json:"lost_percent"`
		} `json:"sum"`
		CPUUtilizationPercent struct {
			HostTotal    float64 `json:"host_total"`
			HostUser     float64 `json:"host_user"`
			HostSystem   float64 `json:"host_system"`
			RemoteTotal  float64 `json:"remote_total"`
			RemoteUser   float64 `json:"remote_user"`
			RemoteSystem float64 `json:"remote_system"`
		} `json:"cpu_utilization_percent"`
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
			Version    string `json:"version"`
			SystemInfo string `json:"system_info"`
			Timestamp  struct {
				Time     string `json:"time"`
				Timesecs int    `json:"timesecs"`
			} `json:"timestamp"`
			AcceptedConnection struct {
				Host string `json:"host"`
				Port int    `json:"port"`
			} `json:"accepted_connection"`
			Cookie       string `json:"cookie"`
			SockBufsize  int    `json:"sock_bufsize"`
			SndbufActual int    `json:"sndbuf_actual"`
			RcvbufActual int    `json:"rcvbuf_actual"`
			TestStart    struct {
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
				JitterMs      float64 `json:"jitter_ms"`
				LostPackets   int     `json:"lost_packets"`
				Packets       int     `json:"packets"`
				LostPercent   float64 `json:"lost_percent"`
				Omitted       bool    `json:"omitted"`
			} `json:"streams"`
			Sum struct {
				Start         float64 `json:"start"`
				End           float64 `json:"end"`
				Seconds       float64 `json:"seconds"`
				Bytes         int     `json:"bytes"`
				BitsPerSecond float64 `json:"bits_per_second"`
				JitterMs      float64 `json:"jitter_ms"`
				LostPackets   int     `json:"lost_packets"`
				Packets       int     `json:"packets"`
				LostPercent   float64 `json:"lost_percent"`
				Omitted       bool    `json:"omitted"`
			} `json:"sum"`
		} `json:"intervals"`
		End struct {
			Streams []struct {
				UDP struct {
					Socket        int     `json:"socket"`
					Start         float64 `json:"start"`
					End           float64 `json:"end"`
					Seconds       float64 `json:"seconds"`
					Bytes         int     `json:"bytes"`
					BitsPerSecond int     `json:"bits_per_second"`
					JitterMs      float64 `json:"jitter_ms"`
					LostPackets   int     `json:"lost_packets"`
					Packets       int     `json:"packets"`
					LostPercent   float64 `json:"lost_percent"`
					OutOfOrder    int     `json:"out_of_order"`
				} `json:"udp"`
			} `json:"streams"`
			Sum struct {
				Start         float64 `json:"start"`
				End           float64 `json:"end"`
				Seconds       float64 `json:"seconds"`
				Bytes         int     `json:"bytes"`
				BitsPerSecond int     `json:"bits_per_second"`
				JitterMs      float64 `json:"jitter_ms"`
				LostPackets   int     `json:"lost_packets"`
				Packets       int     `json:"packets"`
				LostPercent   float64 `json:"lost_percent"`
			} `json:"sum"`
			CPUUtilizationPercent struct {
				HostTotal    float64 `json:"host_total"`
				HostUser     float64 `json:"host_user"`
				HostSystem   float64 `json:"host_system"`
				RemoteTotal  float64 `json:"remote_total"`
				RemoteUser   float64 `json:"remote_user"`
				RemoteSystem float64 `json:"remote_system"`
			} `json:"cpu_utilization_percent"`
		} `json:"end"`
	} `json:"server_output_json"`
}
