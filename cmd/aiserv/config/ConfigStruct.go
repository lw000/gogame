package config

type CfgStruct struct {
	Debug   int64 `json:"debug"`
	GateWay struct {
		Host string `json:"host"`
		Port int64  `json:"port"`
	} `json:"gateWay"`
	LoggerServ struct {
		Host string `json:"host"`
		Port int64  `json:"port"`
	} `json:"loggerServ"`
}
