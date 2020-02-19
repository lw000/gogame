package rpcservice

import (
	log "github.com/alecthomas/log4go"
	"gogame/protos/logger"
	"golang.org/x/net/context"
)

type RpcLoggerServer struct {
}

func (r *RpcLoggerServer) WriteLogger(context context.Context, req *loggersvr.Request) (*loggersvr.Response, error) {
	log.Info(req)
	return &loggersvr.Response{Status: 1}, nil
}
