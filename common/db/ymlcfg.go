package tydb

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type YamlMysqlCfg struct {
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Database     string `yaml:"database"`
	MaxOpenConns int64  `yaml:"maxOpenConns"`
	MaxOdleConns int64  `yaml:"maxOdleConns"`
}

func LoadConfigWithYamlData(data []byte) (*YamlMysqlCfg, error) {
	cfg := &YamlMysqlCfg{}
	err := cfg.LoadWithData(data)
	return cfg, err
}

func LoadConfigWithYaml(file string) (*YamlMysqlCfg, error) {
	cfg := &YamlMysqlCfg{}
	err := cfg.Load(file)
	return cfg, err
}

func (c *YamlMysqlCfg) LoadWithData(data []byte) error {
	err := yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}

	return nil
}

func (c *YamlMysqlCfg) Load(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return c.LoadWithData(data)
}

func (c YamlMysqlCfg) String() string {
	return fmt.Sprintf("{Username:%s Password:%s Host:%s Database:%s MaxOpenConns:%d MaxOdleConns:%d}",
		c.Username, c.Password, c.Host, c.Database, c.MaxOpenConns, c.MaxOdleConns)
}
