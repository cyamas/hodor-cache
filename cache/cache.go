package cache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]Item
	mu   sync.RWMutex
}

type Item struct {
	value      interface{}
	expiration time.Time
}

func New() *Cache {
	return &Cache{
		data: make(map[string]Item),
	}
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	expiration := time.Now().Add(duration)
	c.data[key] = Item{
		value:      value,
		expiration: expiration,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.data[key]
	if !found {
		return nil, false
	}
	if time.Now().After(item.expiration) {
		delete(c.data, key)
		return nil, false
	}
	return item, true
}
