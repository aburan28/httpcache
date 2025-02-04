package otter

import (
	"log"

	"github.com/maypok86/otter"
)

type Config struct {
}

type OtterCache struct {
	cache   *otter.Cache[string, []byte]
	maxSize int
}

func NewOtterCache() *OtterCache {
	cache, err := otter.MustBuilder[string, []byte](10_000).
		CollectStats().
		Cost(func(key string, value []byte) uint32 {
			return 1
		}).
		Build()

	if err != nil {
		log.Fatalf("failed to create Otter cache: %v", err)
	}

	return &OtterCache{cache: &cache}
}

func (o *OtterCache) Set(key string, value []byte) {
	o.cache.Set(key, value)
}

func (o *OtterCache) Get(key string) ([]byte, bool) {
	return o.cache.Get(key)
}

func (o *OtterCache) Delete(key string) {
	o.cache.Delete(key)
}

func (o *OtterCache) Size() int {
	return o.cache.Size()
}
