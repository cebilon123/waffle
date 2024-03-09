package ratelimit

import (
	"context"
	"github.com/google/uuid"
	"net"
	"time"

	"waffle/internal/cache"
)

var (
	defaultPurgeDuration = time.Minute * 5
)

type Limiter interface {
	GetRate(ctx context.Context, address net.IP) *Rate
	SetRate(ctx context.Context, address net.IP, limitedUntil time.Time) string
}

type InMemoryLimiter struct {
	cache *cache.Cache[stringIpAddress, Rate]
}

func NewInMemoryLimiter(expirationTime time.Duration) *InMemoryLimiter {
	purgeDuration := expirationTime * 2
	if purgeDuration < defaultPurgeDuration {
		purgeDuration = defaultPurgeDuration
	}

	c := cache.New[stringIpAddress, Rate](expirationTime, purgeDuration)

	return &InMemoryLimiter{
		cache: c,
	}
}

var _ Limiter = (*InMemoryLimiter)(nil)

func (i *InMemoryLimiter) GetRate(_ context.Context, address net.IP) *Rate {
	rate, _ := i.cache.Get(address.String())

	return rate
}

func (i *InMemoryLimiter) SetRate(_ context.Context, address net.IP, limitedUntil time.Time) string {
	addrString := address.String()

	rate, ok := i.cache.Get(addrString)
	if ok && rate.LimitedUntil.After(limitedUntil) {
		return ""
	}

	id := uuid.NewString()

	i.cache.Set(addrString, Rate{
		UUID:         id,
		IpAddress:    address,
		LimitedUntil: limitedUntil,
	})

	return id
}

type stringIpAddress string
