package rpcservice

import (
	"context"
	"fmt"
	log "github.com/alecthomas/log4go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"net"
	"strings"
)

type RpcServer struct {
	port   int64
	listen net.Listener
	serv   *grpc.Server
}

func authenticateClient(ctx context.Context) (string, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		clientUsername := strings.Join(md["username"], "")
		clientPassword := strings.Join(md["password"], "")
		if clientUsername != "valineliu" {
			return "", fmt.Errorf("unknown user %s", clientUsername)
		}
		if clientPassword != "root" {
			return "", fmt.Errorf("wrong password %s", clientPassword)
		}
		log.Info("authenticated client: %s", clientUsername)
		return "9527", nil
	}
	return "", fmt.Errorf("missing credentials")
}

// unaryInterceptor calls authenticateClient with current context
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	clientID, err := authenticateClient(ctx)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, "clientID", clientID)
	return handler(ctx, req)
}

func (r *RpcServer) StartService(port int64, f func(s *grpc.Server)) error {
	r.port = port
	var er error
	address := fmt.Sprintf(":%d", r.port)
	r.listen, er = net.Listen("tcp", address)
	if er != nil {
		log.Error(er)
		return er
	}

	//opts := []grpc.ServerOption{grpc.UnaryInterceptor(unaryInterceptor)}
	//r.serv = grpc.NewServer(opts...)
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
