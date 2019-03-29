package main

import (
	"demo/gogame/cmd/aiserv/global"
	"demo/gogame/common/sys"
	"demo/gogame/rpc/client"
	"fmt"
	"log"
	"time"
)

var (
	rpcLoggerCli *rpcclient.RpcLoggerClient
)

func Test() {
	//测试日志写入服务
	go func() {
		for {
			er := rpcLoggerCli.WriteLogger("aiserv-1")
			if er != nil {

			}
			time.Sleep(time.Second * time.Duration(1))
		}
	}()
}

func main() {
	ggsys.RegisterOnInterrupt(func() {

	})

	if er := global.LoadGlobalConfig(); er != nil {
		log.Panic(er)
	}

	rpcLoggerCli = &rpcclient.RpcLoggerClient{}
	if er := rpcLoggerCli.Start(fmt.Sprintf("%s:%d", global.Cfg.LoggerServ.Host, global.Cfg.LoggerServ.Port)); er != nil {
		log.Panic(er)
	}

	Test()

	select {}
}
