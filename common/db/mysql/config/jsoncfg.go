package ggdbconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// {
// "username": "root",
// "password": "Aabbcc123!@#",
// "host": "47.96.230.81:3306",
// "database": "mservice",
// "MaxOpenConns": 20,
// "MaxOdleConns": 5
// }

type JsonConfigStruct struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Database     string `json:"database"`
	MaxOpenConns int64  `json:"MaxOpenConns"`
	MaxOdleConns int64  `json:"MaxOdleConns"`
}

// type Config struct {
// 	Cfg JsonConfigStruct
// }

func LoadJsonWithData(data []byte) (*JsonConfigStruct, error) {
	cfg := &JsonConfigStruct{}
	err := cfg.LoadWithData(data)
	return cfg, err
}

func LoadJson(file string) (*JsonConfigStruct, error) {
	cfg := &JsonConfigStruct{}
	err := cfg.Load(file)
	return cfg, err
}

func (c *JsonConfigStruct) LoadWithData(data []byte) error {
	if err := json.Unmarshal(data, c); err != nil {
		return err
	}
	return nil
}

func (c *JsonConfigStruct) Load(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return c.LoadWithData(data)
}

func (c JsonConfigStruct) String() string {
	return fmt.Sprintf("{Username:%s Password:%s Host:%s Database:%s MaxOpenConns:%d MaxOdleConns:%d}",
		c.Username, c.Password, c.Host, c.Database, c.MaxOpenConns, c.MaxOdleConns)
}
