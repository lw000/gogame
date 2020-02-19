package main

import (
	"gogame/cmd/loggersrv/global"
	"gogame/protos/logger"
	"gogame/rpc/service"
	"google.golang.org/grpc"
	"log"
)

func main() {
	if err := global.LoadGlobalConfig(); err != nil {
		log.Panic(err)
	}

	rpcLoggerSvr := &rpcservice.RpcServer{}
	if err := rpcLoggerSvr.StartService(global.Cfg.Port, func(s *grpc.Server) {
		loggersvr.RegisterLoggerServer(s, &rpcservice.RpcLoggerServer{})
	}); err != nil {
		log.Panic(err)
	}

	select {}
}
