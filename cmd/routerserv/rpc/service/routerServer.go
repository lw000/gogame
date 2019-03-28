package service

import (
	"context"
	"demo/gogame/proto/router"
	log "github.com/alecthomas/log4go"
	"io"
)

type RpcRouterServer struct {
	f func(stream routersvr.Router_BidStreamServer, req *routersvr.ForwardRequest)
}

func (r *RpcRouterServer) HandleMessage(f func(stream routersvr.Router_BidStreamServer, req *routersvr.ForwardRequest)) {
	r.f = f
}

func (r *RpcRouterServer) RegisterService(context context.Context, req *routersvr.RequestRegisterService) (*routersvr.ResponseRegisterService, error) {
	log.Info(req)
	return &routersvr.ResponseRegisterService{Status: 0, Msg: "success"}, nil
}

func (r *RpcRouterServer) ForwardingData(context context.Context, req *routersvr.ForwardRequest) (*routersvr.ForwardResponse, error) {
	log.Info(req)
	return &routersvr.ForwardResponse{ServiceId: req.ServiceId, MainId: req.MainId, SubId: req.SubId, Uuid: req.Uuid, Output: req.Input}, nil
}

func (r *RpcRouterServer) BidStream(stream routersvr.Router_BidStreamServer) error {
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
