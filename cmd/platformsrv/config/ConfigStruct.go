package config

type CfgStruct struct {
	DBServe struct {
		Host string `json:"host"`
		Port int64  `json:"port"`
	} `json:"dbServe"`
	Debug       int64 `json:"debug"`
	LoggerServe struct {
		Host string `json:"host"`
		Port int64  `json:"port"`
	} `json:"loggerServe"`
	RouterWay struct {
		Host string `json:"host"`
		Port int64  `json:"port"`
	} `json:"routerWay"`
}
