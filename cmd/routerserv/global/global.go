package global

import (
	log "github.com/alecthomas/log4go"
	"gogame/cmd/routerserv/config"
)

var (
	Cfg *config.JsonConfig
)

func LoadGlobalConfig() error {
	log.LoadConfiguration("./conf/log4go.xml")

	var err error
	Cfg, err = config.LoadJsonConfig("./conf/conf.json")
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
