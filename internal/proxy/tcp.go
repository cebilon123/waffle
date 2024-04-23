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
	remoteAddr string
	listenAddr string

	bytesChan chan []byte
}

var _ Sender = (*TCPSender)(nil)

// NewTCPSender creates new instance of the TCPSender, which is used
// to send bytes to other remote connection.
func NewTCPSender(listenAddr, remoteAddr string) *TCPSender {
	return &TCPSender{
		remoteAddr: remoteAddr,
		listenAddr: listenAddr,
		bytesChan:  make(chan []byte),
	}
}

// Start starts the TCPSender which sends TCP bytes to
// desired host.
func (s *TCPSender) Start(ctx context.Context) error {
	rAddr, err := net.ResolveTCPAddr("tcp", s.remoteAddr)
	if err != nil {
		return fmt.Errorf("resolving remote tcp addr: '%s': %w", s.remoteAddr, err)
	}

	lAddr, err := net.ResolveTCPAddr("tcp", s.listenAddr)
	if err != nil {
		return fmt.Errorf("resolving listen tcp addr: '%s': %w", s.remoteAddr, err)
	}

	go s.startRemoteSender(ctx, rAddr, lAddr)

	return nil
}

// Accept should be called in order to accepts new TCP bytes.
// Those bytes then are send to desired TCP location.
func (s *TCPSender) Accept(rw io.ReadWriter) error {
	bytes, err := io.ReadAll(rw)
	if err != nil {
		return fmt.Errorf("read all bytes from read writer: %w", err)
	}

	s.bytesChan <- bytes

	return nil
}

// startRemoteSender is used to start process of sending incoming bytes to the host.
// It should be called in another go routine in order to not block the main routine.
func (s *TCPSender) startRemoteSender(ctx context.Context, listenAddr, remoteAddr *net.TCPAddr) {
	conn, err := net.DialTCP("tcp", listenAddr, remoteAddr)
	if err != nil {
		fmt.Printf("error dialing remote host: %s", err.Error())
	}

	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Printf("connection close failed: %s", err.Error())
		}
	}()

	for {
		select {
		case bytes, ok := <-s.bytesChan:
			if !ok {
				fmt.Printf("bytes channel has been closed, closing connection")
				return
			}

			n, err := conn.Write(bytes)
			if err != nil {
				fmt.Printf("error writing bytes to connection: %s", err.Error())
				continue
			}

			fmt.Printf("written: %d bytes to remote connection", n)
		case <-ctx.Done():
			fmt.Println("context done")
			return
		}
	}
}
