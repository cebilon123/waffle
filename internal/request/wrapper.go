package request

import (
	"net"
	"net/http"
)

type Wrapper struct {
	Request   *http.Request
	IPAddress *net.IP
}

func NewRequestWrapper(r *http.Request, ipAddress *net.IP) *Wrapper {
	return &Wrapper{
		Request:   r,
		IPAddress: ipAddress,
	}
}
