package ratelimit

import (
	"net"
	"testing"
	"time"
)

func TestRate_IsLimited(t *testing.T) {
	limitedUntilBefore := time.Now().Add(time.Duration(-1) * time.Hour)
	limitedUntilAfter := time.Now().Add(time.Hour)

	type fields struct {
		UUID         string
		IpAddress    net.IP
		LimitedUntil time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "limited until is after now, returns true",
			fields: fields{
				LimitedUntil: limitedUntilAfter,
			},
			want: true,
		},
		{
			name: "limited until is before now, returns false",
			fields: fields{
				LimitedUntil: limitedUntilBefore,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rate{
				UUID:         tt.fields.UUID,
				IpAddress:    tt.fields.IpAddress,
				LimitedUntil: tt.fields.LimitedUntil,
			}
			if got := r.IsLimited(); got != tt.want {
				t.Errorf("IsLimited() = %v, want %v", got, tt.want)
			}
		})
	}
}
