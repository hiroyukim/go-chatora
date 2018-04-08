package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vmihailenco/msgpack"
)

func TestServer(t *testing.T) {
	yaml := []byte(`
---
driver_name: sqlite3
`)
	config := NewConfig()
	config.Load(yaml)
	kvdb := NewKeyValueDB(config)
	_, err := kvdb.DB.Exec(`
CREATE TABLE key_values
(
    "key" TEXT NOT NULL UNIQUE,
    value TEXT NOT NULL
);
`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = kvdb.DB.Exec("INSERT INTO key_values (key,value) VALUES('9s','s')")
	if err != nil {
		log.Fatal(err)
	}
	cache := NewARCCache(config)

	proxyHandler := &ProxyHandler{config, kvdb, cache}
	ts := httptest.NewServer(proxyHandler)
	defer ts.Close()

	r, err := http.Get(ts.URL + "/?key=9s")
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}
	data, err := ioutil.ReadAll(r.Body)

	if string(data) != "s" {
		t.Fatalf("Error value: %s", string(data))
	}

	// Del test
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", ts.URL+"/?key=9s", nil)
	r, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if cache.Get("9s") != "" {
		t.Fatalf("ErrorDel : %s", cache.Get("9s"))
	}

	// other methods
	otherMethods := []string{"HEAD", "POST", "OPTIONS", "PUT", "TRACE"}
	for _, method := range otherMethods {
		client := &http.Client{}
		req, err := http.NewRequest(method, ts.URL+"/?key=9s", nil)
		r, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		if r.StatusCode != 404 {
			t.Fatalf("ErrorOtherMethods : %s", method)
		}
	}

	// json test
	r, err = http.Get(ts.URL + "/?key=9s&type=json")
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}
	data, err = ioutil.ReadAll(r.Body)

	if string(data) != `{"value":"s"}` {
		t.Fatalf("Error value: %s", string(data))
	}

	// msgp test
	r, err = http.Get(ts.URL + "/?key=9s&type=msgp")
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}
	data, err = ioutil.ReadAll(r.Body)
	var rmsgp ResponseMsgp
	err = msgpack.Unmarshal(data, &rmsgp)
	if err != nil {
		panic(err)
	}

	if rmsgp.Value != "s" {
		t.Fatalf("Error value: %s", rmsgp.Value)
	}
}
