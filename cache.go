package main

type Cache interface {
	Get(key string) string
	Set(key string, value string)
	Del(key string)
}

func NewCache(config *Config) Cache {
	var cache Cache
	switch config.GetCacheType() {
	case "arc":
		cache = NewARCCache(config)
	case "leveldb":
		cache = NewLevelDBCache(config)
	default:
		panic("unexcpected cache type" + config.GetCacheType())
	}

	return cache
}
