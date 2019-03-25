package client

import (
	"demo/gogame/proto/db"
	"demo/gogame/proto/platform"
	log "github.com/alecthomas/log4go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
)

type RpcDbClient struct {
	c      dbserv.DBClient
	stream dbserv.DB_BidStreamClient
}

func (r *RpcDbClient) Start(conn *grpc.ClientConn) error {
	r.c = dbserv.NewDBClient(conn)
	r.StartStream()
	return nil
}

func (r *RpcDbClient) Stop() {

}

func (r *RpcDbClient) StartStream() {
	var er error
	r.stream, er = r.c.BidStream(context.Background())
	if er != nil {
		log.Error(er)
		return
	}

	go r.run()
}

func (r *RpcDbClient) SendStreamMessage(mainId int32, subId int32, requestId int32, input string) error {
	if er := r.stream.SendMsg(&platform.Request{MainId: mainId, SubId: subId, RequestId: requestId, Input: input}); er != nil {
		log.Error(er)
		return er
	}
	return nil
}

func (r *RpcDbClient) run() {
	var er error
	var resp *dbserv.Response
	for {
		resp, er = r.stream.Recv()
		if er == io.EOF {
			log.Error("接收到服务端的结束信号 %v", er)
			break
		}

		if er != nil {
			log.Error("接收数据错误 %v", er)
			break
		}
		switch resp.MainId {
		case 1:
			log.Info(resp)
		case 2:
			log.Info(resp)
		}
	}
}
