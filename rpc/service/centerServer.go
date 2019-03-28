package rpcservice

import (
	"context"
	"demo/gogame/proto/center"
	log "github.com/alecthomas/log4go"
	"io"
)

type RpcCenterServer struct {
	f func(stream centersvr.Center_BidStreamServer, req *centersvr.Request)
}

func (r *RpcCenterServer) HandleMessage(f func(stream centersvr.Center_BidStreamServer, req *centersvr.Request)) {
	r.f = f
}

func (r *RpcCenterServer) RegisterService(context context.Context, req *centersvr.RequestRegisterService) (*centersvr.ResponseRegisterService, error) {
	log.Info(req)
	return &centersvr.ResponseRegisterService{Status: 0, Msg: "success"}, nil
}

func (r *RpcCenterServer) BidStream(stream centersvr.Center_BidStreamServer) error {
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

			if r.f != nil {
				r.f(stream, req)
			} else {
				log.Warn("onMessage is empty")
			}
		}
	}

	return nil
}
