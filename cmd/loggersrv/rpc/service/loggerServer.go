package service

import (
	"demo/gogame/proto/logger"
	"golang.org/x/net/context"
	"log"
)

type LogServer struct {
}

func (l *LogServer) WriteLogger(context context.Context, req *Loggersvr.Request) (*Loggersvr.Response, error) {
	var status int32 = 0
	switch req.ServerId {
	case 10000:
		log.Println(req)
		status = 1
	case 10001:
		log.Println(req)
		status = 1
	case 10002:
		log.Println(req)
		status = 1
	case 10003:
		log.Println(req)
		status = 1
	}
	log.Println(req)
	return &Loggersvr.Response{Status: status}, nil
}
