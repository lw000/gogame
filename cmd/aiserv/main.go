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
	// 测试日志写入服务
	go func() {
		for {
			err := rpcLoggerCli.WriteLogger("aiserv-1")
			if err != nil {

			}
			time.Sleep(time.Second * time.Duration(1))
		}
	}()
}

func main() {
	ggsys.RegisterOnInterrupt(func() {

	})

	if err := global.LoadGlobalConfig(); err != nil {
		log.Panic(err)
	}

	rpcLoggerCli = &rpcclient.RpcLoggerClient{}
	if err := rpcLoggerCli.Start(fmt.Sprintf("%s:%d", global.Cfg.LoggerServ.Host, global.Cfg.LoggerServ.Port)); err != nil {
		log.Panic(err)
	}

	Test()

	select {}
}
