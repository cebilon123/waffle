package ratelimit

import (
	"net"
	"time"

	"waffle/internal/clock"
)

type Rate struct {
	UUID         string
	IpAddress    net.IP
	LimitedUntil time.Time
}

func (r *Rate) IsLimited() bool {
	return clock.Now().Before(r.LimitedUntil)
}
