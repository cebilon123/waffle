package proxy

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"waffle/internal/certificate"
)

var (
	ciphers = []uint16{
		// TLS 1.3
		tls.TLS_AES_128_GCM_SHA256,
		tls.TLS_AES_256_GCM_SHA384,
		tls.TLS_CHACHA20_POLY1305_SHA256,

		// ECDSA is about 3 times faster than RSA on the server side.
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,

		// RSA is slower on the server side but still widely used.
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
		tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,

		// Added so all ciphers are available

		tls.TLS_RSA_WITH_RC4_128_SHA,
		tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
		tls.TLS_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
		tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA,
		tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
		tls.TLS_FALLBACK_SCSV,
	}

	http11    = "http/1.1"
	http2     = "h2"
	http3     = "h3"
	alpnProto = "acme-tls/1"
)

// Server represents an HTTP server that listens on a specific address.
// It uses a certificate provider to manage TLS certificates and an HTTP handler to process incoming requests.
type Server struct {
	addr                string
	certificateProvider certificate.Provider
	handler             http.Handler
}

// NewServer initializes a new Server instance with the given address, certificate provider, and HTTP handler.
// The address defines where the server will listen for incoming connections.
// The certificateProvider is used to manage the TLS certificates for secure connections.
// The handler (usually a redirect handler) will process HTTP requests that are received by the server.
func NewServer(
	addr string,
	certificateProvider certificate.Provider,
	redirectHandler http.Handler,
) *Server {
	return &Server{
		addr:                addr,
		certificateProvider: certificateProvider,
		handler:             redirectHandler,
	}
}

// Start initializes and starts the server with TLS configuration using the provided certificate provider.
// It first retrieves the CA certificates pool and server TLS certificate from the certificate provider.
// Then, it configures the server to use TLS 1.3, sets the allowed cipher suites, and prepares the server for
// mutual TLS authentication (client certificates are verified if provided).
//
// A TCP listener is created to listen for incoming TLS connections on the specified address.
//
// The HTTP server is configured with various timeouts (read, write, idle), a maximum header size, and a custom error logger.
// The server handles incoming requests using the provided handler.
//
// A graceful shutdown mechanism is implemented: when an interrupt signal is received (e.g., Ctrl+C),
// the server begins shutting down by closing the active listener and processing any outstanding requests.
//
// The function waits for all idle connections to be closed before returning.
//
// If there is any error while starting or serving the server, the error is returned.
func (s *Server) Start() error {
	caCertPool, err := s.certificateProvider.GetCACertificatesPool()
	if err != nil {
		return fmt.Errorf("get ca certificates pool using certificate provider: %w", err)
	}

	serverCertificate, err := s.certificateProvider.GetTLSCertificate()
	if err != nil {
		return fmt.Errorf("get tls certificate using certificate provider: %w", err)
	}

	tlsConfig := &tls.Config{
		MinVersion:   tls.VersionTLS13,
		MaxVersion:   tls.VersionTLS13,
		CipherSuites: ciphers,
		NextProtos:   []string{http3, http2, http11, alpnProto},
		Certificates: []tls.Certificate{*serverCertificate},
		ClientAuth:   tls.VerifyClientCertIfGiven,
		Rand:         rand.Reader,
		RootCAs:      caCertPool,
		ClientCAs:    caCertPool,
	}

	tcpListener, err := tls.Listen("tcp", s.addr, tlsConfig)
	if err != nil {
		return fmt.Errorf("tls tcp listener listen: %w", err)
	}

	server := &http.Server{
		Handler:           s.handler,
		Addr:              fmt.Sprintf("%s%s", "localhost", s.addr),
		ReadHeaderTimeout: 120 * time.Second,
		WriteTimeout:      120 * time.Second,
		IdleTimeout:       120 * time.Second,
		ReadTimeout:       120 * time.Second,
		TLSConfig:         tlsConfig,
		MaxHeaderBytes:    1048576,
		ErrorLog:          log.New(os.Stderr, "proxy server: ", 0),
	}

	// Implementing graceful shutdown - if an interrupt signal is received, stop listening on the port
	idleConnClosedChan := make(chan struct{})

	go func() {
		defer close(idleConnClosedChan)

		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		log.Println("Shutting down server...")
		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listener(s), or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
	}()

	if err := server.Serve(tcpListener); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("server serve: %w", err)
	}
	<-idleConnClosedChan

	return nil
}

// Shutdown is a function which will be called when gracefully shutting down the server.
func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Server shut down gracefully")
	return nil
}
