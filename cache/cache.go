package cache

import (
	"log"
	"sync"
	"time"
)

type Cache struct {
	data map[string]Item
	mu   sync.RWMutex
}

type Item struct {
	Value      any
	expiration time.Time
}

func New() *Cache {
	return &Cache{
		data: make(map[string]Item),
	}
}

func (c *Cache) Display() {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for k, v := range c.data {
		log.Printf("%s: %v", k, v)
	}
}

func (c *Cache) Set(key string, value any, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	expiration := time.Now().Add(duration)
	c.data[key] = Item{
		Value:      value,
		expiration: expiration,
	}
}

func (c *Cache) Get(key string) (any, bool) {
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
	return item.Value, true
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}
