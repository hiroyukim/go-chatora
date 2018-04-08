package main

import (
	"io/ioutil"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
)

type LevelDBCache struct {
	Cache *leveldb.DB
}

func NewLevelDBCache(config *Config) *LevelDBCache {
	var work_dir string
	var err error
	if config.GetWorkDir() == "" {
		work_dir, err = ioutil.TempDir("", "")
		if err != nil {
			panic(err)
		}
	} else {
		work_dir = config.GetWorkDir()
	}
	logrus.WithFields(logrus.Fields{
		"work_dir": work_dir,
	}).Info("Create work_dir")

	db, err := leveldb.OpenFile(filepath.Join(work_dir, "leveldb.db"), nil)
	if err != nil {
		panic(err)
	}

	return &LevelDBCache{db}
}

func (c *LevelDBCache) Get(key string) string {
	value, err := c.Cache.Get([]byte(key), nil)
	if err != nil {
		return ""
	}

	return string(value[:])
}

func (c *LevelDBCache) Set(key string, value string) {
	c.Cache.Put([]byte(key), []byte(value), nil)
}

func (c *LevelDBCache) Del(key string) {
	c.Cache.Delete([]byte(key), nil)
}
