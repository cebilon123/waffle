package ratelimit

import (
	"net"
	"testing"
	"time"

	"waffle/internal/clock"
)

func TestRate_IsLimited(t *testing.T) {
	limitedUntilBefore := clock.Now().Add(time.Duration(-1) * time.Hour)
	limitedUntilAfter := clock.Now().Add(time.Hour)

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
			name: "limited until is after now, returns false",
			fields: fields{
				LimitedUntil: limitedUntilAfter,
			},
			want: false,
		},
		{
			name: "limited until is before now, returns true",
			fields: fields{
				LimitedUntil: limitedUntilBefore,
			},
			want: true,
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
