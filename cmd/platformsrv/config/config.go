package config

import (
	"encoding/json"
	"io/ioutil"
)

type JsonConfig struct {
	RouterWay struct {
		Host string
		Port int64
	}
	LoggerServe struct {
		Host string
		Port int64
	}
	DBServe struct {
		Host string
		Port int64
	}
	Debug int64
}

func NewJsonConfig() *JsonConfig {
	return &JsonConfig{}
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

	cfg.LoggerServe.Host = cfgStruct.LoggerServe.Host
	cfg.LoggerServe.Port = cfgStruct.LoggerServe.Port

	cfg.RouterWay.Host = cfgStruct.RouterWay.Host
	cfg.RouterWay.Port = cfgStruct.RouterWay.Port

	cfg.DBServe.Host = cfgStruct.DBServe.Host
	cfg.DBServe.Port = cfgStruct.DBServe.Port

	return cfg, err
}
