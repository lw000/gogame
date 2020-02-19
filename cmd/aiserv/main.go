package main

import (
	"fmt"
	"gogame/cmd/aiserv/global"
	"gogame/rpc/client"
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

	if err := global.LoadGlobalConfig(); err != nil {
		log.Panic(err)
	}

	rpcLoggerCli = &rpcclient.RpcLoggerClient{}
	if err := rpcLoggerCli.Start(fmt.Sprintf("%s:%d", global.Cfg.LoggerServe.Host, global.Cfg.LoggerServe.Port)); err != nil {
		log.Panic(err)
	}

	Test()

	select {}
}
