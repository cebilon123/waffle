package main

import (
	"context"
	"embed"
	_ "embed"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"waffle/cmd/proxy/server"
)

//go:embed config/config.yml
var yamlConfigBytes []byte

//go:embed .cert/*
var certificates embed.FS

func main() {
	var (
		visualizeServerPort string
		proxyServerPort     string
	)
	flag.StringVar(&visualizeServerPort, "vp", "8080", "Port for visualizer server to listen on")
	flag.StringVar(&proxyServerPort, "p", "8081", "Port for server to listen on")

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.Run(proxyServerPort, visualizeServerPort, yamlConfigBytes, certificates); err != nil {
			log.Fatalln(err)
		}
	}()

	log.Println("Visualizer server started on ", visualizeServerPort)
	log.Println("Proxy server started on ", proxyServerPort)

	<-quit
	log.Println("Shutdown signal received, shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s\n", err)
	}
}
