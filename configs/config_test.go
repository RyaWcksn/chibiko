package configs

import (
	"reflect"
	"testing"
)

func TestReadFromFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name     string
		args     args
		wantConf *Config
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			args: args{
				path: "./config.json",
			},
			wantConf: &Config{
				Database: struct {
					Host     string "json:\"host\""
					Port     string "json:\"port\""
					Database string "json:\"database\""
					Password string "json:\"password\""
					Username string "json:\"username\""
				}{
					Host:     "127.0.0.1",
					Port:     "33061",
					Database: "url",
					Password: "urlpass",
					Username: "urluser",
				},
				Redis: struct {
					Host     string "json:\"host\""
					Password string "json:\"password\""
					Database string "json:\"database\""
				}{
					Host:     "127.0.0.1",
					Password: "urlredis",
					Database: "1",
				},
				Prefix: "http://127.0.0.1:8080/",
			},
			wantErr: false,
		},
		{
			name: "Cannot find env",
			args: args{
				path: "..../config.json",
			},
			wantConf: nil,
			wantErr:  true,
		},
		{
			name: "Invalid config",
			args: args{
				path: "./invalid_config.json",
			},
			wantConf: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotConf, err := ReadFromFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotConf, tt.wantConf) {
				t.Errorf("ReadFromFile() = %v, want %v", gotConf, tt.wantConf)
			}
		})
	}
}
