package client

import (
	"google.golang.org/grpc"
	"log"
)

type RpcDbManager struct {
	conn *grpc.ClientConn
	cli  *RpcDbClient
}

type RpcLoggerManager struct {
	conn *grpc.ClientConn
	cli  *RpcLoggerClient
}

type RpcPlatformManager struct {
	conn *grpc.ClientConn
	cli  *RpcPlatformClient
}

func (r *RpcPlatformManager) Cli() *RpcPlatformClient {
	return r.cli
}

func (r *RpcPlatformManager) Start(address string) error {
	var er error
	r.conn, er = grpc.Dial(address, grpc.WithInsecure())
	if er != nil {
		log.Fatalf("did not connect:%v", er)
		return er
	}

	r.cli = &RpcPlatformClient{}
	er = r.cli.Start(r.conn)
	if er != nil {
		log.Panic(er)
	}

	return nil
}

func (r *RpcPlatformManager) Stop() error {
	er := r.conn.Close()
	return er
}

func (r *RpcLoggerManager) Cli() *RpcLoggerClient {
	return r.cli
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

func (r *RpcDbManager) Cli() *RpcDbClient {
	return r.cli
}

func (r *RpcDbManager) Start(address string) error {
	var er error
	r.conn, er = grpc.Dial(address, grpc.WithInsecure())
	if er != nil {
		log.Fatalf("did not connect:%v", er)
		return er
	}

	r.cli = &RpcDbClient{}
	er = r.cli.Start(r.conn)
	if er != nil {
		log.Panic(er)
	}

	return nil
}

func (r *RpcDbManager) Stop() error {
	er := r.conn.Close()
	return er
}
