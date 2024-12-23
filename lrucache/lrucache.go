package lrucache

import (
	"log"

	lru "github.com/hashicorp/golang-lru/v2"
)

// LRUCache is a wrapper to implement the httpcache.Cache interface
type LRUCache struct {
	cache *lru.Cache[string, []byte]
}

// NewLRUCache creates a new LRUCache with the specified size
func NewLRUCache(size int) *LRUCache {
	c, err := lru.New[string, []byte](size)
	if err != nil {
		log.Fatalf("failed to create LRU cache: %v", err)
	}
	return &LRUCache{cache: c}
}

// Size returns the current size of the cache
func (c *LRUCache) Size() int {
	return c.cache.Len()
}

// Get retrieves a value from the cache
func (c *LRUCache) Get(key string) ([]byte, bool) {
	value, ok := c.cache.Get(key)
	return value, ok
}

// Set adds a value to the cache
func (c *LRUCache) Set(key string, value []byte) {
	c.cache.Add(key, value)
}

// Delete removes a value from the cache
func (c *LRUCache) Delete(key string) {
	c.cache.Remove(key)
}
