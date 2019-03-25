package client

import (
	"demo/gogame/proto/logger"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
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

func (r *RpcLoggerClient) WriteLogger(msg string) error {
	ctx := context.Background()
	reply, er := r.c.WriteLogger(ctx, &Loggersvr.Request{ServerId: 10000, ServerTag: "platformsrv", Msg: msg})
	if er != nil {
		log.Printf("did not connect:%v", er)
		return er
	}
	log.Println(reply)

	return nil
}
