package main

import lru "github.com/hashicorp/golang-lru"

type ARCCache struct {
	Cache *lru.ARCCache
}

func newARCCache(config *Config) *ARCCache {
	cache, err := lru.NewARC(config.getLruMaxSize())

	if err != nil {
		panic(err)
	}

	return &ARCCache{cache}
}

func (c *ARCCache) get(key string) string {
	cache_value, cache_ok := c.Cache.Get(key)
	if !cache_ok {
		return ""
	} else {
		return cache_value.(string)
	}
}

func (c *ARCCache) set(key string, value string) {
	c.Cache.Add(key, value)
}

func (c *ARCCache) del(key string) {
	c.Cache.Remove(key)
}
