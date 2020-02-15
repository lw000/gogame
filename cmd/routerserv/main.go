package main

import (
	"demo/gogame/cmd/gateway/global"
	"demo/gogame/pcl"
	"demo/gogame/protos/router"
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
	// 测试日志写入服务
	go func() {
		for {
			err := rpcLoggerCli.WriteLogger("routerserv-1")
			if err != nil {
				log.Println("error")
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
		case 0: // 注册协议
			pcls, err := ggpcl.DecodePcl(req.Msg)
			if err != nil {
				log.Println(err)
				err = stream.SendMessage(req.ServiceId, req.MsgType, req.Uuid, []byte("0"))
				if err != nil {
					log.Println(err)
				}
				return
			}

			log.Printf("%+v", pcls)

			clientConfig := ClientSession{ServiceId: req.ServiceId, Clientuuid: stream.ClientUuid()}
			clientConfig.protocols.Store(req.ServiceVersion, pcls)
			clientStreams.Store(req.ServiceId, clientConfig)

			err = stream.SendMessage(req.ServiceId, req.MsgType, req.Uuid, []byte("1"))
			if err != nil {
				log.Println(err)
			}
		case 1: // 转发消息
			v1, ok := clientStreams.Load(req.Cuuid)
			if !ok {
				log.Println("没有找到路由信息")
				return
			}
			transStream := v1.(*rpcservice.RpcRouterServerStream)
			err := transStream.SendMessage(req.ServiceId, 1, req.Uuid, req.Msg)
			if err != nil {
				log.Println(err)
			}
		}
	})

	rpcSvr := &rpcservice.RpcServer{}
	if err := rpcSvr.StartService(global.Cfg.Port, func(s *grpc.Server) {
		routersvr.RegisterRouterServer(s, rpcRouterServ)
	}); err != nil {
		log.Panic(err)
	}

	Test()

	select {}
}
