package main

import (
	"log"

	"waffle/internal/domain"
	"waffle/internal/proxy"
)

func main() {
	dns := domain.NewYAMLBasedNameSystem()

	server := proxy.NewServer(dns, ":8080")

	if err := server.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
