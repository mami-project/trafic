package config

import (
	"reflect"
	"testing"
)

func TestFlowClientConfig_ToArgs(t *testing.T) {
	type fields struct {
		ServerAddr       string
		BufSize          string
		DSCP             string
		ClientPort       uint16
		FlowBytes        string
		FlowIterations   uint64
		FlowInterval     uint64
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			"good stuff",
			fields{
				ClientPort:       12345,
				FlowIterations:   6,
				FlowInterval:     5,
				DSCP:             "AF12",
				FlowBytes:        "1M",
				ServerAddr:       "iperf-server",
			},
			[]string{
				"flowsim",
				"client",
				"-I",
				"iperf-server",
				"-T",
				"AF12",
				"-t",
				"5",
				"-p",
				"12345",
				"-N",
				"1M",
				"-n",
				"6",
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &FlowClientConfig{
				ServerAddr: tt.fields.ServerAddr,
				BufSize: tt.fields.BufSize,
				DSCP: tt.fields.DSCP,
				ClientPort: tt.fields.ClientPort,
				FlowBytes: tt.fields.FlowBytes,
				FlowIterations: tt.fields.FlowIterations,
				FlowInterval: tt.fields.FlowInterval,
			}
			got, err := cfg.ToArgs()
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientConfig.ToArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientConfig.ToArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
