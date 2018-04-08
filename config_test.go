package main

import (
	"log"
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	config := NewConfig()

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

	config := NewConfig()
	config.Load(yaml)

	if config.GetPort() != 80 {
		log.Fatal("GetPort error: " + string(config.GetPort()))
	}

	if config.GetLruMaxSize() != 1000000 {
		log.Fatal("GetLruMaxSize error: " + string(config.GetLruMaxSize()))
	}

	if config.GetDataSourceName() != "[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]" {
		log.Fatal("GetDataSourceName error: " + config.GetDataSourceName())
	}

	if config.GetDriverName() != "mysql" {
		log.Fatal("GetDriverName error:" + config.GetDriverName())
	}
}
