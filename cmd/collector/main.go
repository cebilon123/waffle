package main

import (
	"context"
	"log"
	"waffle/internal/packet"

	"waffle/internal/worker"
)

const networkInterfaceDescription = "WAN Miniport (Network Monitor)"

func main() {
	ctx := context.Background()

	log.Println("starting collector")

	// NEXT TODO: add BPF filter builder
	// https://www.ibm.com/docs/en/qsip/7.4?topic=queries-berkeley-packet-filters
	collector := worker.NewCollector(worker.CollectorConfig{
		Protocol: "ip",
		Port:     "8080",
	}, packet.NewWindowsNetworkInterfaceProvider(networkInterfaceDescription))

	if err := collector.Run(ctx); err != nil {
		panic(err.Error())
	}
}
