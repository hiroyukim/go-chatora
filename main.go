package main

import (
	"flag"
	"log"
	"net/http"
	"runtime"

	_ "net/http/pprof"
)

const version = "0.0.1"

var (
	option_config_path string
	DEBUG_MODE         bool
	FORCED_CACHE_FLAG  bool
)

func init() {
	flag.StringVar(&option_config_path, "c", "", "help message for c option")
	flag.Parse()
}

func main() {
	if option_config_path == "" {
		panic("require option_config_path(c)" + option_config_path)
	}

	config := newConfig()
	config.loadFile(option_config_path)

	DEBUG_MODE = config.getDebugMode()
	FORCED_CACHE_FLAG = config.getForcedCacheFlag()

	if DEBUG_MODE {
		runtime.SetBlockProfileRate(1)
		go func() {
			log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
		}()
	}

	kvdb := newKeyValueDB(config)
	cache := newCache(config)

	runServer(config, kvdb, cache)
}
