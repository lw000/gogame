package main

import (
	"demo/gogame/cmd/ai/global"
	"demo/gogame/cmd/ai/rpc"
	"demo/gogame/common/sys"
	"fmt"
	"log"
)

func main() {
	ggsys.RegisterOnInterrupt(func() {

	})

	if er := global.LoadGlobalConfig(); er != nil {
		log.Panic(er)
	}

	gr := rpc.AiRpc{}
	if er := gr.Start(); er != nil {
		log.Panic(er)
	}

	gr.StartRpcLoggerClient(fmt.Sprintf("%s:%d", global.Cfg.LoggerServ.Host, global.Cfg.LoggerServ.Port))

	select {}
}
