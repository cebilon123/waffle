package main

import (
	"context"
	"embed"
	_ "embed"
	"fmt"
	"os"
	
	"waffle/cmd/proxy/server"
)

//go:embed config/config.yml
var yamlConfigBytes []byte

//go:embed .cert/*
var certificates embed.FS

func main() {
	ctx := context.Background()
	if err := server.Run(ctx, yamlConfigBytes, certificates); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
