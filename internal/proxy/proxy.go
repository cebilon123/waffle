package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"waffle/internal/domain"
)

type Server struct {
	dns  domain.NameSystem
	addr string
}

func NewServer(dns domain.NameSystem, addr string) *Server {
	return &Server{dns: dns}
}

func (s *Server) Start() error {
	remote, _ := url.Parse("http://google.com")

	handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL)
			r.Host = remote.Host
			w.Header().Set("X-Ben", "Rad")
			p.ServeHTTP(w, r)
		}
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	http.HandleFunc("/", handler(proxy))
	if err := http.ListenAndServe(s.addr, nil); err != nil {
		return fmt.Errorf("start reverse proxy server: %w", err)
	}

	return nil
}
