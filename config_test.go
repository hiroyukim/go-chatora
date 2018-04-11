package main

import (
	"log"
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	config := newConfig()

	if reflect.TypeOf(config) != reflect.TypeOf(&Config{}) {
		log.Fatal("NewConfig error: " + reflect.TypeOf(config).String())
	}
}

func TestLoad(t *testing.T) {
	yaml := []byte(`
---
driver_name: mysql
data_source_name: '[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]'
lru_max_size: 1000000
table_name: kvs
port: 80
cache_type: arc
`)

	config := newConfig()
	config.load(yaml)

	if config.getPort() != 80 {
		log.Fatal("GetPort error: " + string(config.getPort()))
	}

	if config.getLruMaxSize() != 1000000 {
		log.Fatal("GetLruMaxSize error: " + string(config.getLruMaxSize()))
	}

	if config.getDataSourceName() != "[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]" {
		log.Fatal("GetDataSourceName error: " + config.getDataSourceName())
	}

	if config.getDriverName() != "mysql" {
		log.Fatal("GetDriverName error:" + config.getDriverName())
	}
}
