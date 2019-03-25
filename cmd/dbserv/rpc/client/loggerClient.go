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

func (r *RpcLoggerClient) SendMessage(msg string) error {
	ctx := context.Background()
	reply, er := r.c.WriteLogger(ctx, &Loggersvr.Request{ServerId: 10002, ServerTag: "dbserv", Msg: msg})
	if er != nil {
		log.Printf("did not connect:%v", er)
		return er
	}
	log.Println(reply)

	return nil
}

func (r *RpcLoggerClient) Test() {
	for {
		er := r.SendMessage("dbserv")
		if er != nil {

		}
		time.Sleep(time.Second * time.Duration(1))
	}
}
