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

type CollectorConfig struct {
	Protocol string
	Port     string
}

type Collector struct {
	cfg        *CollectorConfig
	serializer PacketSerializer
}

func NewCollector(cfg CollectorConfig) *Collector {
	return &Collector{
		cfg: &cfg,
	}
}

func (c *Collector) Run(ctx context.Context) error {
	handle, err := pcap.OpenLive("eth0", 1600, true, pcap.BlockForever)
	if err != nil {
		return fmt.Errorf("pcap open live: %w", err)
	}

	if err := handle.SetBPFFilter(fmt.Sprintf("%s port %s", c.cfg.Protocol, c.cfg.Port)); err != nil {
		return fmt.Errorf("set BPFF filter: %w", err)
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
