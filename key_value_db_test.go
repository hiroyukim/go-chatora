package main

import (
	"log"
	"reflect"
	"testing"
)

func TestNewKeyValueDB(t *testing.T) {
	yaml := []byte(`
---
driver_name: mysql
data_source_name: '[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]'
`)

	config := newConfig()
	config.load(yaml)
	db := newKeyValueDB(config)

	if reflect.TypeOf(db) != reflect.TypeOf(&KeyValueDB{}) {
		log.Fatal("TestNewKeyValueDB error: " + reflect.TypeOf(db).String())
	}
}

func TestGetValue(t *testing.T) {
	yaml := []byte(`
---
driver_name: sqlite3
`)

	config := newConfig()
	config.load(yaml)
	db := newKeyValueDB(config)

	_, err := db.DB.Exec(`
CREATE TABLE key_values
(
    "key" TEXT NOT NULL UNIQUE,
    value TEXT NOT NULL
);
`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.DB.Exec("INSERT INTO key_values (key,value) VALUES('9s','s')")
	if err != nil {
		log.Fatal(err)
	}

	var value string
	value, err = db.getValue("9s")

	if err != nil {
		log.Fatal(err)
	}

	if value != "s" {
		log.Fatal("TestGetValue error: " + value)
	}
}
