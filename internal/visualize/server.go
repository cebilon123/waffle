package visualize

import (
	"context"
	"log"
	"net/http"

	"nhooyr.io/websocket"
)

// Server represents a web server that listens on a specific port and includes a visualizer component.
// The visualizer is used to analyze or display certain data related to the server's functionality.
type Server struct {
	port       string
	visualizer *Visualizer
}

// NewServer initializes and returns a new instance of the Server struct.
// It takes a port as input to define where the server will listen for incoming connections.
// A new Visualizer is also created and attached to the server during initialization.
func NewServer(port string) *Server {
	return &Server{
		port:       port,
		visualizer: NewVisualizer(),
	}
}

// GetVisualizer returns the Visualizer instance associated with the server.
// This allows other components to access and use the visualizer for data visualization.
func (s *Server) GetVisualizer() *Visualizer {
	return s.visualizer
}

// Start begins the server's operation, listening on the specified port.
// It defines an HTTP handler that attempts to upgrade incoming HTTP connections to WebSocket connections.
// If a connection upgrade fails, the error is silently handled.
// The server listens for incoming requests and serves them using the handler.
// If the server fails to start, an error is returned and handled.
func (s *Server) Start(ctx context.Context) {
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := websocket.Accept(w, r, nil)
		if err != nil {
			log.Fatalln("Accept error:", err)
		}
	})

	if err := http.ListenAndServe(s.port, handlerFunc); err != nil {
		log.Fatalln("Error while listening and serving visualizer:", err)
	}
}
