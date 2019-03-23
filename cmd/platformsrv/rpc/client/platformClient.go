package client

import (
	"demo/gogame/proto/platform"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"sync/atomic"
	"time"
)

type RpcPlatformClient struct {
	c platform.PlatformClient
}

func (r *RpcPlatformClient) Start(conn *grpc.ClientConn) error {
	r.c = platform.NewPlatformClient(conn)
	return nil
}

func (r *RpcPlatformClient) Stop() {

}

func (r *RpcPlatformClient) Test() {
	ctx := context.Background()

	for {
		reply, er := r.c.RegisterService(ctx, &platform.RequestRegisterService{ServiceId: 1000, ServiceName: "platform", ServiceVersion: "1.0.1"})
		if er != nil {
			log.Printf("did not connect:%v", er)
			return
		}
		log.Printf("[%d] [%s]", reply.Status, reply.Msg)
		time.Sleep(time.Millisecond * time.Duration(100))
	}
}

func (r *RpcPlatformClient) TestStream() {
	stream, er := r.c.BidStream(context.Background())
	if er != nil {
		log.Println(er)
		return
	}

	var requestId int32 = 0

	go func() {
		for {
			atomic.AddInt32(&requestId, 1)
			if er = stream.SendMsg(&platform.Request{MainId: 1, SubId: 10000, RequestId: requestId, Input: "message-1"}); er != nil {
				log.Println(er)
				return
			}
			time.Sleep(time.Millisecond * time.Duration(100))
		}
	}()

	for {
		var resp *platform.Response
		resp, er = stream.Recv()
		if er == io.EOF {
			log.Println("接收到服务端的结算信号", er)
			break
		}

		if er != nil {
			log.Println("接收数据错误", er)
			break
		}
		switch resp.MainId {
		case 1:
			log.Println("resp", resp)
		case 2:
			log.Println("resp", resp)
		}
	}
}
