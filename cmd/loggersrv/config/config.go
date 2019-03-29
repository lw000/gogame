package config

import (
	"demo/gogame/common/db"
	"encoding/json"
	"io/ioutil"
)

type JsonConfig struct {
	MysqlCfg *ggdb.MysqlConfigStruct
	Port     int64
	Debug    int64
}

func NewJsonConfig() *JsonConfig {
	return &JsonConfig{
		MysqlCfg: &ggdb.MysqlConfigStruct{},
	}
}

func LoadJsonConfig(file string) (*JsonConfig, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var cfgStruct CfgStruct
	if err = json.Unmarshal(data, &cfgStruct); err != nil {
		return nil, err
	}

	cfg := NewJsonConfig()
	cfg.Debug = cfgStruct.Debug
	cfg.Port = cfgStruct.Port

	cfg.MysqlCfg.Database = cfgStruct.Mysql.Database
	cfg.MysqlCfg.Host = cfgStruct.Mysql.Host
	cfg.MysqlCfg.MaxOdleConns = cfgStruct.Mysql.MaxOdleConns
	cfg.MysqlCfg.MaxOpenConns = cfgStruct.Mysql.MaxOpenConns
	cfg.MysqlCfg.Password = cfgStruct.Mysql.Password
	cfg.MysqlCfg.Username = cfgStruct.Mysql.Username

	return cfg, err
}
