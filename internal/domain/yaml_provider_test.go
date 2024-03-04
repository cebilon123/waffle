package domain

import (
	"testing"

	"waffle/internal/config"
)

func TestYamlNameSystemProvider_GetAddress(t *testing.T) {
	type fields struct {
		cfg *config.YamlConfig
	}
	type args struct {
		host string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "there is no host registered, error is returned",
			fields: fields{
				cfg: &config.YamlConfig{},
			},
			args: args{
				host: "not.com",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "there host registered, address is returned",
			fields: fields{
				cfg: &config.YamlConfig{
					DNS: []*config.Dns{
						{
							Host:    "yes.com",
							Address: "127.0.0.1:8080",
						},
					},
				},
			},
			args: args{
				host: "yes.com",
			},
			want:    "127.0.0.1:8080",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &YamlNameSystemProvider{
				cfg: tt.fields.cfg,
			}
			got, err := y.GetAddress(tt.args.host)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}
