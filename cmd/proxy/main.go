package main

import (
	_ "embed"
	"log"

	cert "waffle/.cert"
	"waffle/internal/certificate"
	"waffle/internal/domain"
	"waffle/internal/proxy"
)

func main() {
	dns := domain.NewYAMLBasedNameSystem()

	certificateProvider := certificate.NewLocalCertificatesProvider(loadLocalCustomCACerts())

	server := proxy.NewServer(dns, ":8080", certificateProvider)

	if err := server.Start(); err != nil {
		log.Fatal(err.Error())
	}
}

func loadLocalCustomCACerts() [][]byte {
	certBytes, _ := cert.Certificates.ReadFile("ca.crt")

	return [][]byte{certBytes}
}
