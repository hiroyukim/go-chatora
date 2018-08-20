package main

import "github.com/vmihailenco/msgpack"

type Cache interface {
	get(key string) string
	set(key string, value string)
	del(key string)
}

func newCache(config *Config) Cache {
	var cache Cache
	switch config.getCacheType() {
	case "arc":
		cache = newARCCache(config)
	case "leveldb":
		cache = newLevelDBCache(config)
	default:
		panic("unexcpected cache type" + config.getCacheType())
	}

	return cache
}

type CacheData struct {
	data  string
	utime int64
}

func newCacheData(data string, utime int64) *CacheData {
	return &CacheData{data, utime}
}

func unmarshalCacheData(cache_value []byte) *CacheData {
	var c CacheData
	err = msgpack.Unmarshal(cache_value, &c)
	if err != nil {
		return nil
	}
	return c
}

func (c *CacheData) timeOver(now int64, duration int64) bool {
	return now > s.utime+duration
}

func (c *CacheData) marshalCacheData() ([]byte, error) {
	return msgpack.Marshal(c)
}
