package main

import (
	"log"
	"reflect"
	"testing"
)

func TestNewLevelDBCache(t *testing.T) {
	config := newConfig()
	var cache Cache
	cache = newLevelDBCache(config)

	if reflect.TypeOf(cache) != reflect.TypeOf(&LevelDBCache{}) {
		log.Fatal("TestNewLevelDBCache error: " + reflect.TypeOf(cache).String())
	}
}

func TestLevelDBCacheMethod(t *testing.T) {
	config := newConfig()
	var cache Cache
	cache = newLevelDBCache(config)

	cache.set("2b", "B")
	value := cache.get("2b")

	if value != "B" {
		log.Fatal("TestMethod error: " + value)
	}
}
