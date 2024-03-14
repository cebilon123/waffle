package visualize

import (
	"context"
	"net/http"

	"nhooyr.io/websocket"
)

type Server struct {
	port       string
	visualizer *Visualizer
}

func NewServer(port string) *Server {
	return &Server{
		port:       port,
		visualizer: NewVisualizer(),
	}
}

func (s *Server) GetVisualizer() *Visualizer {
	return s.visualizer
}

func (s *Server) Start(ctx context.Context) {
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := websocket.Accept(w, r, nil)
		if err != nil {

		}
	})

	if err := http.ListenAndServe(s.port, handlerFunc); err != nil {

	}
}
