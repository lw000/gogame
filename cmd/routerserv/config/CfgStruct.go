package config

type CfgStruct struct {
	Debug      int64 `json:"debug"`
	LoggerServ struct {
		Host string `json:"host"`
		Port int64  `json:"port"`
	} `json:"loggerServ"`
	Mysql struct {
		Comment string `json:"comment"`
		Dev     struct {
			MaxOdleConns int64  `json:"MaxOdleConns"`
			MaxOpenConns int64  `json:"MaxOpenConns"`
			Database     string `json:"database"`
			Host         string `json:"host"`
			Network      string `json:"network"`
			Password     string `json:"password"`
			Username     string `json:"username"`
		} `json:"dev"`
		Prod struct {
			MaxOdleConns int64  `json:"MaxOdleConns"`
			MaxOpenConns int64  `json:"MaxOpenConns"`
			Database     string `json:"database"`
			Host         string `json:"host"`
			Network      string `json:"network"`
			Password     string `json:"password"`
			Username     string `json:"username"`
		} `json:"prod"`
		Status int64 `json:"status"`
		Test   struct {
			MaxOdleConns int64  `json:"MaxOdleConns"`
			MaxOpenConns int64  `json:"MaxOpenConns"`
			Database     string `json:"database"`
			Host         string `json:"host"`
			Network      string `json:"network"`
			Password     string `json:"password"`
			Username     string `json:"username"`
		} `json:"test"`
	} `json:"mysql"`
	Port  int64 `json:"port"`
	Redis struct {
		Comment string `json:"comment"`
		Dev     struct {
			DB           int64  `json:"db"`
			Host         string `json:"host"`
			MinIdleConns int64  `json:"minIdleConns"`
			PoolSize     int64  `json:"poolSize"`
			Psd          string `json:"psd"`
		} `json:"dev"`
		Prod struct {
			DB           int64  `json:"db"`
			Host         string `json:"host"`
			MinIdleConns int64  `json:"minIdleConns"`
			PoolSize     int64  `json:"poolSize"`
			Psd          string `json:"psd"`
		} `json:"prod"`
		Status int64 `json:"status"`
		Test   struct {
			DB           int64  `json:"db"`
			Host         string `json:"host"`
			MinIdleConns int64  `json:"minIdleConns"`
			PoolSize     int64  `json:"poolSize"`
			Psd          string `json:"psd"`
		} `json:"test"`
	} `json:"redis"`
}
