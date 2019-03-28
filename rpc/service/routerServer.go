package rpcservice

import (
	"context"
	"demo/gogame/proto/router"
	log "github.com/alecthomas/log4go"
	"io"
)

type RpcRouterServer struct {
	ServiceId         int32
	onMessage         func(stream routersvr.Router_ForwardingDataStreamServer, req *routersvr.ForwardMessage)
	onConnected       func(stream routersvr.Router_ForwardingDataStreamServer)
	onDisConnected    func(stream routersvr.Router_ForwardingDataStreamServer)
	onRegisterService func(protocols []*routersvr.RouterProtocol) error
}

func (r *RpcRouterServer) HandleConnect(f func(stream routersvr.Router_ForwardingDataStreamServer)) {
	r.onConnected = f
}

func (r *RpcRouterServer) HandleDisConnected(f func(stream routersvr.Router_ForwardingDataStreamServer)) {
	r.onDisConnected = f
}

func (r *RpcRouterServer) HandleMessage(f func(stream routersvr.Router_ForwardingDataStreamServer, req *routersvr.ForwardMessage)) {
	r.onMessage = f
}

func (r *RpcRouterServer) HandleRegisterService(f func(protocols []*routersvr.RouterProtocol) error) {
	r.onRegisterService = f
}

func (r *RpcRouterServer) RegisterService(context context.Context, req *routersvr.RequestRegisterService) (*routersvr.ResponseRegisterService, error) {
	er := r.onRegisterService(req.Protocols)
	if er != nil {
		return &routersvr.ResponseRegisterService{Status: 0, Msg: er.Error()}, nil
	}
	return &routersvr.ResponseRegisterService{Status: 1, Msg: "ok"}, nil
}

func (r *RpcRouterServer) ForwardingData(context context.Context, req *routersvr.ForwardMessage) (*routersvr.ForwardMessage, error) {
	log.Info(req)
	return &routersvr.ForwardMessage{ServiceId: req.ServiceId, Uuid: req.Uuid, Msg: req.Msg}, nil
}

func (r *RpcRouterServer) ForwardingDataStream(stream routersvr.Router_ForwardingDataStreamServer) error {
	log.Info(stream)

	if r.onConnected != nil {
		r.onConnected(stream)
	}

	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Error("收到客户端通过context发出的终止信号")
			if r.onDisConnected != nil {
				r.onDisConnected(stream)
			}
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

			if r.onMessage != nil {
				r.onMessage(stream, req)
			} else {
				log.Warn("onMessage is empty")
			}
		}
	}

	return nil
}
