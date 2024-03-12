package ddos

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/emirpasic/gods/sets/treeset"
)

type IPTreeSetProvider interface {
	GetSet() *treeset.Set
}

// SyncIPTreeSetProvider is a treeset.Set provider
// with additional functionality: it tries to update itself from time to time
// based on the configuration.
type SyncIPTreeSetProvider struct {
	set      *treeset.Set
	interval time.Duration
	fetch    fetchIPSliceFunc

	mu sync.RWMutex
}

var _ IPTreeSetProvider = (*SyncIPTreeSetProvider)(nil)

// NewSyncIPTreeSetProvider creates new SyncIPTreeSetProvider based
// on the interval of set refresh.
// Param: fetch is a func that is being executed per each interval
func NewSyncIPTreeSetProvider(interval time.Duration, fetch fetchIPSliceFunc) *SyncIPTreeSetProvider {
	treeSetProvider := &SyncIPTreeSetProvider{
		set:      nil,
		interval: interval,
		fetch:    fetch,
	}

	return treeSetProvider
}

func (ip *SyncIPTreeSetProvider) GetSet() *treeset.Set {
	ip.mu.RLock()
	defer ip.mu.RUnlock()

	return ip.set
}

func (ip *SyncIPTreeSetProvider) startUpdating() {
	go func() {
		for {
			go ip.tryUpdateSet()

			time.Sleep(ip.interval)
		}
	}()
}

func (ip *SyncIPTreeSetProvider) tryUpdateSet() {
	ip.mu.Lock()
	defer ip.mu.Unlock()

	ctx := context.Background()

	ips, err := ip.fetch(ctx)
	if err != nil {
		log.Println(fmt.Errorf("failed to fetch ips: %w", err).Error())

		return
	}

	ip.set = treeset.NewWith(ipComparator, ips)
}

type fetchIPSliceFunc func(ctx context.Context) ([]net.IP, error)

func ipComparator(a interface{}, b interface{}) int {
	aIP := a.(net.IP)
	bIP := b.(net.IP)

	return bytes.Compare(aIP, bIP)
}
