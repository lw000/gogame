package ggrdsex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//{
//	"status": 0,
//	"comment": "dev=0开发环境 test=1测试环境 prod=2正式环境",
//	"dev": {
//		"host": "192.168.1.201:6379",
//		"psd": "",
//		"db": 0,
//		"poolSize": 20,
//		"minIdleConns": 5
//	},
//	"test": {
//		"host": "192.168.1.201:6379",
//		"psd": "",
//		"db": 0,
//		"poolSize": 20,
//		"minIdleConns": 5
//	},
//	"prod": {
//		"host": "127.0.0.1:6379",
//		"psd": "123456",
//		"db": 0,
//		"poolSize": 20,
//		"minIdleConns": 5
//	}
//}

type RedisConfigItemStruct struct {
	Host         string `json:"host"`
	Psd          string `json:"psd"`
	Db           int64  `json:"db"`
	PoolSize     int64  `json:"poolSize"`
	MinIdleConns int64  `json:"minIdleConns"`
}

type configStruct struct {
	Status  int64                 `json:"status"`
	Comment string                `json:"comment"`
	Dev     RedisConfigItemStruct `json:"dev"`
	Test    RedisConfigItemStruct `json:"test"`
	Prod    RedisConfigItemStruct `json:"prod"`
}

type Config struct {
	Cfg RedisConfigItemStruct
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

	var cfgstruct configStruct
	if err := json.Unmarshal(data, &cfgstruct); err != nil {
		return err
	}

	if cfgstruct.Status == 0 {
		c.Cfg = cfgstruct.Dev
	}

	if cfgstruct.Status == 1 {
		c.Cfg = cfgstruct.Test
	}

	if cfgstruct.Status == 1 {
		c.Cfg = cfgstruct.Prod
	}

	return nil
}

func (c Config) String() string {
	return fmt.Sprintf("{Host:%s Psd:%s Db:%d PoolSize:%d MinIdleConns:%d}", c.Cfg.Host, c.Cfg.Psd, c.Cfg.Db, c.Cfg.PoolSize, c.Cfg.MinIdleConns)
}
