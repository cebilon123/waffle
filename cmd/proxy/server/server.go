package server

import (
	"context"
	"embed"
	"log"
	"os"
	"os/signal"
	"time"

	"waffle/internal/certificate"
	"waffle/internal/config"
	"waffle/internal/domain"
	"waffle/internal/proxy"
	"waffle/internal/ratelimit"
	"waffle/internal/redirect"
	"waffle/internal/waf"
	"waffle/internal/waf/guard"
)

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

	guardHandler := waf.NewHandler(redirect.NewHandler(yamlDnsProvider), defender, limiter)

	proxyServer := proxy.NewServer(":8080", certificateProvider, guardHandler)

	log.Println("Starting Waffle Proxy on port :8080 🚀")

	if err := proxyServer.Start(); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

func loadLocalCustomCACerts(certificates embed.FS) [][]byte {
	certBytes, _ := certificates.ReadFile(".cert/ca.crt")

	return [][]byte{certBytes}
}

func loadLocalCertPEMBlock(certificates embed.FS) []byte {
	certBytes, _ := certificates.ReadFile(".cert/server.crt")

	return certBytes
}

func loadLocalKeyPEMBlock(certificates embed.FS) []byte {
	certBytes, _ := certificates.ReadFile(".cert/server.key")

	return certBytes
}
