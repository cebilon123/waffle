package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"waffle/internal/domain"
)

const (
	pathCertFile = ".cert/server.crt"
	pathKeyFile  = ".cert/server.key"
)

type Server struct {
	dns  domain.NameSystem
	addr string
}

func NewServer(dns domain.NameSystem, addr string) *Server {
	return &Server{
		dns:  dns,
		addr: addr,
	}
}

func (s *Server) Start() error {
	remote, _ := url.Parse("http://google.com")

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			r.Host = remote.Host
			w.Header().Set("X-Ben", "Rad")
			p.ServeHTTP(w, r)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	http.HandleFunc("/", handler(proxy))

	if err := http.ListenAndServeTLS(s.addr, pathCertFile, pathKeyFile, nil); err != nil {
		return fmt.Errorf("start reverse proxy server: %w", err)
	}

	return nil
}
