package config

type CfgStruct struct {
	Debug      int64 `json:"debug"`
	LoggerServ struct {
		Host string `json:"host"`
		Port int64  `json:"port"`
	} `json:"loggerServ"`
	Port       int64 `json:"port"`
	RouterServ struct {
		Host string `json:"host"`
		Port int64  `json:"port"`
	} `json:"routerServ"`
}
