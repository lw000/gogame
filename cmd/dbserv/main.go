package main

import (
	"demo/gogame/cmd/dbserv/global"
	"demo/gogame/common/sys"
	"demo/gogame/proto/db"
	"demo/gogame/rpc/client"
	"demo/gogame/rpc/service"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

var (
	rpcDbSvr     *rpcservice.RpcDbServer
	rpcLoggerCli *rpcclient.RpcLoggerClient
)

func Test() {
	go func() {
		for {
			er := rpcLoggerCli.WriteLogger("dbserv-1")
			if er != nil {

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

	rpcDbSvr = &rpcservice.RpcDbServer{}
	rpcDbSvr.HandleMessage(func(stream dbsvr.DB_BidStreamServer, req *dbsvr.Request) {
		log.Println(req)

		switch req.MainId {
		case 0:
			er := stream.Send(&dbsvr.Response{MainId: req.MainId, SubId: req.SubId, RequestId: req.RequestId, Output: req.Input})
			if er != nil {
				log.Println(er)
			}
		case 1:
			er := stream.Send(&dbsvr.Response{MainId: req.MainId, SubId: req.SubId, RequestId: req.RequestId, Output: req.Input})
			if er != nil {
				log.Println(er)
			}
		}
	})

	rpcSvr := &rpcservice.RpcServer{}
	if er := rpcSvr.StartService(global.Cfg.Port, func(s *grpc.Server) {
		dbsvr.RegisterDBServer(s, rpcDbSvr)
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
