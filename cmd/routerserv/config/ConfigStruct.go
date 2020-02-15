package config

type CfgStruct struct {
	Debug       int64 `json:"debug"`
	LoggerServe struct {
		Host string `json:"host"`
		Port int64  `json:"port"`
	} `json:"loggerServe"`
	Port int64 `json:"port"`
}
