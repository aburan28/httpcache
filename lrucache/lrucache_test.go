package lrucache

import (
	"fmt"
	"testing"

	"github.com/tj/assert"
)

func TestMaxCacheSize(t *testing.T) {
	maxCacheSize := 3
	lruCache := NewLRUCache(maxCacheSize)

	for i := 0; i < 5; i++ {
		lruCache.Set(fmt.Sprintf("%d", i), []byte("value"))

	}

	if lruCache.Size() > maxCacheSize {
		t.Errorf("expected cache size to be %d, got %d", maxCacheSize, lruCache.Size())
	}
}

func TestSizeOfCache(t *testing.T) {
	maxCacheSize := 5
	lruCache := NewLRUCache(maxCacheSize)

	for i := 0; i < 10; i++ {
		lruCache.Set(fmt.Sprintf("%d", i), []byte("value"))

	}

	assert.Equal(t, lruCache.Size(), maxCacheSize, "expected cache size to be %d, got %d", maxCacheSize, lruCache.Size())
}

func TestLeastRecentlyUsedEviction(t *testing.T) {
	maxCacheSize := 3
	lruCache := NewLRUCache(maxCacheSize)
	lruCache.Set("0", []byte("value1"))
	lruCache.Set("1", []byte("value1"))
	lruCache.Set("2", []byte("value2"))
	lruCache.Set("3", []byte("value3"))
	lruCache.Set("4", []byte("value4"))
	value, found := lruCache.Get("1")

	assert.False(t, found, "expected key '1' to be evicted, but it was found in the cache")
	assert.Nil(t, value, "expected key '1' to be evicted (value should be nil)")

}

func TestLeastRecentlyUsedWithGetCalls(t *testing.T) {
	maxCacheSize := 3
	lruCache := NewLRUCache(maxCacheSize)
	lruCache.Set("0", []byte("value1"))
	lruCache.Set("1", []byte("value1"))
	lruCache.Set("2", []byte("value2"))
	lruCache.Get("0")
	lruCache.Set("3", []byte("value3"))
	lruCache.Get("1")
	lruCache.Set("4", []byte("value4"))
	value, found := lruCache.Get("2")

	assert.False(t, found, "expected key '2' to be evicted, but it was found in the cache")
	assert.Nil(t, value, "expected key '2' to be evicted (value should be nil)")

}
