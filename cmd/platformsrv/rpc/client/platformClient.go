package client

import (
	"demo/gogame/proto/platform"
	log "github.com/alecthomas/log4go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
)

type RpcPlatformClient struct {
	c      platform.PlatformClient
	stream platform.Platform_BidStreamClient
}

func (r *RpcPlatformClient) Stream() platform.Platform_BidStreamClient {
	return r.stream
}

func (r *RpcPlatformClient) Start(conn *grpc.ClientConn) error {
	r.c = platform.NewPlatformClient(conn)
	r.StartStream()
	return nil
}

func (r *RpcPlatformClient) Stop() {

}

func (r *RpcPlatformClient) RegisterService() {
	ctx := context.Background()
	reply, er := r.c.RegisterService(ctx, &platform.RequestRegisterService{ServiceId: 1000, ServiceName: "platform", ServiceVersion: "1.0.1"})
	if er != nil {
		log.Error("did not connect:%v", er)
		return
	}

	if reply.Status != 1 {
		log.Error(reply)
	}
}

func (r *RpcPlatformClient) StartStream() {
	var er error
	r.stream, er = r.c.BidStream(context.Background())
	if er != nil {
		log.Error(er)
		return
	}

	go r.run()
}

func (r *RpcPlatformClient) SendStreamMessage(mainId int32, subId int32, requestId int32, input string) error {
	if er := r.stream.SendMsg(&platform.Request{MainId: mainId, SubId: subId, RequestId: requestId, Input: input}); er != nil {
		log.Error(er)
		return er
	}
	return nil
}

func (r *RpcPlatformClient) run() {
	var er error
	var resp *platform.Response
	for {
		resp, er = r.stream.Recv()
		if er == io.EOF {
			log.Error("接收到服务端的结束信号", er)
			break
		}

		if er != nil {
			log.Error("接收数据错误", er)
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
