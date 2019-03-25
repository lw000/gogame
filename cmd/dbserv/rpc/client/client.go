package client

import (
	"google.golang.org/grpc"
	"log"
)

type RpcLoggerManager struct {
	conn *grpc.ClientConn
	cli  RpcLoggerClient
}

func (r *RpcLoggerManager) Start(address string) error {
	var er error
	r.conn, er = grpc.Dial(address, grpc.WithInsecure())
	if er != nil {
		log.Fatalf("did not connect:%v", er)
		return er
	}

	r.cli = NewRpcLoggerClient()
	er = r.cli.Start(r.conn)
	if er != nil {
		log.Panic(er)
	}
	return nil
}

func (r *RpcLoggerManager) Stop() error {
	er := r.conn.Close()
	return er
}

func (r *RpcLoggerManager) WriteLogger(msg string) error {
	return r.cli.WriteLogger(msg)
}
