//Contains the LRU cache interface and struct definitions.

package cache

import (
	"sync"
	"time"
)

// CacheItem represents a single item in the cache.
type CacheItem struct {
	Value      string
	Expiration int64 // Time when the item expires in Unix timestamp
}

// LRUCache represents a thread-safe LRU cache.
type LRUCache struct {
	capacity int
	items    map[string]*CacheItem
	mutex    sync.RWMutex
}

// NewLRUCache creates a new LRUCache with the given capacity.
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		items:    make(map[string]*CacheItem),
	}
}

// Set adds an item to the cache with a specified expiration in seconds.
func (c *LRUCache) Set(key, value string, expiration int64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(c.items) >= c.capacity {
		// Evict an item (simple approach, not LRU eviction)
		for k := range c.items {
			delete(c.items, k)
			break
		}
	}

	c.items[key] = &CacheItem{
		Value:      value,
		Expiration: time.Now().Unix() + expiration,
	}
}

// Get retrieves an item from the cache.
func (c *LRUCache) Get(key string) (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, found := c.items[key]
	if !found {
		return "", false
	}

	if time.Now().Unix() > item.Expiration {
		delete(c.items, key)
		return "", false
	}

	return item.Value, true
}

// Delete removes an item from the cache.
func (c *LRUCache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.items, key)
}

// evictExpired removes expired items from the cache.
func (c *LRUCache) evictExpired() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for key, item := range c.items {
		if time.Now().UnixNano() > item.Expiration {
			delete(c.items, key)
		}
	}
}
