package global

import (
	"demo/gogame/cmd/platformsrv/config"
	log "github.com/alecthomas/log4go"
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
