package service

import (
	"demo/gogame/proto/db"
	log "github.com/alecthomas/log4go"
	"io"
)

type RpcDbServer struct {
}

func (c *RpcDbServer) BidStream(stream dbserv.DB_BidStreamServer) error {
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
				log.Error("服务端数据接收出错", er)
				return er
			}

			log.Info(req)

			switch req.MainId {
			case 0:
				er = stream.SendMsg(&dbserv.Response{MainId: req.MainId, SubId: req.SubId, RequestId: req.RequestId, Output: req.Input})
				if er != nil {
					log.Error(er)
					return er
				}
				return nil
			case 1:
				er = stream.SendMsg(&dbserv.Response{MainId: req.MainId, SubId: req.SubId, RequestId: req.RequestId, Output: req.Input})
				if er != nil {
					log.Error(er)
					return er
				}
			}
		}
	}

	return nil
}
