package rpcclient

import (
	"demo/gogame/protos/platform"
	log "github.com/alecthomas/log4go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
)

type RpcPlatformClient struct {
	conn   *grpc.ClientConn
	client platformsvr.PlatformClient
}

type RpcPlatformStream struct {
	onMessage func(*platformsvr.Response)
	stream    platformsvr.Platform_BidStreamClient
}

func (r *RpcPlatformClient) Start(address string) error {
	var er error
	r.conn, er = grpc.Dial(address, grpc.WithInsecure())
	if er != nil {
		log.Error("did not connect: %v", er)
		return er
	}
	r.client = platformsvr.NewPlatformClient(r.conn)

	return nil
}

func (r *RpcPlatformClient) Stop() error {
	er := r.conn.Close()
	if er != nil {

	}
	return er
}

func (r *RpcPlatformClient) CreateStream(onMessage func(response *platformsvr.Response)) (*RpcPlatformStream, error) {
	var er error
	var rpcStream RpcPlatformStream
	rpcStream.stream, er = r.client.BidStream(context.Background())
	if er != nil {
		log.Error(er)
		return nil, er
	}
	rpcStream.onMessage = onMessage

	go rpcStream.run()

	return &rpcStream, nil
}

func (r *RpcPlatformStream) SendMessage(mainId int32, subId int32, requestId int32, input string) error {
	if er := r.stream.SendMsg(&platformsvr.Request{MainId: mainId, SubId: subId, RequestId: requestId, Input: input}); er != nil {
		log.Error(er)
		return er
	}
	return nil
}

func (r *RpcPlatformStream) CloseSend() error {
	return r.stream.CloseSend()
}

func (r *RpcPlatformStream) run() {
	var er error
	var resp *platformsvr.Response
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
