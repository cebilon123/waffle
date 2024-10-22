package redirect

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"waffle/internal/domain"
)

// Handler is a struct that implements the http.Handler interface.
// It is responsible for processing incoming HTTP requests and forwarding them to the appropriate backend server
// based on the DNS resolution provided by the NameSystemProvider.
type Handler struct {
	dns domain.NameSystemProvider
}

// NewHandler creates and returns a new instance of Handler.
// It takes a NameSystemProvider as a parameter, which is used to resolve hostnames to addresses.
func NewHandler(dns domain.NameSystemProvider) *Handler {
	return &Handler{
		dns: dns,
	}
}

// Ensures that Handler implements the http.Handler interface.
var _ http.Handler = (*Handler)(nil)

// ServeHTTP processes an incoming HTTP request.
// It resolves the target address using the DNS provider based on the request's host.
// If an error occurs during resolution, it responds with a 502 Bad Gateway status.
// Otherwise, it sets up a reverse proxy to the resolved address and forwards the request,
// adjusting the request URL and headers as needed to maintain the correct context for the proxying.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	addr, err := h.dns.GetAddress(r.Host)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(addr)

	r.URL.Host = addr.Host
	r.URL.Scheme = addr.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = addr.Host
	path := r.URL.Path
	r.URL.Path = strings.TrimLeft(path, addr.Path)
	fmt.Printf("[ TinyRP ] Redirecting request to %s at %s\n", r.URL, time.Now().UTC())
	proxy.ServeHTTP(w, r)
}
