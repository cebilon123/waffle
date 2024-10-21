package ratelimit

import (
	"net"
	"time"
)

type Rate struct {
	UUID         string
	IpAddress    net.IP
	LimitedUntil time.Time
}

func (r *Rate) IsLimited() bool {
	return time.Now().Before(r.LimitedUntil)
}
