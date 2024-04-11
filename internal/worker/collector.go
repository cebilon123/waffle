package worker

import (
	"context"
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type PacketSerializer interface {
	SerializePackets(ctx context.Context, packetsChan chan<- gopacket.Packet) error
}

type NetworkInterfaceProvider interface {
	GetNetworkInterface() (*pcap.Interface, error)
}

type CollectorConfig struct {
	Protocol string
	Port     string
}

type Collector struct {
	cfg                  *CollectorConfig
	serializer           PacketSerializer
	netInterfaceProvider NetworkInterfaceProvider
}

func NewCollector(cfg CollectorConfig, deviceProvider NetworkInterfaceProvider) *Collector {
	return &Collector{
		cfg:                  &cfg,
		netInterfaceProvider: deviceProvider,
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

	if err := handle.SetBPFFilter("ip"); err != nil {
		return fmt.Errorf("set BPFF filter: %w", err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

	packetsChan := make(chan gopacket.Packet)
	defer func() {
		close(packetsChan)

		log.Println("collector closed")
	}()

	//go func() {
	//	if err := c.serializer.SerializePackets(ctx, packetsChan); err != nil {
	//		log.Println("error in serialize packets")
	//	}
	//}()

	for {
		select {
		case packet, ok := <-packetSource.Packets():
			if !ok {
				log.Println("error reading packet")
			}
			log.Println(packet.String())
			//packetsChan <- packet

		case <-ctx.Done():
			return nil
		}
	}
}
