package client

import (
	"context"
	"demo/grpc_test/proto/chat"
	"demo/grpc_test/proto/helloworld"
	"google.golang.org/grpc"
	"io"
	"log"
	"sync/atomic"
	"time"
)

type RpcClient struct {
}

var (
	count int64
	conn  *grpc.ClientConn
)

func TestGreeter() {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()
	ctx := context.Background()
	client := helloworld.NewGreeterClient(conn)
	for i := 0; i < 100; i++ {
		begin := time.Now()
		reply, er := client.SayHello(ctx, &helloworld.HelloRequest{Name: "world"})
		if er != nil {
			log.Printf("did not connect:%v", er)
			return
		}
		atomic.AddInt64(&count, 1)
		log.Printf("[%d] [%v] Greeting: %s", count, time.Now().Sub(begin), reply.Message)
	}
}

func TestMathService() {
	ctx := context.Background()
	client := helloworld.NewMathServiceClient(conn)
	{
		reply, er := client.Add(ctx, &helloworld.AddRequest{A: 10, B: 30})
		if er != nil {
			log.Printf("did not connect:%v", er)
			return
		}
		atomic.AddInt64(&count, 1)
		log.Printf("Add: [%d] [%v]", count, reply.C)
	}

	{
		reply, er := client.Sub(ctx, &helloworld.SubRequest{A: 10, B: 100})
		if er != nil {
			log.Printf("did not connect:%v", er)
			return
		}
		atomic.AddInt64(&count, 1)
		log.Printf("Sub: [%d] [%v]", count, reply.C)
	}
}

func TestChat() {
	client := chat.NewChatClient(conn)
	stream, er := client.BidStream(context.Background())
	if er != nil {
		log.Println(er)
		return
	}

	var requestId int32 = 0
	go func() {
		for {
			atomic.AddInt32(&requestId, 1)
			if er = stream.SendMsg(&chat.Request{MainId: 1, SubId: 10000, RequestId: requestId, Input: "message-1"}); er != nil {
				log.Println(er)
				return
			}
			//time.Sleep(time.Second * time.Duration(1))
		}
	}()

	go func() {
		for {
			atomic.AddInt32(&requestId, 1)
			if er = stream.SendMsg(&chat.Request{MainId: 2, SubId: 10000, RequestId: requestId, Input: "message-2"}); er != nil {
				log.Println(er)
				return
			}
			//time.Sleep(time.Second * time.Duration(1))
		}
	}()

	go func() {
		for {
			atomic.AddInt32(&requestId, 1)
			if er = stream.SendMsg(&chat.Request{MainId: 3, SubId: 10000, RequestId: requestId, Input: "message-3"}); er != nil {
				log.Println(er)
				return
			}
			//time.Sleep(time.Second * time.Duration(1))
		}
	}()

	for {
		var resp *chat.Response
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
		case 3:
			log.Println("resp", resp)
		}
	}
}

func StartRpcClient(address string) {
	var er error
	conn, er = grpc.Dial(address, grpc.WithInsecure())
	if er != nil {
		log.Fatalf("did not connect:%v", er)
	}

	defer func() {
		er = conn.Close()
	}()

	go TestGreeter()

	go TestMathService()

	TestChat()

	select {}
}
