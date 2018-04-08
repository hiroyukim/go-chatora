package main

import (
	"log"
	"reflect"
	"testing"
)

func TestNewLevelDBCache(t *testing.T) {
	config := NewConfig()
	var cache Cache
	cache = NewLevelDBCache(config)

	if reflect.TypeOf(cache) != reflect.TypeOf(&LevelDBCache{}) {
		log.Fatal("TestNewLevelDBCache error: " + reflect.TypeOf(cache).String())
	}
}

func TestLevelDBCacheMethod(t *testing.T) {
	config := NewConfig()
	var cache Cache
	cache = NewLevelDBCache(config)

	cache.Set("2b", "B")
	value := cache.Get("2b")

	if value != "B" {
		log.Fatal("TestMethod error: " + value)
	}
}
