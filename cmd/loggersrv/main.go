package main

import (
	"demo/gogame/cmd/loggersrv/global"
	"demo/gogame/cmd/loggersrv/rpc"
	"log"
)

func main() {
	if er := global.LoadGlobalConfig(); er != nil {
		log.Panic(er)
	}

	l := rpc.LoggerRpc{}
	if er := l.Start(global.Cfg.Port); er != nil {
		log.Panic(er)
	}

	select {}
}
