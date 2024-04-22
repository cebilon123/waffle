package proxy

import (
	"context"
	"fmt"
	"io"
	"net"
)

type Sender interface {
	Accept(rw io.ReadWriter) error
}

// TCPReceiver is responsible for the reading
// incoming tcp connections
type TCPReceiver struct {
	localAddr string // localAddr is an address of the proxy server
	sender    Sender // sender is a sender where tcp bytes are supposed to be redirected
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

			continue
		}

		if err := r.sender.Accept(conn); err != nil {
			fmt.Printf("failed to accept connection with sender: %s", err.Error())
		}
	}

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

	bytesChan chan []byte
}

var _ Sender = (*TCPSender)(nil)

func NewTCPSender(listenAddr, remoteAddr string) *TCPSender {
	return &TCPSender{
		pipes:      make(map[string]any),
		remoteAddr: remoteAddr,
		bytesChan:  make(chan []byte),
	}
}

func (s *TCPSender) Start(ctx context.Context) error {
	rAddr, err := net.ResolveTCPAddr("tcp", s.remoteAddr)
	if err != nil {
		return fmt.Errorf("resolving remote tcp addr: '%s': %w", s.remoteAddr, err)
	}

	go s.startRemoteSender(ctx, rAddr)

	return nil
}

func (s *TCPSender) Accept(rw io.ReadWriter) error {
	bytes, err := io.ReadAll(rw)
	if err != nil {
		return fmt.Errorf("read all bytes from read writer: %w", err)
	}

	s.bytesChan <- bytes

	return nil
}

func (s *TCPSender) startRemoteSender(ctx context.Context, remoteAddr *net.TCPAddr) {
	net.DialTCP("tcp")
	for bytes := range s.bytesChan {

	}
}
