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
	reply, er := r.c.WriteLogger(ctx, &Loggersvr.Request{ServerId: 10001, ServerTag: "gamesrv", Msg: msg})
	if er != nil {
		log.Printf("did not connect:%v", er)
		return
	}
	log.Println(reply)
}

func (r *RpcLoggerClient) Test() {
	for {
		r.SendMessage("gamesrv")
		time.Sleep(time.Second * time.Duration(1))
	}
}
