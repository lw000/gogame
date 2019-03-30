package main

import (
	"demo/gogame/cmd/gateway/global"
	"demo/gogame/pcl"
	"demo/gogame/proto/router"
	"demo/gogame/rpc/client"
	"demo/gogame/rpc/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

type ClientSession struct {
	ServiceId  int32
	Clientuuid string
	protocols  sync.Map
}

var (
	rpcLoggerCli  *rpcclient.RpcLoggerClient
	rpcRouterServ *rpcservice.RpcRouterServer

	clientStreams sync.Map
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

	rpcRouterServ = &rpcservice.RpcRouterServer{}
	rpcRouterServ.HandleConnect(func(stream *rpcservice.RpcRouterServerStream) {
		log.Println("子服务器连接", stream)
	})

	rpcRouterServ.HandleDisConnected(func(stream *rpcservice.RpcRouterServerStream) {
		clientStreams.Delete(stream.ServiceId)
		log.Println("子服务器断开", stream)
	})

	rpcRouterServ.HandleMessage(func(stream *rpcservice.RpcRouterServerStream, req *routersvr.RequestMessage) {
		switch req.MsgType {
		case 0: //注册协议
			pcls, er := ggpcl.DecodePcl(req.Msg)
			if er != nil {
				log.Println(er)
				er = stream.SendMessage(req.ServiceId, req.MsgType, req.Uuid, []byte("0"))
				if er != nil {
					log.Println(er)
				}
				return
			}

			log.Printf("%+v", pcls)

			clientConfig := ClientSession{ServiceId: req.ServiceId, Clientuuid: stream.ClientUuid()}
			clientConfig.protocols.Store(req.ServiceVersion, pcls)
			clientStreams.Store(req.ServiceId, clientConfig)

			er = stream.SendMessage(req.ServiceId, req.MsgType, req.Uuid, []byte("1"))
			if er != nil {
				log.Println(er)
			}
		case 1: //转发消息
			v1, ok := clientStreams.Load(req.Cuuid)
			if !ok {
				log.Println("没有找到路由信息")
				return
			}
			transStream := v1.(*rpcservice.RpcRouterServerStream)
			er := transStream.SendMessage(req.ServiceId, 1, req.Uuid, req.Msg)
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
