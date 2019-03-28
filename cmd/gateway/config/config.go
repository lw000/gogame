package config

import (
	"demo/gogame/common/db"
	"demo/gogame/common/rdsex"
	"encoding/json"
	"io/ioutil"
)

type JsonConfig struct {
	RdsCfg     *ggrdsex.RedisConfigItemStruct
	MysqlCfg   *ggdb.MysqlConfigItemStruct
	RouterServ struct {
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
		RdsCfg:   &ggrdsex.RedisConfigItemStruct{},
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

	cfg.LoggerServ.Host = cfgStruct.LoggerServ.Host
	cfg.LoggerServ.Port = cfgStruct.LoggerServ.Port

	cfg.RouterServ.Host = cfgStruct.RouterServ.Host
	cfg.RouterServ.Port = cfgStruct.RouterServ.Port

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

	switch cfgStruct.Redis.Status {
	case 0:
		cfg.RdsCfg.Host = cfgStruct.Redis.Dev.Host
		cfg.RdsCfg.Db = cfgStruct.Redis.Dev.DB
		cfg.RdsCfg.Psd = cfgStruct.Redis.Dev.Psd
		cfg.RdsCfg.PoolSize = cfgStruct.Redis.Dev.PoolSize
		cfg.RdsCfg.MinIdleConns = cfgStruct.Redis.Dev.MinIdleConns
	case 1:
		cfg.RdsCfg.Host = cfgStruct.Redis.Test.Host
		cfg.RdsCfg.Db = cfgStruct.Redis.Test.DB
		cfg.RdsCfg.Psd = cfgStruct.Redis.Test.Psd
		cfg.RdsCfg.PoolSize = cfgStruct.Redis.Test.PoolSize
		cfg.RdsCfg.MinIdleConns = cfgStruct.Redis.Test.MinIdleConns
	case 2:
		cfg.RdsCfg.Host = cfgStruct.Redis.Prod.Host
		cfg.RdsCfg.Db = cfgStruct.Redis.Prod.DB
		cfg.RdsCfg.Psd = cfgStruct.Redis.Prod.Psd
		cfg.RdsCfg.PoolSize = cfgStruct.Redis.Prod.PoolSize
		cfg.RdsCfg.MinIdleConns = cfgStruct.Redis.Prod.MinIdleConns
	default:
	}

	return cfg, err
}
