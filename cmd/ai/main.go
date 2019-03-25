package main

import (
	"demo/gogame/cmd/ai/global"
	"demo/gogame/cmd/ai/rpc"
	"demo/gogame/common/sys"
	"fmt"
	"log"
	"time"
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

	go func() {
		for {
			er := gr.RpcLoggerMgr().WriteLogger("aiserv test message")
			if er != nil {

			}
			time.Sleep(time.Second * time.Duration(1))
		}
	}()

	select {}
}
