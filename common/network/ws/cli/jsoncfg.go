package ggwscfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// {
// 	"scheme": "ws",
// 	"host": "47.96.230.81:8830",
// 	"path": ""
// }

type WsConfigStruct struct {
	Host string `json:"host"`
	Path string `json:"path"`
}

func LoadConfig(file string) (*WsConfigStruct, error) {
	cfg := &WsConfigStruct{}
	err := cfg.Load(file)
	return cfg, err
}

func (c *WsConfigStruct) Load(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, c); err != nil {
		return err
	}

	return nil
}

func (c WsConfigStruct) String() string {
	return fmt.Sprintf("{Host:%s Path:%s}", c.Host, c.Path)
}
