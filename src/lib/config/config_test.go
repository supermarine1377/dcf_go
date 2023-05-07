package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name     string
		IsSetEnv bool
		Env      map[string]string
		want     *Config
		wantErr  bool
	}{
		{
			name:     "1st",
			IsSetEnv: false,
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "2nd",
			IsSetEnv: true,
			Env:      map[string]string{"FMP_API_KEY": ""},
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "3rd",
			IsSetEnv: true,
			Env:      map[string]string{"FMP_API_KEY": "api_key"},
			want: &Config{
				FMPAPIKey: "api_key",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.IsSetEnv {
				for k, v := range tt.Env {
					t.Setenv(k, v)
				}
			}
			got, err := NewConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("unexpected diff: %s", diff)
			}
		})
	}
}
