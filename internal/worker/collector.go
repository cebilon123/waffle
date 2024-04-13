package worker

import (
	"context"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
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

func (c *Collector) Run(ctx context.Context) error {
	netInterface, err := c.netInterfaceProvider.GetNetworkInterface()
	if err != nil {
		return fmt.Errorf("get network interface using net interface provider: %w", err)
	}

	handle, err := pcap.OpenLive(netInterface.Name, 1600, true, pcap.BlockForever)
	if err != nil {
		return fmt.Errorf("pcap open live: %w", err)
	}

	if err := handle.SetBPFFilter(c.cfg.BPF); err != nil {
		return fmt.Errorf("set BPF filter: %w", err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	packetsChan := make(chan gopacket.Packet)

	defer func() {
		close(packetsChan)

		log.Println("collector closed")
	}()

	go func() {
		if err := c.serializer.SerializePackets(ctx, packetsChan); err != nil {
			log.Println("error in serialize packets")
		}
	}()

	for {
		select {
		case packet, ok := <-packetSource.Packets():
			if !ok {
				log.Println("error reading packet")
			}

			packetsChan <- packet

		case <-ctx.Done():
			return nil
		}
	}
}
