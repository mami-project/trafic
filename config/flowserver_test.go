package config

import (
	"reflect"
	"testing"
)

func TestFlowServerConfig_ToArgs(t *testing.T) {
	type fields struct {
		ServerAddr     string
		ServerPort     uint16
		OneOff         bool
		DSCP           string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			"flowsim server test",
			fields{
				ServerAddr: "iperf-server",
				ServerPort: 12345,
				OneOff:     true,
				DSCP:       "AF12",
			},
			[]string{
				"flowsim",
				"server",
				"-I", "iperf-server",
				"-p", "12345",
				"-T", "AF12",
				"-1",
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &FlowServerConfig{
				ServerAddr: tt.fields.ServerAddr,
				ServerPort: tt.fields.ServerPort,
				OneOff:     tt.fields.OneOff,
				DSCP:       tt.fields.DSCP,
			}
			got, err := cfg.ToArgs()
			if (err != nil) != tt.wantErr {
				t.Errorf("FlowServerConfig.ToArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlowServerConfig.ToArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
