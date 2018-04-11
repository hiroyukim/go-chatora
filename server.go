package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack"
)

type ResponseJson struct {
	Value string `json:"value"`
}

type ResponseMsgp struct {
	Value string
}

func (h *ProxyHandler) indexDel(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	key := q.Get("key")
	h.Cache.del(key)

	responseType := q.Get("type")
	h.writeResponse(responseType, "", w, req)
}

func (h *ProxyHandler) indexGet(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	key := q.Get("key")
	responseType := q.Get("type")

	cache_value := h.Cache.get(key)
	var value string
	var err error
	if cache_value != "" {
		value = cache_value
		if DEBUG_MODE {
			logrus.WithFields(logrus.Fields{
				"key":   key,
				"value": value,
			}).Info("OK cache")
		}
	} else {
		value, err = h.Kvdb.getValue(key)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"key": key,
			}).Warn(err)
		} else {
			h.Cache.set(key, value)
			if DEBUG_MODE {
				logrus.WithFields(logrus.Fields{
					"key":   key,
					"value": value,
				}).Info("Add cache")
			}
		}
	}

	h.writeResponse(responseType, value, w, req)
}

const (
	CONTENT_TYPE       = "Content-Type"
	CONTENT_TYPE_JSON  = "application/json"
	CONTENT_TYPE_PLAIN = "text/plain"
	CONTENT_TYPE_MSGP  = "application/x-msgpack"
)

const (
	RESPONSE_TYPE_JSON = "json"
	RESPONSE_TYPE_MSGP = "msgp"
)

func (h *ProxyHandler) writeResponse(responseType string, value string, w http.ResponseWriter, req *http.Request) {
	switch responseType {
	case RESPONSE_TYPE_JSON:
		w.Header().Set(CONTENT_TYPE, CONTENT_TYPE_JSON)
		data := ResponseJson{value}
		json_string, _ := json.Marshal(data)
		fmt.Fprintf(w, string(json_string))
	case RESPONSE_TYPE_MSGP:
		w.Header().Set(CONTENT_TYPE, CONTENT_TYPE_MSGP)
		data := ResponseMsgp{value}
		b, _ := msgpack.Marshal(data)
		fmt.Fprintf(w, string(b))
	default:
		w.Header().Set(CONTENT_TYPE, CONTENT_TYPE_PLAIN)
		fmt.Fprintf(w, value)
	}
}

type ProxyHandler struct {
	Config *Config
	Kvdb   *KeyValueDB
	Cache  Cache
}

func (h *ProxyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		h.indexGet(w, req)
	case "DELETE":
		h.indexDel(w, req)
	default:
		http.NotFound(w, req)
	}
}

func runServer(config *Config, kvdb *KeyValueDB, cache Cache) {
	//TODO memcached protocol
	proxy_handler := &ProxyHandler{config, kvdb, cache}

	//TODO unix socket
	address := ":" + strconv.Itoa(config.getPort())
	logrus.WithFields(logrus.Fields{
		"address": address,
	}).Info("start server")
	if err := http.ListenAndServe(address, proxy_handler); err != nil {
		log.Fatal(err)
	}
}
