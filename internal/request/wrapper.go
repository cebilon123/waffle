package request

import (
	"net"
	"net/http"
)

// Wrapper is used to wrap request, in order to not pass it every time.
type Wrapper struct {
	Request *http.Request
	// IPAddress is a real ip address of the client. The difference between this one and request remoteAddr is
	// that it can be read from the headers (i.e. if request was forwarded by some kind of proxy).
	IPAddress *net.IP
}

func NewRequestWrapper(r *http.Request, ipAddress *net.IP) *Wrapper {
	return &Wrapper{
		Request:   r,
		IPAddress: ipAddress,
	}
}
