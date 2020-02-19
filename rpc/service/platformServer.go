package rpcservice

import (
	log "github.com/alecthomas/log4go"
	"gogame/protos/platform"
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
			req, err := stream.Recv()
			if err == io.EOF {
				log.Error("客户端发送数据流结束")
				return nil
			}

			if err != nil {
				log.Error("服务端数据接收出错 %v", err)
				return err
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
