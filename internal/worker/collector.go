package worker

import (
	"context"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type PacketSerializer interface {
	SerializePackets(ctx context.Context, packetsChan <-chan gopacket.Packet) error
}

type NetworkInterfaceProvider interface {
	GetNetworkInterface() (*pcap.Interface, error)
}

type CollectorConfig struct {
	// BPF is a filter to filter out desired packets. IMPORTANT filter should only read the incoming packets.
	BPF string
}

// Collector collects packets from the network interfaces
// on the device it works (it could be a PC, server).
// Then based on the BPF it filters the packets.
type Collector struct {
	cfg                  *CollectorConfig
	serializer           PacketSerializer
	netInterfaceProvider NetworkInterfaceProvider
}

// NewCollector creates a new collector.
//
//	cfg - configuration of the collector
//	networkInterfaceProvider - provider of the network interface
//	serializer - serializer used to serialize packets.
func NewCollector(
	cfg CollectorConfig,
	networkInterfaceProvider NetworkInterfaceProvider,
	serializer PacketSerializer,
) *Collector {
	return &Collector{
		cfg:                  &cfg,
		netInterfaceProvider: networkInterfaceProvider,
		serializer:           serializer,
	}
}

// Run starts the packet capturing process for the Collector.
// It retrieves the network interface using the netInterfaceProvider.
// Then, it opens a live packet capture session on the interface using pcap, with a
// specified snapshot length of 1600 bytes, setting the device to promiscuous mode
// and waiting indefinitely for packets to arrive.
//
// If a BPF is provided in the configuration (c.cfg.BPF),
// it applies the filter to capture only relevant packets.
//
// The packets are read using a gopacket PacketSource and passed through a channel
// to be serialized by the Serializer in a separate goroutine.
//
// The function listens for packets in an infinite loop, and if the context is canceled
// (signaling termination), it gracefully exits.
//
// Deferred actions ensure that the packet channel is closed and a log message is generated
// when the collector is closed
func (c *Collector) Run(ctx context.Context) error {
	netInterface, err := c.netInterfaceProvider.GetNetworkInterface()
	if err != nil {
		return fmt.Errorf("get network interface using net interface provider: %w", err)
	}

	handle, err := pcap.OpenLive(netInterface.Name, 1600, true, pcap.BlockForever)
	if err != nil {
		return fmt.Errorf("pcap open live: %w", err)
	}
	defer handle.Close()

	if err := handle.SetBPFFilter(c.cfg.BPF); err != nil {
		return fmt.Errorf("set BPF filter: %w", err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packetsChan := make(chan gopacket.Packet)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	defer func() {
		close(packetsChan)
		log.Println("collector closed")
	}()

	go func() {
		if err := c.serializer.SerializePackets(ctx, packetsChan); err != nil {
			log.Printf("Error in serialize packets: %v\n", err)
		}
	}()

	go func() {
		select {
		case <-signalChan:
			log.Println("Received shutdown signal")
			cancel()
		case <-ctx.Done():

		}
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("Shutting down gracefully")
			return nil
		case packet := <-packetSource.Packets():
			if packet == nil {
				log.Println("No more packets to read, shutting down")
				return nil
			}
			packetsChan <- packet
		}
	}
}
