package redirect

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
	
	"waffle/internal/domain"
)

type Handler struct {
	dns domain.NameSystemProvider
}

func NewHandler(dns domain.NameSystemProvider) *Handler {
	return &Handler{
		dns: dns,
	}
}

var _ http.Handler = (*Handler)(nil)

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
	//trim reverseProxyRoutePrefix
	path := r.URL.Path
	r.URL.Path = strings.TrimLeft(path, addr.Path)
	fmt.Printf("[ TinyRP ] Redirecting request to %s at %s\n", r.URL, time.Now().UTC())
	proxy.ServeHTTP(w, r)
}
