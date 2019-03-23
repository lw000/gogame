package service

import (
	"context"
	"demo/gogame/proto/platform"
	"io"
	"log"
)

type RpcPlatformServer struct {
}

func (c *RpcPlatformServer) BidStream(stream platform.Platform_BidStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("收到客户端通过context发出的终止信号")
			return ctx.Err()
		default:
			req, er := stream.Recv()
			if er == io.EOF {
				log.Println("客户端发送数据流结束")
				return nil
			}

			if er != nil {
				log.Println("服务端数据接收出错", er)
				return er
			}

			log.Println(req)

			switch req.MainId {
			case 0:
				er = stream.SendMsg(&platform.Response{MainId: req.MainId, SubId: req.SubId, RequestId: req.RequestId, Output: req.Input})
				if er != nil {
					return er
				}
				return nil
			case 1:
				er = stream.SendMsg(&platform.Response{MainId: req.MainId, SubId: req.SubId, RequestId: req.RequestId, Output: req.Input})
				if er != nil {
					return er
				}
			}
		}
	}

	return nil
}

func (c *RpcPlatformServer) RegisterService(context context.Context, req *platform.RequestRegisterService) (*platform.ResponseRegisterService, error) {
	log.Println(req)
	return &platform.ResponseRegisterService{Status: 0, Msg: "success"}, nil
}
