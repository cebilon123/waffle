package server

import (
	"context"
	"embed"
	"log"
	"os"
	"os/signal"
	"time"
	"waffle/internal/visualize"

	"waffle/internal/certificate"
	"waffle/internal/config"
	"waffle/internal/domain"
	"waffle/internal/proxy"
	"waffle/internal/ratelimit"
	"waffle/internal/redirect"
	"waffle/internal/waf"
	"waffle/internal/waf/guard"
)

// Run initializes and starts the Waffle Proxy server with the provided context, configuration, and embedded certificates.
//
// It first sets up signal handling to allow graceful shutdown on receiving an interrupt signal.
//
// The function loads environment-specific configurations, then parses the provided YAML configuration
// to initialize a DNS provider for managing domain names.
//
// Next, it sets up the certificate provider using locally embedded certificates,
// loading custom CA certificates, certificate PEM blocks, and key PEM blocks.
//
// A defense coordinator is initialized to handle security measures like XSS protection,
// along with an in-memory rate limiter to control the number of requests allowed per a given time window (5 minutes).
//
// Additionally, a server for visualizing traffic is set up on port :8081.
//
// The WAF (Web Application Firewall) handler is constructed using a redirect handler, the defender (for security),
// the rate limiter, and a visualizer from the visualization server.
//
// Finally, the main proxy server is started on port :8080, with the guard handler and certificate provider.
// If the proxy server fails to start, the function logs a fatal error.
//
// The function returns nil upon normal completion.
func Run(ctx context.Context, yamlConfigBytes []byte, certificates embed.FS) error {
	_, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	_, err := config.LoadEnvironmentConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	yamlCfg, err := config.NewYamlConfig(yamlConfigBytes)
	if err != nil {
		log.Fatal(err.Error())
	}

	yamlDnsProvider := domain.NewYamlNameSystemProvider(yamlCfg)

	certificateProvider := certificate.NewLocalCertificatesProvider(
		loadLocalCustomCACerts(certificates),
		loadLocalCertPEMBlock(certificates),
		loadLocalKeyPEMBlock(certificates),
	)

	defender := guard.NewDefenseCoordinator([]guard.Defender{&guard.XSS{}})

	limiter := ratelimit.NewInMemoryLimiter(time.Minute * 5)

	visualizeServer := visualize.NewServer(":8081")

	guardHandler := waf.NewHandler(
		redirect.NewHandler(yamlDnsProvider),
		defender,
		limiter,
		visualizeServer.GetVisualizer(),
	)

	proxyServer := proxy.NewServer(":8080", certificateProvider, guardHandler)

	log.Println("Starting Waffle Proxy on port :8080 🚀")

	if err := proxyServer.Start(); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

// loadLocalCustomCACerts reads the local custom CA certificates from the embedded file system.
// It reads the CA certificate file (ca.crt) located in the ".cert" directory and returns it as a slice of byte slices.
// This CA certificate is used for establishing trust during TLS/SSL handshakes.
func loadLocalCustomCACerts(certificates embed.FS) [][]byte {
	certBytes, _ := certificates.ReadFile(".cert/ca.crt")
	return [][]byte{certBytes}
}

// loadLocalCertPEMBlock reads the local server certificate (server.crt) from the embedded file system.
// It returns the certificate as a byte slice, which is later used to serve the server's public certificate in TLS/SSL connections.
func loadLocalCertPEMBlock(certificates embed.FS) []byte {
	certBytes, _ := certificates.ReadFile(".cert/server.crt")
	return certBytes
}

// loadLocalKeyPEMBlock reads the private key (server.key) from the embedded file system.
// It returns the private key as a byte slice, which is paired with the server certificate during TLS/SSL handshakes.
func loadLocalKeyPEMBlock(certificates embed.FS) []byte {
	certBytes, _ := certificates.ReadFile(".cert/server.key")
	return certBytes
}
