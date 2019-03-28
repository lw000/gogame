package rpcservice

import (
	"context"
	"demo/gogame/proto/db"
	log "github.com/alecthomas/log4go"
	"io"
)

type RpcDbServer struct {
	f func(stream dbsvr.DB_BidStreamServer, req *dbsvr.Request)
}

func (r *RpcDbServer) HandleMessage(f func(stream dbsvr.DB_BidStreamServer, req *dbsvr.Request)) {
	r.f = f
}

func (r *RpcDbServer) RegisterService(context context.Context, req *dbsvr.RequestRegisterService) (*dbsvr.ResponseRegisterService, error) {
	log.Info(req)
	return &dbsvr.ResponseRegisterService{Status: 0, Msg: "success"}, nil
}

func (r *RpcDbServer) BidStream(stream dbsvr.DB_BidStreamServer) error {
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
