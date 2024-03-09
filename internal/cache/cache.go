package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// Cache is a cache build based on github.com/patrickmn/go-cache.
// It extends its functionality by introducing the generics, which
// simplifies obtaining the value.
type Cache[K any, T any] struct {
	cache *cache.Cache
}

func New[K any, T any](defaultExpiration time.Duration, cleanupInterval time.Duration) *Cache[K, T] {
	c := cache.New(defaultExpiration, cleanupInterval)

	return &Cache[K, T]{
		cache: c,
	}
}

func (c *Cache[K, T]) Set(key string, val any) {
	c.cache.Set(key, val, cache.DefaultExpiration)
}

func (c *Cache[K, T]) Get(key string) (*T, bool) {
	v, ok := c.cache.Get(key)
	if !ok {
		return nil, false
	}

	vTyped, ok := v.(T)
	if !ok {
		return nil, false
	}

	return &vTyped, true
}
