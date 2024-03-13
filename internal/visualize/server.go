package visualize

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}
