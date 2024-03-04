package main

import (
	_ "embed"
	"log"

	_ "github.com/go-sql-driver/mysql"

	cert "waffle/.cert"
	"waffle/internal/certificate"
	"waffle/internal/config"
	"waffle/internal/proxy"
)

func main() {
	_, err := config.LoadEnvironmentConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	certificateProvider := certificate.NewLocalCertificatesProvider(loadLocalCustomCACerts())

	server := proxy.NewServer(dns, ":8080", certificateProvider)

	log.Println("Starting Waffle Proxy on port :8080 🚀")

	if err := server.Start(); err != nil {
		log.Fatal(err.Error())
	}
}

func loadLocalCustomCACerts() [][]byte {
	certBytes, _ := cert.Certificates.ReadFile("ca.crt")

	return [][]byte{certBytes}
}
