package main

import (
	"embed"
	_ "embed"
	"log"

	"waffle/internal/certificate"
	"waffle/internal/config"
	"waffle/internal/domain"
	"waffle/internal/proxy"
)

//go:embed config/config.yml
var yamlConfigBytes []byte

//go:embed .cert/*
var Certificates embed.FS

func main() {
	_, err := config.LoadEnvironmentConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	yamlCfg, err := config.NewYamlConfig(yamlConfigBytes)
	if err != nil {
		log.Fatal(err.Error())
	}

	yamlDnsProvider := domain.NewYamlNameSystemProvider(yamlCfg)

	certificateProvider := certificate.NewLocalCertificatesProvider(loadLocalCustomCACerts(), loadLocalCertPEMBlock(), loadLocalKeyPEMBlock())

	server := proxy.NewServer(yamlDnsProvider, ":8080", certificateProvider)

	log.Println("Starting Waffle Proxy on port :8080 🚀")

	if err := server.Start(); err != nil {
		log.Fatal(err.Error())
	}
}

func loadLocalCustomCACerts() [][]byte {
	certBytes, _ := Certificates.ReadFile(".cert/ca.crt")

	return [][]byte{certBytes}
}

func loadLocalCertPEMBlock() []byte {
	certBytes, _ := Certificates.ReadFile(".cert/server.crt")

	return certBytes
}

func loadLocalKeyPEMBlock() []byte {
	certBytes, _ := Certificates.ReadFile(".cert/server.key")

	return certBytes
}
