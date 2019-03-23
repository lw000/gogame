package ggwscfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//{
//	"status": 0,
//	"comment": "dev=0开发环境 test=1测试环境 prod=2正式环境",
//	"dev": {
//		"scheme": "ws",
//		"host": "127.0.0.1:8830",
//		"path": ""
//	},
//	"test": {
//		"scheme": "ws",
//		"host": "192.168.1.186:8830",
//		"path": ""
//	},
//	"prod": {
//		"scheme": "ws",
//		"host": "47.96.230.81:8830",
//		"path": ""
//	}
//}

type WsConfigItemStruct struct {
	Scheme string `json:"scheme"`
	Host   string `json:"host"`
	Path   string `json:"path"`
}

type configStruct struct {
	Status  int64              `json:"status"`
	Comment string             `json:"comment"`
	Dev     WsConfigItemStruct `json:"dev"`
	Test    WsConfigItemStruct `json:"test"`
	Prod    WsConfigItemStruct `json:"prod"`
}

type Config struct {
	Cfg WsConfigItemStruct
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
	if err = json.Unmarshal(data, &cfgstruct); err != nil {
		return err
	}

	if cfgstruct.Status == 0 {
		c.Cfg = cfgstruct.Dev
	}

	if cfgstruct.Status == 1 {
		c.Cfg = cfgstruct.Test
	}

	if cfgstruct.Status == 2 {
		c.Cfg = cfgstruct.Prod
	}

	return nil
}

func (c Config) String() string {
	return fmt.Sprintf("{Scheme:%s Host:%s Path:%s}", c.Cfg.Scheme, c.Cfg.Host, c.Cfg.Path)
}
