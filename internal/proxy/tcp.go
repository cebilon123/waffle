package proxy

import (
	"fmt"
	"net"
)

type Sender interface {
	Accept() error
}

// TCPReceiver is responsible for the reading
// incoming tcp connections
type TCPReceiver struct {
	localAddr string // localAddr is an address of the proxy server
}

func NewTCPReceiver(localAddr string) *TCPReceiver {
	return &TCPReceiver{
		localAddr: localAddr,
	}
}

func (r *TCPReceiver) Run() error {
	lAddr, err := net.ResolveTCPAddr("tcp", r.localAddr)
	if err != nil {
		return fmt.Errorf("resolving local tcp addr: '%s': %w", r.localAddr, err)
	}

	listener, err := net.ListenTCP("tcp", lAddr)
	if err != nil {
		return fmt.Errorf("creating TCP listener for local address: %w", err)
	}

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Printf("failed to accept TCP connection: %s", err.Error())
		}
	}

	return nil
}

// TCPSender is responsible for the sending
// of the incoming tcp connections to
// desired TCP location.
type TCPSender struct {
	// pipes is a map of ipaddr string
	// and channel of bytes used to send
	// TCP connection to desired location
	pipes      map[string]any
	remoteAddr string
}

func NewTCPSender(remoteAddr string) *TCPSender {
	return &TCPSender{
		pipes:      make(map[string]any),
		remoteAddr: remoteAddr,
	}
}

func (s *TCPSender) Run() error {
	rAddr, err := net.ResolveTCPAddr("tcp", r.remoteAddr)
	if err != nil {
		return fmt.Errorf("resolving remote tcp addr: '%s': %w", r.remoteAddr, err)
	}
}

type senderConstruct struct {
	locationIPAddr *net.IPAddr
}
