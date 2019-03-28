package global

import (
	"demo/gogame/cmd/gateway/config"
	log "github.com/alecthomas/log4go"
)

var (
	Cfg *config.JsonConfig
)

func LoadGlobalConfig() error {
	log.LoadConfiguration("./conf/log4go.xml")

	var er error
	Cfg, er = config.LoadJsonConfig("./conf/conf.json")
	if er != nil {
		log.Error(er)
		return er
	}

	return nil
}
