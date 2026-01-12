package cache

import (
	"sync"
	"time"
)

type CacheItem struct {
	Value      interface{}
	Expiration int64
}

type MemoryCache struct {
	items map[string]CacheItem
	mu    sync.RWMutex
}

var (
	cacheInstance *MemoryCache
	once          sync.Once
)

func GetCache() *MemoryCache {
	once.Do(func() {
		cacheInstance = &MemoryCache{
			items: make(map[string]CacheItem),
		}
	})
	return cacheInstance
}

func (c *MemoryCache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	var expiration int64
	if ttl > 0 {
		expiration = time.Now().Add(ttl).Unix()
	}
	
	c.items[key] = CacheItem{
		Value:      value,
		Expiration: expiration,
	}
}

func (c *MemoryCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	item, found := c.items[key]
	if !found {
		return nil, false
	}
	
	if item.Expiration > 0 && time.Now().Unix() > item.Expiration {
		return nil, false
	}
	
	return item.Value, true
}

func (c *MemoryCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

func (c *MemoryCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[string]CacheItem)
}

func (c *MemoryCache) DeletePrefix(prefix string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	for key := range c.items {
		if len(key) >= len(prefix) && key[:len(prefix)] == prefix {
			delete(c.items, key)
		}
	}
}
