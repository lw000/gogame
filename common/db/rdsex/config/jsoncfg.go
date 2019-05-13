package ggrdsexconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// {
// "host": "127.0.0.1:6379",
// "psd": "123456",
// "db": 0,
// "poolSize": 20,
// "minIdleConns": 5
// }

type JsonConfigStruct struct {
	Host         string `json:"host"`
	Psd          string `json:"psd"`
	Db           int64  `json:"db"`
	PoolSize     int64  `json:"poolSize"`
	MinIdleConns int64  `json:"minIdleConns"`
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

	if err = json.Unmarshal(data, c); err != nil {
		return err
	}

	return nil
}

func (c JsonConfigStruct) String() string {
	return fmt.Sprintf("{Host:%s Psd:%s Db:%d PoolSize:%d MinIdleConns:%d}", c.Host, c.Psd, c.Db, c.PoolSize, c.MinIdleConns)
}
