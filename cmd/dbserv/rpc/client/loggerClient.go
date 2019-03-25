package client

import (
	"demo/gogame/proto/logger"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

type RpcLoggerClient interface {
	Start(conn *grpc.ClientConn) error
	Stop()
	WriteLogger(msg string) error
}

type rpcLoggerClient struct {
	c Loggersvr.LoggerClient
}

func NewRpcLoggerClient() RpcLoggerClient {
	return &rpcLoggerClient{}
}

func (r *rpcLoggerClient) Start(conn *grpc.ClientConn) error {
	r.c = Loggersvr.NewLoggerClient(conn)

	return nil
}

func (r *rpcLoggerClient) Stop() {

}

func (r *rpcLoggerClient) WriteLogger(msg string) error {
	ctx := context.Background()
	reply, er := r.c.WriteLogger(ctx, &Loggersvr.Request{ServerId: 10002, ServerTag: "dbserv", Msg: msg})
	if er != nil {
		log.Printf("did not connect:%v", er)
		return er
	}
	log.Println(reply)

	return nil
}
