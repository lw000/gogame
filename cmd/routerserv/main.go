package main

import (
	"demo/gogame/cmd/gateway/global"
	"demo/gogame/cmd/routerserv/rpc/service"
	"demo/gogame/proto/router"
	"demo/gogame/rpc/client"
	"demo/gogame/rpc/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

var (
	rpcLoggerCli  *rpcclient.RpcLoggerClient
	rpcRouterServ *service.RpcRouterServer
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

	rpcRouterServ = &service.RpcRouterServer{}
	rpcRouterServ.HandleMessage(func(stream routersvr.Router_BidStreamServer, req *routersvr.ForwardRequest) {
		log.Println(req)

		switch req.MainId {
		case 10000:
			er := stream.Send(&routersvr.ForwardResponse{MainId: req.MainId, SubId: req.SubId, Uuid: req.Uuid, Output: req.Input})
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

	rpcLoggerCli = &rpcclient.RpcLoggerClient{}
	if er := rpcLoggerCli.Start(fmt.Sprintf("%s:%d", global.Cfg.LoggerServ.Host, global.Cfg.LoggerServ.Port)); er != nil {
		log.Panic(er)
	}

	Test()

	select {}
}
