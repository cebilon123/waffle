package request

import (
	"fmt"
	"net"
	"net/http"
	"reflect"
	"testing"
)

func TestGetRealIPAddress(t *testing.T) {
	defaultAddress := net.ParseIP("127.0.0.1")

	type args struct {
		r http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    net.IP
		wantErr bool
	}{
		{
			name: "no header, ip address should be local",
			args: args{
				r: http.Request{
					RemoteAddr: "127.0.0.1",
				},
			},
			want: net.ParseIP("127.0.0.1"),
		},
		{
			name: "no header, no remote addr, error returned",
			args: args{
				r: http.Request{},
			},
			wantErr: true,
		},
		{
			name: "header: x-client-ip, should return address",
			args: args{
				r: http.Request{
					Header: map[string][]string{
						HeaderXClientIP: {
							defaultAddress.String(),
						},
					},
				},
			},
			want: defaultAddress,
		},
		{
			name: "header: x-forwarded-for, should return first address",
			args: args{
				r: http.Request{
					Header: map[string][]string{
						HeaderXForwardedFor: {
							fmt.Sprintf("%s,%s, %s", defaultAddress.String(), "192.168.100.1", "192.127.32.1"),
						},
					},
				},
			},
			want: defaultAddress,
		},
		{
			name: "header: x-Real-Ip, should return address",
			args: args{
				r: http.Request{
					Header: map[string][]string{
						HeaderXRealIP: {
							defaultAddress.String(),
						},
					},
				},
			},
			want: defaultAddress,
		},
		{
			name: "header: x-cluster-client-ip, should return address",
			args: args{
				r: http.Request{
					Header: map[string][]string{
						HeaderXClusterClientIP: {
							defaultAddress.String(),
						},
					},
				},
			},
			want: defaultAddress,
		},
		{
			name: "header: True-Client-Ip, should return address",
			args: args{
				r: http.Request{
					Header: map[string][]string{
						HeaderTrueClientIP: {
							defaultAddress.String(),
						},
					},
				},
			},
			want: defaultAddress,
		},
		{
			name: "header: Fastly-Client-Ip, should return address",
			args: args{
				r: http.Request{
					Header: map[string][]string{
						HeaderFastlyClientIP: {
							defaultAddress.String(),
						},
					},
				},
			},
			want: defaultAddress,
		},
		{
			name: "header: Cf-Connecting-Ip, should return address",
			args: args{
				r: http.Request{
					Header: map[string][]string{
						HeaderCFConnectingIP: {
							defaultAddress.String(),
						},
					},
				},
			},
			want: defaultAddress,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRealIPAddress(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRealIPAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRealIPAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}
