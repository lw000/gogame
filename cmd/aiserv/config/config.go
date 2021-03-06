package config

import (
	"encoding/json"
	"io/ioutil"
)

type JsonConfig struct {
	GateWay struct {
		Host string
		Port int64
	}
	LoggerServe struct {
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

	cfg.GateWay.Host = cfgStruct.GateWay.Host
	cfg.GateWay.Port = cfgStruct.GateWay.Port

	return cfg, err
}
