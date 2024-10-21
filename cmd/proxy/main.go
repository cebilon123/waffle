package main

import (
	"context"
	"embed"
	_ "embed"
	"flag"
	"log"
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
	flag.StringVar(&visualizeServerPort, "p", "8081", "Port for server to listen on")
	flag.StringVar(&proxyServerPort, "p", "8081", "Port for server to listen on")

	ctx := context.Background()
	if err := server.Run(ctx, proxyServerPort, visualizeServerPort, yamlConfigBytes, certificates); err != nil {
		log.Fatalln(err)
	}
}
