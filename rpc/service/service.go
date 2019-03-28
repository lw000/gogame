package rpcservice

import (
	"fmt"
	log "github.com/alecthomas/log4go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type RpcServer struct {
	port   int64
	listen net.Listener
	serv   *grpc.Server
}

func (r *RpcServer) StartService(port int64, f func(s *grpc.Server)) error {
	r.port = port

	address := fmt.Sprintf(":%d", r.port)

	var er error
	r.listen, er = net.Listen("tcp", address)
	if er != nil {
		log.Error(er)
		return er
	}

	r.serv = grpc.NewServer()

	f(r.serv)

	reflection.Register(r.serv)

	log.Info("Listening and serving RPC on [:%d]", r.port)

	go r.run()

	return nil
}

func (r *RpcServer) run() {
	if er := r.serv.Serve(r.listen); er != nil {
		log.Error("failed to serve: %v", er)
	}
}
