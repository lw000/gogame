package config

type CfgStruct struct {
	Debug int64 `json:"debug"`
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
	Port int64 `json:"port"`
}
