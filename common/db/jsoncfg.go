package ggdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//{
//	"status": 0,
//	"comment": "dev=0开发环境 test=1测试环境 prod=2正式环境",
//	"dev": {
//		"username": "root",
//		"password": "root",
//		"network": "tcp",
//		"host": "192.168.1.101:3306",
//		"database": "mservice",
//		"MaxOpenConns": 20,
//		"MaxOdleConns": 5
//	},
//	"test": {
//		"username": "root",
//		"password": "root",
//		"network": "tcp",
//		"host": "192.168.1.101:3306",
//		"database": "mservice",
//		"MaxOpenConns": 20,
//		"MaxOdleConns": 5
//	},
//	"prod": {
//		"username": "root",
//		"password": "Aabbcc123!@#",
//		"network": "tcp",
//		"host": "47.96.230.81:3306",
//		"database": "mservice",
//		"MaxOpenConns": 20,
//		"MaxOdleConns": 5
//	}
//}

type MysqlConfigItemStruct struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Network      string `json:"network"`
	Host         string `json:"host"`
	Database     string `json:"database"`
	MaxOpenConns int64  `json:"MaxOpenConns"`
	MaxOdleConns int64  `json:"MaxOdleConns"`
}

type configStruct struct {
	Status  int64                 `json:"status"`
	Comment string                `json:"comment"`
	Dev     MysqlConfigItemStruct `json:"dev"`
	Test    MysqlConfigItemStruct `json:"test"`
	Prod    MysqlConfigItemStruct `json:"prod"`
}

type Config struct {
	Cfg MysqlConfigItemStruct
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

	if cfgstruct.Status == 2 {
		c.Cfg = cfgstruct.Prod
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
	return fmt.Sprintf("{Username:%s Password:%s Host:%s Network:%s Database:%s MaxOpenConns:%d MaxOdleConns:%d}",
		c.Cfg.Username, c.Cfg.Password, c.Cfg.Host, c.Cfg.Network, c.Cfg.Database, c.Cfg.MaxOpenConns, c.Cfg.MaxOdleConns)
}
