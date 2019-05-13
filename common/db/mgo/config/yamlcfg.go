package ggmgoconfig

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type YamlConfigStruct struct {
	Address  []string `yaml:"address"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
	Db       string   `yaml:"db"`
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
	return fmt.Sprintf("{Address:%v Username:%s Password:%s Db:%s}", c.Address, c.Username, c.Password, c.Db)
}
