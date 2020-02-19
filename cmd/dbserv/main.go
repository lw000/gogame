package main

import (
	"fmt"
	"gogame/cmd/dbserv/global"
	"gogame/protos/db"
	"gogame/rpc/client"
	"gogame/rpc/service"
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
			err := rpcLoggerCli.WriteLogger("dbserv-1")
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

	rpcDbSvr = &rpcservice.RpcDbServer{}
	rpcDbSvr.HandleMessage(func(stream dbsvr.DB_BidStreamServer, req *dbsvr.Request) {
		log.Println(req)

		switch req.MainId {
		case 0:
			err := stream.Send(&dbsvr.Response{MainId: req.MainId, SubId: req.SubId, RequestId: req.RequestId, Output: req.Input})
			if err != nil {
				log.Println(err)
			}
		case 1:
			err := stream.Send(&dbsvr.Response{MainId: req.MainId, SubId: req.SubId, RequestId: req.RequestId, Output: req.Input})
			if err != nil {
				log.Println(err)
			}
		}
	})

	rpcSvr := &rpcservice.RpcServer{}
	if err := rpcSvr.StartService(global.Cfg.Port, func(s *grpc.Server) {
		dbsvr.RegisterDBServer(s, rpcDbSvr)
	}); err != nil {
		log.Panic(err)
	}

	rpcLoggerCli = &rpcclient.RpcLoggerClient{}
	if err := rpcLoggerCli.Start(fmt.Sprintf("%s:%d", global.Cfg.LoggerServe.Host, global.Cfg.LoggerServe.Port)); err != nil {
		log.Panic(err)
	}

	Test()

	select {}
}
