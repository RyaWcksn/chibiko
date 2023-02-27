package configs

import (
	"testing"
)

func TestReadFromFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			args: args{
				path: "./config.json",
			},
			wantErr: false,
		},
		{
			name: "Cannot find env",
			args: args{
				path: "..../config.json",
			},
			wantErr: true,
		},
		{
			name: "Invalid config",
			args: args{
				path: "./invalid_config.json",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ReadFromFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
