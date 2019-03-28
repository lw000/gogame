package rpcclient

import (
	"demo/gogame/proto/center"
	log "github.com/alecthomas/log4go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
)

type RpcCenterClient struct {
	conn *grpc.ClientConn
	c    centersvr.CenterClient
}

type RpcCenterStream struct {
	onMessage func(response *centersvr.Response)
	stream    centersvr.Center_BidStreamClient
}

func (r *RpcCenterClient) Start(address string) error {
	var er error
	r.conn, er = grpc.Dial(address, grpc.WithInsecure())
	if er != nil {
		log.Error("did not connect: %v", er)
		return er
	}
	r.c = centersvr.NewCenterClient(r.conn)

	return nil
}

func (r *RpcCenterClient) Stop() {

}

func (r *RpcCenterClient) CreateStream(onMessage func(response *centersvr.Response)) (*RpcCenterStream, error) {
	var er error
	var rpcStream RpcCenterStream
	rpcStream.stream, er = r.c.BidStream(context.Background())
	if er != nil {
		log.Error(er)
		return nil, er
	}

	rpcStream.onMessage = onMessage

	go rpcStream.run()

	return &rpcStream, nil
}

func (r *RpcCenterStream) SendMessage(mainId int32, subId int32, requestId int32, input string) error {
	if er := r.stream.SendMsg(&centersvr.Request{MainId: mainId, SubId: subId, RequestId: requestId, Input: input}); er != nil {
		log.Error(er)
		return er
	}
	return nil
}

func (r *RpcCenterStream) CloseSend() error {
	return r.stream.CloseSend()
}

func (r *RpcCenterStream) run() {
	var er error
	var resp *centersvr.Response
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
