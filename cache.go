package main

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
