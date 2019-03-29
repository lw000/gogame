package ggdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//{
//"username": "root",
//"password": "Aabbcc123!@#",
//"host": "47.96.230.81:3306",
//"database": "mservice",
//"MaxOpenConns": 20,
//"MaxOdleConns": 5
//}

type MysqlConfigStruct struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Database     string `json:"database"`
	MaxOpenConns int64  `json:"MaxOpenConns"`
	MaxOdleConns int64  `json:"MaxOdleConns"`
}

type Config struct {
	Cfg MysqlConfigStruct
}

func LoadConfigWithData(data []byte) (*Config, error) {
	cfg := &Config{}
	err := cfg.LoadWithData(data)
	return cfg, err
}

func LoadConfigWithJson(file string) (*Config, error) {
	cfg := &Config{}
	err := cfg.Load(file)
	return cfg, err
}

func (c *Config) LoadWithData(data []byte) error {
	if err := json.Unmarshal(data, &c.Cfg); err != nil {
		return err
	}
	return nil
}

func (c *Config) Load(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return c.LoadWithData(data)
}

func (c Config) String() string {
	return fmt.Sprintf("{Username:%s Password:%s Host:%s Database:%s MaxOpenConns:%d MaxOdleConns:%d}",
		c.Cfg.Username, c.Cfg.Password, c.Cfg.Host, c.Cfg.Database, c.Cfg.MaxOpenConns, c.Cfg.MaxOdleConns)
}
