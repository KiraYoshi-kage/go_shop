package cache

import (
	"errors"
	"sync"
	"time"
)

type item struct {
	value      interface{}
	expiration int64
}

type MemoryCache struct {
	items map[string]item
	mu    sync.RWMutex
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		items: make(map[string]item),
	}
}

func (c *MemoryCache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, errors.New("key not found")
	}

	if item.expiration > 0 && item.expiration < time.Now().UnixNano() {
		return nil, errors.New("key expired")
	}

	return item.value, nil
}

func (c *MemoryCache) Set(key string, value interface{}, expiration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	var exp int64
	if expiration > 0 {
		exp = time.Now().Add(expiration).UnixNano()
	}

	c.items[key] = item{
		value:      value,
		expiration: exp,
	}
	return nil
}

func (c *MemoryCache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
	return nil
}
