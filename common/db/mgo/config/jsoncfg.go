package ggmgoconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// {
// "address": "127.0.0.1:6379",
// username:"levi"
// "password": "123456",
// "db": "log",
// }

type JsonConfigStruct struct {
	Address  []string `json:"address"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Db       string   `json:"db"`
}

func LoadConfig(file string) (*JsonConfigStruct, error) {
	cfg := &JsonConfigStruct{}
	err := cfg.Load(file)
	return cfg, err
}

func (c *JsonConfigStruct) Load(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, &c); err != nil {
		return err
	}

	return nil
}

func (c JsonConfigStruct) String() string {
	return fmt.Sprintf("{Address:%v Username:%s Password:%s Db:%s}", c.Address, c.Username, c.Password, c.Db)
}
