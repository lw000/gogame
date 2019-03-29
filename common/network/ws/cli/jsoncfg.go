package ggwscfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//{
//		"scheme": "ws",
//		"host": "47.96.230.81:8830",
//		"path": ""
//}

type WsConfigStruct struct {
	Host string `json:"host"`
	Path string `json:"path"`
}

type Config struct {
	Cfg WsConfigStruct
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
	return fmt.Sprintf("{Host:%s Path:%s}", c.Cfg.Host, c.Cfg.Path)
}
