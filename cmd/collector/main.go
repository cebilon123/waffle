package main

import (
	"context"
	"flag"
	"log"
	"time"
	"waffle/internal/packet"

	"waffle/internal/worker"
)

const (
	defaultNetworkInterfaceDescription = "Intel(R) I211 Gigabit Network Connection"
)

func main() {
	var (
		networkInterface string
	)
	flag.StringVar(&networkInterface, "i", defaultNetworkInterfaceDescription, "Identification of the interface")

	ctx := context.Background()

	log.Println("starting collector")

	inMemoryPacketSerializer := packet.NewMemoryPacketSerializer(time.Minute * 5)

	// NEXT TODO: add BPF filter builder
	// https://www.ibm.com/docs/en/qsip/7.4?topic=queries-berkeley-packet-filters
	cfg := worker.CollectorConfig{
		BPF: "ip",
	}

	collector := worker.NewCollector(
		cfg,
		packet.NewWindowsNetworkInterfaceProvider(networkInterface),
		inMemoryPacketSerializer)

	if err := collector.Run(ctx); err != nil {
		panic(err.Error())
	}
}
