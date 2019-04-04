package config

type CfgStruct struct {
	Debug int64 `json:"debug"`
	Mysql struct {
		MaxOdleConns int64  `json:"MaxOdleConns"`
		MaxOpenConns int64  `json:"MaxOpenConns"`
		Database     string `json:"database"`
		Host         string `json:"host"`
		Network      string `json:"network"`
		Password     string `json:"password"`
		Username     string `json:"username"`
	} `json:"mysql"`
	Port int64 `json:"port"`
}
