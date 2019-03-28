package main

import (
	"demo/gogame/cmd/gateway/global"
	"demo/gogame/common/utilty"
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
	rpcRouterServ.HandleConnect(func(stream routersvr.Router_ForwardingDataStreamServer) {
		services.Store(ggutilty.UUID(), stream)
	})

	rpcRouterServ.HandleDisConnected(func(stream routersvr.Router_ForwardingDataStreamServer) {

	})

	rpcRouterServ.HandleRegisterService(func(protocols []*routersvr.RouterProtocol) error {

		return nil
	})

	rpcRouterServ.HandleMessage(func(stream routersvr.Router_ForwardingDataStreamServer, req *routersvr.ForwardMessage) {
		log.Println(req)
		switch req.ServiceId {
		case ggconstant.CRouterServiceId:
			er := stream.Send(&routersvr.ForwardMessage{ServiceId: req.ServiceId, Uuid: req.Uuid, Msg: req.Msg})
			if er != nil {
				log.Println(er)
			}
		case ggconstant.CGatewayServiceId:
			er := stream.Send(&routersvr.ForwardMessage{ServiceId: req.ServiceId, Uuid: req.Uuid, Msg: req.Msg})
			if er != nil {
				log.Println(er)
			}
		case ggconstant.CPlatformServiceId:
			er := stream.Send(&routersvr.ForwardMessage{ServiceId: req.ServiceId, Uuid: req.Uuid, Msg: req.Msg})
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
