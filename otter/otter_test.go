package otter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetGetItems(t *testing.T) {
	cache := NewOtterCache()

	cache.Set("key1", []byte("value1"))
	cache.Set("key2", []byte("value2"))
	cache.Set("key3", []byte("value3"))
	k1, err := cache.Get("key1")
	require.Equal(t, "value1", string(k1))
	require.Equal(t, err, true)
	k2, err := cache.Get("key2")
	require.Equal(t, "value2", string(k2))
	require.Equal(t, err, true)

	k3, err := cache.Get("key3")
	require.Equal(t, "value3", string(k3))
	require.Equal(t, err, true)
	require.Equal(t, 3, cache.Size())
}

func TestMaxItems(t *testing.T) {
	cache := NewOtterCache()
	for i := 0; i < 10000; i++ {
		cache.Set(fmt.Sprint(i), []byte("value"))
	}
	require.Equal(t, 10000, cache.Size())
}
