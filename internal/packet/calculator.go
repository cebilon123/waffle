package packet

import (
	"context"
	"sync"

	"github.com/google/gopacket"

	"waffle/internal/worker"
)

// Calculator is used in the packet serialization process.
// Implements PacketSerializer tho this struct recalculates
// the average size of the packets.
// https://www.tandfonline.com/doi/full/10.1080/23742917.2017.1384213
// https://journals.vilniustech.lt/index.php/MLA/article/view/4481/3817
type Calculator struct {
	mu sync.Mutex
}

var _ worker.PacketSerializer = (*Calculator)(nil)

func (c *Calculator) SerializePackets(ctx context.Context, packetsChan <-chan gopacket.Packet) error {
	//TODO implement me
	panic("implement me")
}
