package cache

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	items map[string]item
	mu    sync.RWMutex
}

type item struct {
	value interface{}
	ttl   int64
}

func New() *Cache {
	cache := &Cache{items: make(map[string]item)}
	go cache.scanCache()
	return cache
}

func (r *Cache) Set(key string, value interface{}, ttl time.Duration) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[key] = item{value, time.Now().Add(ttl).Unix()}
}

func (r *Cache) Get(key string) (interface{}, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if item, ok := r.items[key]; ok {
		return item.value, nil
	}

	return nil, errors.New("no such key")
}

func (r *Cache) Delete(key string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.items, key)
}

func (r *Cache) scanCache() {
	for {
		r.clean()
	}
}

func (r *Cache) clean() {
	for key, value := range r.items {
		if time.Now().Unix() > value.ttl {
			r.Delete(key)
		}
	}
}
