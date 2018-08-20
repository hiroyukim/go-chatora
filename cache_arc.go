package main

import (
	lru "github.com/hashicorp/golang-lru"
	"github.com/sirupsen/logrus"
)

type ARCCache struct {
	cache *lru.ARCCache
}

func newARCCache(config *Config) *ARCCache {
	cache, err := lru.NewARC(config.getLruMaxSize())

	if err != nil {
		panic(err)
	}

	return &ARCCache{cache}
}

func (c *ARCCache) get(key string) *CacheData {
	cache_value, cache_ok := c.cache.Get(key)
	if !cache_ok {
		nil
	} else {
		return unmarshalCacheData([]byte{cache_value.(string)})
	}
}

func (c *ARCCache) set(key string, cacheData *CacheData) {
	b, err := cachdData.marshalCacheData()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"key": key,
			"err": err,
		}).Warn(err)
	}
	c.cache.Add(key, string(b))
}

func (c *ARCCache) del(key string) {
	c.cache.Remove(key)
}
