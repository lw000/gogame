package config

import (
	"demo/gogame/common/db"
	"encoding/json"
	"io/ioutil"
)

type JsonConfig struct {
	MysqlCfg *ggdb.MysqlConfigItemStruct
	Port     int64
	Debug    int64
}

func NewJsonConfig() *JsonConfig {
	return &JsonConfig{
		MysqlCfg: &ggdb.MysqlConfigItemStruct{},
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

	switch cfgStruct.Mysql.Status {
	case 0:
		cfg.MysqlCfg.Database = cfgStruct.Mysql.Dev.Database
		cfg.MysqlCfg.Host = cfgStruct.Mysql.Dev.Host
		cfg.MysqlCfg.MaxOdleConns = cfgStruct.Mysql.Dev.MaxOdleConns
		cfg.MysqlCfg.MaxOpenConns = cfgStruct.Mysql.Dev.MaxOpenConns
		cfg.MysqlCfg.Network = cfgStruct.Mysql.Dev.Network
		cfg.MysqlCfg.Password = cfgStruct.Mysql.Dev.Password
		cfg.MysqlCfg.Username = cfgStruct.Mysql.Dev.Username
	case 1:
		cfg.MysqlCfg.Database = cfgStruct.Mysql.Test.Database
		cfg.MysqlCfg.Host = cfgStruct.Mysql.Test.Host
		cfg.MysqlCfg.MaxOdleConns = cfgStruct.Mysql.Test.MaxOdleConns
		cfg.MysqlCfg.MaxOpenConns = cfgStruct.Mysql.Test.MaxOpenConns
		cfg.MysqlCfg.Network = cfgStruct.Mysql.Test.Network
		cfg.MysqlCfg.Password = cfgStruct.Mysql.Test.Password
		cfg.MysqlCfg.Username = cfgStruct.Mysql.Test.Username
	case 2:
		cfg.MysqlCfg.Database = cfgStruct.Mysql.Prod.Database
		cfg.MysqlCfg.Host = cfgStruct.Mysql.Prod.Host
		cfg.MysqlCfg.MaxOdleConns = cfgStruct.Mysql.Prod.MaxOdleConns
		cfg.MysqlCfg.MaxOpenConns = cfgStruct.Mysql.Prod.MaxOpenConns
		cfg.MysqlCfg.Network = cfgStruct.Mysql.Prod.Network
		cfg.MysqlCfg.Password = cfgStruct.Mysql.Prod.Password
		cfg.MysqlCfg.Username = cfgStruct.Mysql.Prod.Username
	default:
	}

	return cfg, err
}
