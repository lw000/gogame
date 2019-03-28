package main

import (
	"demo/gogame/cmd/gamesrv/global"
	"demo/gogame/common/sys"
	"demo/gogame/proto/center"
	"demo/gogame/proto/db"
	"demo/gogame/rpc/client"
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

var (
	rpcLoggerCli *rpcclient.RpcLoggerClient
	rpcCenterCli *rpcclient.RpcCenterClient
	rpcDbCli     *rpcclient.RpcDbClient

	rpcCenterStream *rpcclient.RpcCenterStream
	rpcDbStream     *rpcclient.RpcDbStream
)

func Test() {
	//网关数据发送测试
	go func() {
		for {
			var requestId int32 = 0
			atomic.AddInt32(&requestId, 1)
			er := rpcCenterStream.SendMessage(1, 10000, requestId, "gamesrv-1")
			if er != nil {
				log.Println(er)
				return
			}
			time.Sleep(time.Second * time.Duration(1))
		}
	}()

	//测试日志写入服务
	go func() {
		for {
			er := rpcLoggerCli.WriteLogger("gamesrv-1")
			if er != nil {

			}
			time.Sleep(time.Second * time.Duration(1))
		}
	}()

	//测试数据库服务
	go func() {
		for {
			var requestId int32 = 0
			atomic.AddInt32(&requestId, 1)
			er := rpcDbStream.SendMessage(1, 10000, requestId, "gamesrv-1")
			if er != nil {
				log.Println(er)
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

	rpcCenterCli = &rpcclient.RpcCenterClient{}
	if er := rpcCenterCli.Start(fmt.Sprintf("%s:%d", global.Cfg.GateWay.Host, global.Cfg.GateWay.Port)); er != nil {
		log.Panic(er)
	}

	var er error
	rpcCenterStream, er = rpcCenterCli.CreateStream(func(response *centersvr.Response) {
		switch response.MainId {
		case 1:
			log.Println(response)
		case 2:
			log.Println(response)
		}
	})
	if er != nil {
		log.Panic(er)
	}

	rpcDbCli = &rpcclient.RpcDbClient{}
	if er = rpcDbCli.Start(fmt.Sprintf("%s:%d", global.Cfg.DBServ.Host, global.Cfg.DBServ.Port)); er != nil {
		log.Panic(er)
	}

	rpcDbStream, er = rpcDbCli.CreateStream(func(response *dbsvr.Response) {
		switch response.MainId {
		case 1:
			log.Println(response)
		case 2:
			log.Println(response)
		}
	})
	if er != nil {
		log.Panic(er)
	}

	Test()

	select {}
}
