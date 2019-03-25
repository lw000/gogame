package service

import (
	"demo/gogame/proto/logger"
	log "github.com/alecthomas/log4go"
	"golang.org/x/net/context"
)

type RocLoggerServer struct {
}

func (l *RocLoggerServer) WriteLogger(context context.Context, req *Loggersvr.Request) (*Loggersvr.Response, error) {
	var status int32 = 0
	switch req.ServerId {
	case 10000:
		log.Info(req)
		status = 1
	case 10001:
		log.Info(req)
		status = 1
	case 10002:
		log.Info(req)
		status = 1
	case 10003:
		log.Info(req)
		status = 1
	}
	return &Loggersvr.Response{Status: status}, nil
}
