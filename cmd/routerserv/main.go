package main

import (
	"demo/gogame/cmd/gateway/global"
	"demo/gogame/constant"
	"demo/gogame/proto/router"
	"demo/gogame/rpc/client"
	"demo/gogame/rpc/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

var (
	rpcLoggerCli  *rpcclient.RpcLoggerClient
	rpcRouterServ *rpcservice.RpcRouterServer

	services sync.Map
)

func Test() {
	//测试日志写入服务
	go func() {
		for {
			er := rpcLoggerCli.WriteLogger("routerserv-1")
			if er != nil {
				log.Println("error")
			}
			time.Sleep(time.Second * time.Duration(1))
		}
	}()
}

func main() {
	if er := global.LoadGlobalConfig(); er != nil {
		log.Panic(er)
	}

	rpcLoggerCli = &rpcclient.RpcLoggerClient{}
	if er := rpcLoggerCli.Start(fmt.Sprintf("%s:%d", global.Cfg.LoggerServ.Host, global.Cfg.LoggerServ.Port)); er != nil {
		log.Panic(er)
	}

	rpcRouterServ = &rpcservice.RpcRouterServer{ServiceId: ggconstant.CRouterServiceId}
	rpcRouterServ.HandleConnect(func(stream *rpcservice.RpcRouterServerStream) {
		log.Println("子服务器连接", stream)
		services.Store(stream.Uuid(), stream)
	})

	rpcRouterServ.HandleDisConnected(func(stream *rpcservice.RpcRouterServerStream) {
		log.Println("子服务器断开", stream)
		services.Delete(stream.Uuid())
	})

	rpcRouterServ.HandleMessage(func(stream *rpcservice.RpcRouterServerStream, req *routersvr.RequestMessage) {
		log.Println(req)

		switch req.ServiceId {
		case ggconstant.CGatewayServiceId:
			er := stream.SendMessage(req.ServiceId, req.Uuid, req.Msg)
			if er != nil {
				log.Println(er)
			}
		case ggconstant.CPlatformServiceId:
			er := stream.SendMessage(req.ServiceId, req.Uuid, req.Msg)
			if er != nil {
				log.Println(er)
			}
		}
	})

	rpcSvr := &rpcservice.RpcServer{}
	if er := rpcSvr.StartService(global.Cfg.Port, func(s *grpc.Server) {
		routersvr.RegisterRouterServer(s, rpcRouterServ)
	}); er != nil {
		log.Panic(er)
	}

	Test()

	select {}
}
