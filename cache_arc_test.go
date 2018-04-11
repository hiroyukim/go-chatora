package main

import (
	"log"
	"reflect"
	"testing"
)

func TestNewARCCache(t *testing.T) {
	config := newConfig()
	var cache Cache
	cache = newARCCache(config)

	if reflect.TypeOf(cache) != reflect.TypeOf(&ARCCache{}) {
		log.Fatal("TestNewARCCache error: " + reflect.TypeOf(cache).String())
	}
}

func TestARCCacheMethod(t *testing.T) {
	config := newConfig()
	var cache Cache
	cache = newARCCache(config)

	cache.set("2b", "B")
	value := cache.get("2b")

	if value != "B" {
		log.Fatal("TestMethod error: " + value)
	}
}
