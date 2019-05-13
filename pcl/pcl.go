package ggpcl

type PclStruct struct {
	ServiceCmd struct {
		MainID int64 `json:"mainId"`
		SubIds []struct {
			Dest  string `json:"dest"`
			SubID int64  `json:"subId"`
		} `json:"subIds"`
	} `json:"serviceCmd"`
	ServiceID      int64  `json:"serviceId"`
	ServiceName    string `json:"serviceName"`
	ServiceVersion string `json:"serviceVersion"`
}
