package cache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	item map[string]interface{}
	mu   sync.Mutex
}

func New() *Cache {
	return &Cache{
		item: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	c.item[key] = value
	go func() {
		time.Sleep(ttl)
		c.Delete(key)
	}()
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (interface{}, error) {
	if value, ok := c.item[key]; ok {
		return value, nil
	}

	return nil, errors.New("no such key")
}

func (c *Cache) Delete(key string) {
	delete(c.item, key)
}
