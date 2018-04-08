package main

import (
	"log"
	"reflect"
	"testing"
)

func TestNewARCCache(t *testing.T) {
	config := NewConfig()
	var cache Cache
	cache = NewARCCache(config)

	if reflect.TypeOf(cache) != reflect.TypeOf(&ARCCache{}) {
		log.Fatal("TestNewARCCache error: " + reflect.TypeOf(cache).String())
	}
}

func TestARCCacheMethod(t *testing.T) {
	config := NewConfig()
	var cache Cache
	cache = NewARCCache(config)

	cache.Set("2b", "B")
	value := cache.Get("2b")

	if value != "B" {
		log.Fatal("TestMethod error: " + value)
	}
}
