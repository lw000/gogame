package main

import (
	"fmt"
	"gogame/cmd/gamesrv/global"
	"gogame/protos/db"
	"gogame/protos/router"
	"gogame/rpc/client"
	"log"
	"sync/atomic"
	"time"
)

var (
	rpcLoggerCli *rpcclient.RpcLoggerClient
	rpcRouterCli *rpcclient.RpcRouterClient
	rpcDbCli     *rpcclient.RpcDbClient

	rpcRouterStream *rpcclient.RpcRouterStreamClient
	rpcDbStream     *rpcclient.RpcDbStream
)

func Test() {
	// 网关数据发送测试
	go func() {
		for {
			var requestId int32 = 0
			atomic.AddInt32(&requestId, 1)
			err := rpcRouterStream.SendMessage("", []byte("gamesrv-1"))
			if err != nil {
				log.Println(err)
				return
			}
			time.Sleep(time.Second * time.Duration(1))
		}
	}()

	// 测试日志写入服务
	go func() {
		for {
			err := rpcLoggerCli.WriteLogger("gamesrv-1")
			if err != nil {

			}
			time.Sleep(time.Second * time.Duration(1))
		}
	}()

	// 测试数据库服务
	go func() {
		for {
			var requestId int32 = 0
			atomic.AddInt32(&requestId, 1)
			err := rpcDbStream.SendMessage(1, 10000, requestId, "gamesrv-1")
			if err != nil {
				log.Println(err)
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

	rpcRouterCli = &rpcclient.RpcRouterClient{}
	if err := rpcRouterCli.Start(fmt.Sprintf("%s:%d", global.Cfg.RouterWay.Host, global.Cfg.RouterWay.Port)); err != nil {
		log.Panic(err)
	}

	var err error
	rpcRouterStream, err = rpcRouterCli.CreateStream(func(response *routersvr.ReponseMessage) {
		switch response.ServiceId {
		case 1:
			log.Println(response)
		case 2:
			log.Println(response)
		}
	})
	if err != nil {
		log.Panic(err)
	}

	rpcDbCli = &rpcclient.RpcDbClient{}
	if err = rpcDbCli.Start(fmt.Sprintf("%s:%d", global.Cfg.DBServe.Host, global.Cfg.DBServe.Port)); err != nil {
		log.Panic(err)
	}

	rpcDbStream, err = rpcDbCli.CreateStream(func(response *dbsvr.Response) {
		switch response.MainId {
		case 1:
			log.Println(response)
		case 2:
			log.Println(response)
		}
	})
	if err != nil {
		log.Panic(err)
	}

	Test()

	select {}
}
