package main

import (
	"demo/gogame/cmd/loggersrv/global"
	"demo/gogame/proto/logger"
	"demo/gogame/rpc/service"
	"google.golang.org/grpc"
	"log"
)

func main() {
	if er := global.LoadGlobalConfig(); er != nil {
		log.Panic(er)
	}

	rpcLoggerSvr := &rpcservice.RpcServer{}
	if er := rpcLoggerSvr.StartService(global.Cfg.Port, func(s *grpc.Server) {
		loggersvr.RegisterLoggerServer(s, &rpcservice.RpcLoggerServer{})
	}); er != nil {
		log.Panic(er)
	}

	select {}
}
