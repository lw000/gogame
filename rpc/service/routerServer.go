package rpcservice

import (
	"demo/gogame/common/utils"
	"demo/gogame/protos/router"
	log "github.com/alecthomas/log4go"
	"io"
)

type RpcRouterServer struct {
	onMessage      func(s *RpcRouterServerStream, req *routersvr.RequestMessage)
	onConnected    func(s *RpcRouterServerStream)
	onDisConnected func(s *RpcRouterServerStream)
}

type RpcRouterServerStream struct {
	ServiceId  int32
	clientUuid string
	stream     routersvr.Router_BindStreamServer
}

func (r *RpcRouterServer) HandleConnect(f func(s *RpcRouterServerStream)) {
	r.onConnected = f
}

func (r *RpcRouterServer) HandleDisConnected(f func(*RpcRouterServerStream)) {
	r.onDisConnected = f
}

func (r *RpcRouterServer) HandleMessage(f func(s *RpcRouterServerStream, req *routersvr.RequestMessage)) {
	r.onMessage = f
}

func (r *RpcRouterServer) BindStream(stream routersvr.Router_BindStreamServer) error {
	serverStream := &RpcRouterServerStream{stream: stream, clientUuid: ggutils.UUID()}
	if r.onConnected != nil {
		r.onConnected(serverStream)
	}
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Error("收到客户端通过context发出的终止信号")
			if r.onDisConnected != nil {
				r.onDisConnected(serverStream)
			}
			return ctx.Err()
		default:
			req, er := stream.Recv()
			if er == io.EOF {
				log.Error("客户端发送数据流结束")
				if r.onDisConnected != nil {
					r.onDisConnected(serverStream)
				}
				return nil
			}

			if er != nil {
				log.Error("服务端数据接收出错 %v", er)
				if r.onDisConnected != nil {
					r.onDisConnected(serverStream)
				}
				return er
			}

			if r.onMessage != nil {
				r.onMessage(serverStream, req)
			} else {
				log.Warn("onMessage is empty")
			}
		}
	}

	return nil
}

func (r RpcRouterServerStream) ClientUuid() string {
	return r.clientUuid
}

func (r RpcRouterServerStream) SendMessage(serviceId int32, msgType int32, uuid string, msg []byte) error {
	er := r.stream.Send(&routersvr.ReponseMessage{ServiceId: serviceId, MsgType: msgType, Cuuid: r.clientUuid, Uuid: uuid, Msg: msg})
	if er != nil {
		log.Error(er)
		return er
	}
	return nil
}
