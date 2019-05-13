package config

import (
	ggdbconfig "demo/gogame/common/db/mysql/config"
	ggrdsexconfig "demo/gogame/common/db/rdsex/config"
	"encoding/json"
	"io/ioutil"
)

type JsonConfig struct {
	RdsCfg   *ggrdsexconfig.JsonConfigStruct
	MysqlCfg *ggdbconfig.JsonConfigStruct
	GateWay  struct {
		Host string
		Port int64
	}
	LoggerServ struct {
		Host string
		Port int64
	}
	Port  int64
	Debug int64
}

func NewJsonConfig() *JsonConfig {
	return &JsonConfig{
		RdsCfg:   &ggrdsexconfig.JsonConfigStruct{},
		MysqlCfg: &ggdbconfig.JsonConfigStruct{},
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

	cfg.LoggerServ.Host = cfgStruct.LoggerServ.Host
	cfg.LoggerServ.Port = cfgStruct.LoggerServ.Port

	cfg.MysqlCfg.Database = cfgStruct.Mysql.Database
	cfg.MysqlCfg.Host = cfgStruct.Mysql.Host
	cfg.MysqlCfg.MaxOdleConns = cfgStruct.Mysql.MaxOdleConns
	cfg.MysqlCfg.MaxOpenConns = cfgStruct.Mysql.MaxOpenConns
	cfg.MysqlCfg.Password = cfgStruct.Mysql.Password
	cfg.MysqlCfg.Username = cfgStruct.Mysql.Username

	cfg.RdsCfg.Host = cfgStruct.Redis.Host
	cfg.RdsCfg.Db = cfgStruct.Redis.DB
	cfg.RdsCfg.Psd = cfgStruct.Redis.Psd
	cfg.RdsCfg.PoolSize = cfgStruct.Redis.PoolSize
	cfg.RdsCfg.MinIdleConns = cfgStruct.Redis.MinIdleConns

	return cfg, err
}
