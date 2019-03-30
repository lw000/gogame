package ggpcl

import (
	"encoding/json"
	"io/ioutil"
)

func LoadPcl(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func DecodePcl(data []byte) (*PclStruct, error) {
	var cfgStruct PclStruct
	if er := json.Unmarshal(data, &cfgStruct); er != nil {
		return nil, er
	}
	return &cfgStruct, nil
}
