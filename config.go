package main

import (
	"bytes"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	DEFAULT_LRU_MAX_SIZE = 1000
	DEFAULT_TABLE_NAME   = "key_values"
	DEFAULT_PORT         = 8080
	DEFAULT_CACHE_TYPE   = "arc"
	DEFAULT_DEBUG_MODE   = false
)

type Config struct {
	Config *viper.Viper
}

func newConfig() *Config {
	config := viper.New()

	config.SetDefault("lru_max_size", DEFAULT_LRU_MAX_SIZE)
	config.SetDefault("table_name", DEFAULT_TABLE_NAME)
	config.SetDefault("port", DEFAULT_PORT)
	config.SetDefault("cache_type", DEFAULT_CACHE_TYPE)
	config.SetDefault("debug_mode", DEFAULT_DEBUG_MODE)

	return &Config{config}
}

func (c *Config) load(yaml []byte) {
	c.Config.SetConfigType("yaml")
	c.Config.ReadConfig(bytes.NewBuffer(yaml))
}

func (c *Config) loadFile(config_path string) {
	yaml, err := ioutil.ReadFile(config_path)
	if err != nil {
		panic(err)
	}
	c.Config.SetConfigType("yaml")
	c.Config.ReadConfig(bytes.NewBuffer(yaml))
}

const (
	CONFIG_KEY_DRIVER_NAME      = "driver_name"
	CONFIG_KEY_DATA_SOURCE_NAME = "data_source_name"
	DEFAULT_DATABASE_NAME       = "key_value.db"
	CONFIG_KEY_LRU_MAX_SIZE     = "lru_max_size"
	CONFIG_KEY_TABLE_NAME       = "table_name"
	CONFIG_KEY_PORT             = "port"
	CONFIG_KEY_CACHE_TYPE       = "cache_type"
	CONFIG_KEY_DEBUG_MODE       = "debug_mode"
	CONFIG_KEY_WORK_DIR         = "work_dir"
)

func (c *Config) getDriverName() string {
	return c.Config.Get(CONFIG_KEY_DRIVER_NAME).(string)
}

func (c *Config) getDataSourceName() string {
	if c.Config.Get(CONFIG_KEY_DATA_SOURCE_NAME) == nil {
		workDir := c.getWorkDir()
		c.Config.Set(CONFIG_KEY_DATA_SOURCE_NAME, filepath.Join(workDir, DEFAULT_DATABASE_NAME))
	}
	return c.Config.Get(CONFIG_KEY_DATA_SOURCE_NAME).(string)
}

func (c *Config) getLruMaxSize() int {
	return c.Config.GetInt(CONFIG_KEY_LRU_MAX_SIZE)
}

func (c *Config) getTableName() string {
	return c.Config.Get(CONFIG_KEY_TABLE_NAME).(string)
}

func (c *Config) getPort() int {
	return c.Config.GetInt(CONFIG_KEY_PORT)
}

func (c *Config) getCacheType() string {
	return c.Config.Get(CONFIG_KEY_CACHE_TYPE).(string)
}

func (c *Config) getDebugMode() bool {
	return c.Config.GetBool(CONFIG_KEY_DEBUG_MODE)
}

func (c *Config) getWorkDir() string {
	if c.Config.Get(CONFIG_KEY_WORK_DIR) == nil {
		work_dir, err := ioutil.TempDir("", "")
		if err != nil {
			panic(err)
		}
		c.Config.Set(CONFIG_KEY_WORK_DIR, work_dir)
	}
	return c.Config.Get(CONFIG_KEY_WORK_DIR).(string)
}
