package main

import lru "github.com/hashicorp/golang-lru"

type ARCCache struct {
	Cache *lru.ARCCache
}

func NewARCCache(config *Config) *ARCCache {
	cache, err := lru.NewARC(config.GetLruMaxSize())

	if err != nil {
		panic(err)
	}

	return &ARCCache{cache}
}

func (c *ARCCache) Get(key string) string {
	cache_value, cache_ok := c.Cache.Get(key)
	if !cache_ok {
		return ""
	} else {
		return cache_value.(string)
	}
}

func (c *ARCCache) Set(key string, value string) {
	c.Cache.Add(key, value)
}

func (c *ARCCache) Del(key string) {
	c.Cache.Remove(key)
}
