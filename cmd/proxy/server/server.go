package server

import (
	"context"
	"embed"
	"log"
	"os"
	"os/signal"

	"waffle/internal/certificate"
	"waffle/internal/config"
	"waffle/internal/domain"
	"waffle/internal/proxy"
	"waffle/internal/redirect"
)

func Run(ctx context.Context, yamlConfigBytes []byte, certificates embed.FS) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
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

	redirectHandler := redirect.NewHandler(yamlDnsProvider)

	proxyServer := proxy.NewServer(":8080", certificateProvider, redirectHandler)

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
