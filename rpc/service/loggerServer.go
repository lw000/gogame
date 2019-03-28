package rpcservice

import (
	"demo/gogame/proto/logger"
	log "github.com/alecthomas/log4go"
	"golang.org/x/net/context"
	"io"
)

type RpcLoggerServer struct {
	f func(stream loggersvr.Logger_BidStreamServer, req *loggersvr.Request)
}

func (r *RpcLoggerServer) HandleMessage(f func(stream loggersvr.Logger_BidStreamServer, req *loggersvr.Request)) {
	r.f = f
}

func (r *RpcLoggerServer) RegisterService(context context.Context, req *loggersvr.RequestRegisterService) (*loggersvr.ResponseRegisterService, error) {
	log.Info(req)
	return &loggersvr.ResponseRegisterService{Status: 0, Msg: "success"}, nil
}

func (r *RpcLoggerServer) WriteLogger(context context.Context, req *loggersvr.Request) (*loggersvr.Response, error) {
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
	return &loggersvr.Response{Status: status}, nil
}

func (r *RpcLoggerServer) BidStream(stream loggersvr.Logger_BidStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Error("收到客户端通过context发出的终止信号")
			return ctx.Err()
		default:
			req, er := stream.Recv()
			if er == io.EOF {
				log.Error("客户端发送数据流结束")
				return nil
			}

			if er != nil {
				log.Error("服务端数据接收出错 %v", er)
				return er
			}

			log.Info(req)

			if r.f != nil {
				r.f(stream, req)
			} else {
				log.Warn("onMessage is empty")
			}
		}
	}

	return nil
}
