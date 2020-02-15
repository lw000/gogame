package rpcclient

import (
	"demo/gogame/protos/logger"
	log "github.com/alecthomas/log4go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type RpcLoggerClient struct {
	conn   *grpc.ClientConn
	client loggersvr.LoggerClient
}

func (r *RpcLoggerClient) Start(address string) error {
	var err error
	r.conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Error("did not connect:%v", err)
		return err
	}

	r.client = loggersvr.NewLoggerClient(r.conn)

	return nil
}

func (r *RpcLoggerClient) Stop() error {
	err := r.conn.Close()
	if err != nil {

	}
	return err
}

func (r *RpcLoggerClient) WriteLogger(msg string) error {
	ctx := context.Background()
	reply, err := r.client.WriteLogger(ctx, &loggersvr.Request{ServerId: 10000, ServerTag: "platformsrv", Msg: msg})
	if err != nil {
		log.Error("did not connect:%v", err)
		return err
	}
	if reply.Status != 1 {
		log.Error(reply)
	}
	return nil
}
