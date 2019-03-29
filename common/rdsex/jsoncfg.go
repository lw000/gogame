package ggrdsex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//{
//"host": "127.0.0.1:6379",
//"psd": "123456",
//"db": 0,
//"poolSize": 20,
//"minIdleConns": 5
//}

type RdsConfigStruct struct {
	Host         string `json:"host"`
	Psd          string `json:"psd"`
	Db           int64  `json:"db"`
	PoolSize     int64  `json:"poolSize"`
	MinIdleConns int64  `json:"minIdleConns"`
}

type Config struct {
	Cfg RdsConfigStruct
}

func NewConfig() *Config {
	return &Config{}
}

func LoadConfig(file string) (*Config, error) {
	cfg := &Config{}
	err := cfg.Load(file)
	return cfg, err
}

func (c *Config) Load(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, &c.Cfg); err != nil {
		return err
	}

	return nil
}

func (c Config) String() string {
	return fmt.Sprintf("{Host:%s Psd:%s Db:%d PoolSize:%d MinIdleConns:%d}", c.Cfg.Host, c.Cfg.Psd, c.Cfg.Db, c.Cfg.PoolSize, c.Cfg.MinIdleConns)
}
