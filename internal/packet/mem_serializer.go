package packet

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/gopacket"

	"waffle/internal/worker"
)

// MemoryPacketSerializer is used to temp serialize
// packets in the memory, this way we can use this
// serialized packets to validate against DDOS attacks
type MemoryPacketSerializer struct {
	//ttl of the packets
	ttl            time.Duration
	addrPacketsMap map[string][]*inMemoryPacket

	mu sync.Mutex
}

var _ worker.PacketSerializer = (*MemoryPacketSerializer)(nil)

// NewMemoryPacketSerializer creates new memory packet serializer.
//
//	ttl - ttl of the packet, after this time packets are cleaned
//	from the memory.
func NewMemoryPacketSerializer(ttl time.Duration) *MemoryPacketSerializer {
	mps := &MemoryPacketSerializer{
		ttl:            ttl,
		addrPacketsMap: make(map[string][]*inMemoryPacket),
	}

	go mps.runCleaner()

	return mps
}

// SerializePackets is used to serialize packets in memory. Those packets are then removed after TTL in the another
// goroutine that is created when new instance of MemoryPacketSerializer is created by NewMemoryPacketSerializer.
func (m *MemoryPacketSerializer) SerializePackets(ctx context.Context, packetsChan <-chan gopacket.Packet) error {
	for {
		select {
		case packet, ok := <-packetsChan:
			if !ok {
				return nil
			}

			func() {
				m.mu.Lock()
				defer m.mu.Unlock()

				id := packet.NetworkLayer().NetworkFlow().Src().String()

				inMemPacket := &inMemoryPacket{
					obtainedAt: time.Now(),
					packet:     packet,
				}

				v, ok := m.addrPacketsMap[id]
				if !ok {
					m.addrPacketsMap[id] = []*inMemoryPacket{inMemPacket}
					return
				}

				v = append(v, inMemPacket)

				m.addrPacketsMap[id] = v
			}()
		case <-ctx.Done():
			return nil
		}
	}
}

func (m *MemoryPacketSerializer) runCleaner() {
	for {
		// for now let's go with two minutes, but I guess it should be adjusted later
		time.Sleep(time.Minute * 2)

		log.Println("started cleaning the serialized packets")

		now := time.Now()

		m.mu.Lock()

		// each few minutes we are clearing map from the packets that are no longer relevant
		for k, v := range m.addrPacketsMap {
			for i, packet := range v {
				if packet.obtainedAt.Add(m.ttl).After(time.Now()) {
					// TODO: slice out of bounds, we need to gather id, and which packets to remove and then remove those in the other loop
					m.addrPacketsMap[k] = append(m.addrPacketsMap[k][:i], m.addrPacketsMap[k][i+1:]...)
				}
			}
		}

		m.mu.Unlock()

		elapsed := time.Since(now)

		log.Println(fmt.Sprintf("serialzied packets clean done. Elapsed: %d", elapsed))
	}
}

type inMemoryPacket struct {
	obtainedAt time.Time
	packet     gopacket.Packet
}
