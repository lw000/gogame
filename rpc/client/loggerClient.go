package rpcclient

import (
	"demo/gogame/proto/logger"
	log "github.com/alecthomas/log4go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type RpcLoggerClient struct {
	conn *grpc.ClientConn
	c    loggersvr.LoggerClient
}

func (r *RpcLoggerClient) Start(address string) error {
	var er error
	r.conn, er = grpc.Dial(address, grpc.WithInsecure())
	if er != nil {
		log.Error("did not connect:%v", er)
		return er
	}

	r.c = loggersvr.NewLoggerClient(r.conn)

	return nil
}

func (r *RpcLoggerClient) Stop() {

}

func (r *RpcLoggerClient) WriteLogger(msg string) error {
	ctx := context.Background()
	reply, er := r.c.WriteLogger(ctx, &loggersvr.Request{ServerId: 10000, ServerTag: "platformsrv", Msg: msg})
	if er != nil {
		log.Error("did not connect:%v", er)
		return er
	}
	if reply.Status != 1 {
		log.Error(reply)
	}
	return nil
}
