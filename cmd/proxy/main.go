package main

import (
	"database/sql"
	_ "embed"
	"log"

	_ "github.com/go-sql-driver/mysql"

	cert "waffle/.cert"
	"waffle/internal/certificate"
	"waffle/internal/config"
	"waffle/internal/domain"
	"waffle/internal/proxy"
)

func main() {
	cfg, err := config.LoadEnvironmentConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := sql.Open("mysql", cfg.Database.URI)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	dns := domain.NewMysqlNameSystemProvider(db)

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
