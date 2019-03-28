package rpcclient

import (
	"demo/gogame/proto/db"
	log "github.com/alecthomas/log4go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
)

type RpcDbClient struct {
	conn *grpc.ClientConn
	c    dbsvr.DBClient
}

type RpcDbStream struct {
	onMessage func(response *dbsvr.Response)
	stream    dbsvr.DB_BidStreamClient
}

func (r *RpcDbClient) Start(address string) error {
	var er error
	r.conn, er = grpc.Dial(address, grpc.WithInsecure())
	if er != nil {
		log.Error("did not connect:%v", er)
		return er
	}

	r.c = dbsvr.NewDBClient(r.conn)

	return nil
}

func (r *RpcDbClient) Stop() {

}

func (r *RpcDbClient) CreateStream(onMessage func(response *dbsvr.Response)) (*RpcDbStream, error) {
	var er error
	var rpcStream RpcDbStream
	rpcStream.stream, er = r.c.BidStream(context.Background())
	if er != nil {
		log.Error(er)
		return nil, er
	}

	rpcStream.onMessage = onMessage

	go rpcStream.run()

	return &rpcStream, nil
}

func (r *RpcDbStream) SendMessage(mainId int32, subId int32, requestId int32, input string) error {
	if er := r.stream.SendMsg(&dbsvr.Request{MainId: mainId, SubId: subId, RequestId: requestId, Input: input}); er != nil {
		log.Error(er)
		return er
	}
	return nil
}

func (r *RpcDbStream) CloseSend() error {
	return r.stream.CloseSend()
}

func (r *RpcDbStream) run() {
	var er error
	var resp *dbsvr.Response
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

		if r.onMessage != nil {
			r.onMessage(resp)
		}
	}
}
