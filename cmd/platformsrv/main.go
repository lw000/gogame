package main

import (
	"demo/gogame/cmd/platformsrv/global"
	"demo/gogame/cmd/platformsrv/htp"
	"demo/gogame/cmd/platformsrv/platform"
	"demo/gogame/common/sys"
	"demo/gogame/proto/center"
	"demo/gogame/proto/db"
	"demo/gogame/proto/platform"
	"demo/gogame/rpc/client"
	"demo/gogame/rpc/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"sync/atomic"
	"time"
)

var (
	plat *platform.Platform

	rpcPlatformSvr *rpcservice.RpcPlatformServer

	rpcLoggerCli *rpcclient.RpcLoggerClient
	rpcCenterCli *rpcclient.RpcCenterClient
	rpcDbCli     *rpcclient.RpcDbClient

	rpcCenterStream *rpcclient.RpcCenterStream
	rpcDbStream     *rpcclient.RpcDbStream
)

func Test() {
	//测试日志写入服务
	go func() {
		for {
			er := rpcLoggerCli.WriteLogger("platform-1")
			if er != nil {

			}
			time.Sleep(time.Second * time.Duration(1))
		}
	}()

	//网关数据发送测试
	go func() {
		for {
			var requestId int32 = 0
			atomic.AddInt32(&requestId, 1)
			er := rpcCenterStream.SendMessage(1, 10000, requestId, "platform-1")
			if er != nil {
				log.Println(er)
				return
			}
			time.Sleep(time.Second * time.Duration(1))
		}
	}()

	//测试数据库服务
	go func() {
		for {
			var requestId int32 = 0
			atomic.AddInt32(&requestId, 1)
			er := rpcDbStream.SendMessage(1, 10000, requestId, "platform-1")
			if er != nil {
				log.Println(er)
			}
			time.Sleep(time.Second * time.Duration(1))
		}
	}()
}

func main() {
	ggsys.RegisterOnInterrupt(func() {
		if er := plat.Stop(); er != nil {

		}
	})

	plat = platform.NewPlatform(1, "棋牌游戏")
	if er := plat.Start(); er != nil {
		log.Panic(er)
	}

	if er := global.LoadGlobalConfig(); er != nil {
		log.Panic(er)
	}

	pfh := htp.PlatformHtp{}
	if er := pfh.Start(global.Cfg.HTTPPort); er != nil {
		log.Panic(er)
	}

	rpcPlatformSvr = &rpcservice.RpcPlatformServer{}
	rpcPlatformSvr.HandleMessage(func(stream platformsvr.Platform_BidStreamServer, req *platformsvr.Request) {
		log.Println(req)

		switch req.MainId {
		case 0:
			er := stream.Send(&platformsvr.Response{MainId: req.MainId, SubId: req.SubId, RequestId: req.RequestId, Output: req.Input})
			if er != nil {
				log.Println(er)
			}
		case 1:
			er := stream.Send(&platformsvr.Response{MainId: req.MainId, SubId: req.SubId, RequestId: req.RequestId, Output: req.Input})
			if er != nil {
				log.Println(er)
			}
		}
	})
	rpcSvr := &rpcservice.RpcServer{}
	if er := rpcSvr.StartService(global.Cfg.Port, func(serv *grpc.Server) {
		platformsvr.RegisterPlatformServer(serv, rpcPlatformSvr)
	}); er != nil {
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

	rpcDbCli = &rpcclient.RpcDbClient{}
	if er := rpcDbCli.Start(fmt.Sprintf("%s:%d", global.Cfg.DBServ.Host, global.Cfg.DBServ.Port)); er != nil {
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
