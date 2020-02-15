package rpcservice

import (
	"demo/gogame/protos/platform"
	log "github.com/alecthomas/log4go"
	"io"
)

type RpcPlatformServer struct {
	f func(stream platformsvr.Platform_BidStreamServer, req *platformsvr.Request)
}

func (r *RpcPlatformServer) HandleMessage(f func(stream platformsvr.Platform_BidStreamServer, req *platformsvr.Request)) {
	r.f = f
}

func (r *RpcPlatformServer) BidStream(stream platformsvr.Platform_BidStreamServer) error {
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
