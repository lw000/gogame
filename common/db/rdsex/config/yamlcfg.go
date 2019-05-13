package ggrdsexconfig

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type YamlConfigStruct struct {
	Host         string `yaml:"host"`
	Psd          string `yaml:"psd"`
	Db           int64  `yaml:"db"`
	PoolSize     int64  `yaml:"poolSize"`
	MinIdleConns int64  `yaml:"minIdleConns"`
}

func LoadWithYaml(file string) (*YamlConfigStruct, error) {
	cfg := &YamlConfigStruct{}
	err := cfg.Load(file)
	return cfg, err
}

func (c *YamlConfigStruct) Load(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	er := yaml.Unmarshal(data, c)
	if er != nil {
		return er
	}

	return nil
}

func (c YamlConfigStruct) String() string {
	return fmt.Sprintf("{Host:%s Psd:%s Db:%d PoolSize:%d MinIdleConns:%d}", c.Host, c.Psd, c.Db, c.PoolSize, c.MinIdleConns)
}
