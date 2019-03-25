package client

import (
	"demo/gogame/proto/logger"
	log "github.com/alecthomas/log4go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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
		log.Error("did not connect:%v", er)
		return er
	}

	if reply.Status == 1 {

	}

	return nil
}
