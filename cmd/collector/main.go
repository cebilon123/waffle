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

	// question: Why do we need context here? It is not used in collector.Run, except of ctx.Done, but since it is not
	// context.WithTimeout (as example) it can not be closed in any way.
	// Same in c.serializer.SerializePackets(ctx, packetsChan), it can not be closed there as well.
	// Why not just to remove it?
	ctx := context.Background()

	log.Println("starting collector")

	packetSerializer := packet.NewMemoryPacketSerializer(time.Minute * 5)

	// NEXT TODO: add BPF filter builder
	// https://www.ibm.com/docs/en/qsip/7.4?topic=queries-berkeley-packet-filters
	cfg := worker.CollectorConfig{
		BPF: "ip",
	}

	collector := worker.NewCollector(
		cfg,
		packet.NewWindowsNetworkInterfaceProvider(networkInterface),
		packetSerializer)

	if err := collector.Run(ctx); err != nil {
		log.Fatalln("Error during running collector: ", err.Error())
	}
}
