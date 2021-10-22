package cache

import (
	"errors"
	"time"
)

type Cache struct {
	item map[string]interface{}
}

func New() *Cache {
	return &Cache{
		item: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.item[key] = value
	go func() {
		time.Sleep(ttl)
		delete(c.item, key)
	}()
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
