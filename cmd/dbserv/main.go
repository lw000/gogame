package main

import (
	"demo/gogame/cmd/dbserv/global"
	"demo/gogame/cmd/dbserv/rpc"
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

	gr := rpc.DbRpc{}
	if er := gr.Start(); er != nil {
		log.Panic(er)
	}
	gr.StartRpcDbServer(global.Cfg.Port)
	gr.StartRpcLoggerClient(fmt.Sprintf("%s:%d", global.Cfg.LoggerServ.Host, global.Cfg.LoggerServ.Port))

	go func() {
		for {
			er := gr.RpcLoggerMgr().WriteLogger("dbserv")
			if er != nil {

			}
			time.Sleep(time.Second * time.Duration(1))
		}
	}()

	select {}
}
