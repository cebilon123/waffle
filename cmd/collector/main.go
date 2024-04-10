package main

import (
	"context"
	"golang.org/x/sys/windows"
	"log"

	"waffle/internal/worker"
	"waffle/pkg/permission"
)

func main() {
	if !windows.GetCurrentProcessToken().IsElevated() {
		if err := permission.RunMeElevated(); err != nil {
			panic(err.Error())
		}
	}

	ctx := context.Background()

	log.Println("starting collector")

	collector := worker.NewCollector(worker.CollectorConfig{
		Protocol: "tcp",
		Port:     "8080",
	})

	if err := collector.Run(ctx); err != nil {
		panic(err.Error())
	}
}
