package config

import (
	"reflect"
	"testing"
)

func TestServerConfig_ToArgs(t *testing.T) {
	type fields struct {
		AuthUsersFile  string
		Daemon         bool
		OneOff         bool
		PidFile        string
		RSAPrivKeyFile string
		CommonConfig   CommonConfig
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			"--daemon is suppressed",
			fields{
				AuthUsersFile:  "/path/to/auth/users/file",
				Daemon:         true,
				OneOff:         true,
				PidFile:        "/path/to/pid/file",
				RSAPrivKeyFile: "/path/to/rsa/private/key",
				CommonConfig:   CommonConfig{},
			},
			[]string{
				"--server",
				"--authorized-users-path",
				"/path/to/auth/users/file",
				"--one-off",
				"--pidfile",
				"/path/to/pid/file",
				"--rsa-private-key-path",
				"/path/to/rsa/private/key",
				// assume Common's implicit settings
				"--json",
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &ServerConfig{
				AuthUsersFile:  tt.fields.AuthUsersFile,
				Daemon:         tt.fields.Daemon,
				OneOff:         tt.fields.OneOff,
				PidFile:        tt.fields.PidFile,
				RSAPrivKeyFile: tt.fields.RSAPrivKeyFile,
				CommonConfig:   tt.fields.CommonConfig,
			}
			got, err := cfg.ToArgs()
			if (err != nil) != tt.wantErr {
				t.Errorf("ServerConfig.ToArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServerConfig.ToArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
