package client

import (
	"google.golang.org/grpc"
	"log"
)

type RpcLoggerManager struct {
	conn *grpc.ClientConn
	cli  *RpcLoggerClient
}

type RpcPlatformManager struct {
	conn *grpc.ClientConn
}

func (r *RpcPlatformManager) Start(address string) error {
	var er error
	r.conn, er = grpc.Dial(address, grpc.WithInsecure())
	if er != nil {
		log.Fatalf("did not connect:%v", er)
		return er
	}

	pclient := RpcPlatformClient{}
	er = pclient.Start(r.conn)
	if er != nil {
		log.Panic(er)
	}

	return nil
}

func (r *RpcPlatformManager) Stop() error {
	er := r.conn.Close()
	return er
}

func (r *RpcLoggerManager) Start(address string) error {
	var er error
	r.conn, er = grpc.Dial(address, grpc.WithInsecure())
	if er != nil {
		log.Fatalf("did not connect:%v", er)
		return er
	}

	r.cli = &RpcLoggerClient{}
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
