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

func NewConfig() *Config {
	config := viper.New()

	config.SetDefault("lru_max_size", DEFAULT_LRU_MAX_SIZE)
	config.SetDefault("table_name", DEFAULT_TABLE_NAME)
	config.SetDefault("port", DEFAULT_PORT)
	config.SetDefault("cache_type", DEFAULT_CACHE_TYPE)
	config.SetDefault("debug_mode", DEFAULT_DEBUG_MODE)

	return &Config{config}
}

func (c *Config) Load(yaml []byte) {
	c.Config.SetConfigType("yaml")
	c.Config.ReadConfig(bytes.NewBuffer(yaml))
}

func (c *Config) LoadFile(config_path string) {
	yaml, err := ioutil.ReadFile(config_path)
	if err != nil {
		panic(err)
	}
	c.Config.SetConfigType("yaml")
	c.Config.ReadConfig(bytes.NewBuffer(yaml))
}

func (c *Config) GetDriverName() string {
	return c.Config.Get("driver_name").(string)
}

func (c *Config) GetDataSourceName() string {
	if c.Config.Get("data_source_name") == nil {
		workDir := c.GetWorkDir()
		c.Config.Set("data_source_name", filepath.Join(workDir, "key_value.db"))
	}
	return c.Config.Get("data_source_name").(string)
}

func (c *Config) GetLruMaxSize() int {
	return c.Config.GetInt("lru_max_size")
}

func (c *Config) GetTableName() string {
	return c.Config.Get("table_name").(string)
}

func (c *Config) GetPort() int {
	return c.Config.GetInt("port")
}

func (c *Config) GetCacheType() string {
	return c.Config.Get("cache_type").(string)
}

func (c *Config) GetDebugMode() bool {
	return c.Config.GetBool("debug_mode")
}

func (c *Config) GetWorkDir() string {
	if c.Config.Get("work_dir") == nil {
		work_dir, err := ioutil.TempDir("", "")
		if err != nil {
			panic(err)
		}
		c.Config.Set("work_dir", work_dir)
	}
	return c.Config.Get("work_dir").(string)
}
