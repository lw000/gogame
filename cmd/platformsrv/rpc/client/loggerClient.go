package client

import (
	"demo/gogame/proto/logger"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

type RpcLoggerClient struct {
	c Loggersvr.LoggerClient
}

func (r *RpcLoggerClient) Start(conn *grpc.ClientConn) error {
	r.c = Loggersvr.NewLoggerClient(conn)

	return nil
}

func (r *RpcLoggerClient) Stop() {

}

func (r *RpcLoggerClient) SendMessage(msg string) {
	ctx := context.Background()
	reply, er := r.c.WriteLogger(ctx, &Loggersvr.Request{ServerId: 1000, Msg: msg})
	if er != nil {
		log.Printf("did not connect:%v", er)
		return
	}
	log.Println(reply)
}

func (r *RpcLoggerClient) Test() {
	for {
		r.SendMessage("platform")
		time.Sleep(time.Millisecond * time.Duration(100))
	}
}
