package rpcclient

import (
	log "github.com/alecthomas/log4go"
	"gogame/protos/platform"
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
	var err error
	r.conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Error("did not connect: %v", err)
		return err
	}
	r.client = platformsvr.NewPlatformClient(r.conn)

	return nil
}

func (r *RpcPlatformClient) Stop() error {
	err := r.conn.Close()
	if err != nil {

	}
	return err
}

func (r *RpcPlatformClient) CreateStream(onMessage func(response *platformsvr.Response)) (*RpcPlatformStream, error) {
	var err error
	var rpcStream RpcPlatformStream
	rpcStream.stream, err = r.client.BidStream(context.Background())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	rpcStream.onMessage = onMessage

	go rpcStream.run()

	return &rpcStream, nil
}

func (r *RpcPlatformStream) SendMessage(mainId int32, subId int32, requestId int32, input string) error {
	if err := r.stream.SendMsg(&platformsvr.Request{MainId: mainId, SubId: subId, RequestId: requestId, Input: input}); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (r *RpcPlatformStream) CloseSend() error {
	return r.stream.CloseSend()
}

func (r *RpcPlatformStream) run() {
	var err error
	var resp *platformsvr.Response
	for {
		resp, err = r.stream.Recv()
		if err == io.EOF {
			log.Error("接收到服务端的结束信号 %v", err)
			break
		}

		if err != nil {
			log.Error("接收数据错误 %v", err)
			break
		}

		if r.onMessage != nil {
			r.onMessage(resp)
		}
	}
}
