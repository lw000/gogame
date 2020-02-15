package config

type CfgStruct struct {
	Debug       int64 `json:"debug"`
	LoggerServe struct {
		Host string `json:"host"`
		Port int64  `json:"port"`
	} `json:"loggerServe"`
	Mysql struct {
		MaxOdleConns int64  `json:"MaxOdleConns"`
		MaxOpenConns int64  `json:"MaxOpenConns"`
		Database     string `json:"database"`
		Host         string `json:"host"`
		Network      string `json:"network"`
		Password     string `json:"password"`
		Username     string `json:"username"`
	} `json:"mysql"`
	Port  int64 `json:"port"`
	Redis struct {
		DB           int64  `json:"db"`
		Host         string `json:"host"`
		MinIdleConns int64  `json:"minIdleConns"`
		PoolSize     int64  `json:"poolSize"`
		Psd          string `json:"psd"`
	} `json:"redis"`
}
