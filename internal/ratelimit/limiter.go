package ratelimit

import (
	"context"
	"net"
	"time"

	"github.com/google/uuid"

	"waffle/internal/cache"
)

var (
	defaultPurgeDuration = time.Minute * 5
)

// Limiter is an interface that defines rate-limiting functionality for network addresses.
// It includes methods to retrieve the current rate for an IP address and to set a rate limit for an IP address until a specific time.
type Limiter interface {
	GetRate(ctx context.Context, address net.IP) *Rate
	SetRate(ctx context.Context, address net.IP, limitedUntil time.Time) string
}

// InMemoryLimiter is a concrete implementation of the Limiter interface.
// It stores rate limit data in memory using a cache, with each IP address associated with its corresponding rate information.
type InMemoryLimiter struct {
	cache *cache.Cache[stringIpAddress, Rate]
}

// NewInMemoryLimiter creates and returns a new instance of InMemoryLimiter.
// It accepts an expirationTime, which defines how long rate limit data will be cached.
// The cache will automatically purge expired data at twice the expiration time or at a default purge duration, whichever is greater.
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

// Ensures that InMemoryLimiter implements the Limiter interface.
var _ Limiter = (*InMemoryLimiter)(nil)

// GetRate retrieves the rate limit information for the given IP address from the in-memory cache.
// If no rate limit is found for the IP address, it returns nil.
func (i *InMemoryLimiter) GetRate(_ context.Context, address net.IP) *Rate {
	rate, _ := i.cache.Get(stringIpAddress(address.String()))
	return rate
}

// SetRate sets the rate limit for the given IP address until the specified limitedUntil time.
// If a rate limit already exists for the IP address but the existing limit expires later than the provided time,
// the limit is not updated. The method returns a UUID associated with the new rate limit if it is set, or an empty string otherwise.
func (i *InMemoryLimiter) SetRate(_ context.Context, address net.IP, limitedUntil time.Time) string {
	addrString := address.String()

	rate, ok := i.cache.Get(stringIpAddress(addrString))
	if ok && rate.LimitedUntil.After(limitedUntil) {
		return ""
	}

	id := uuid.NewString()

	i.cache.Set(stringIpAddress(addrString), Rate{
		UUID:         id,
		IpAddress:    address,
		LimitedUntil: limitedUntil,
	})

	return id
}

// stringIpAddress is a type alias for string, used to represent IP addresses in the cache.
type stringIpAddress string
