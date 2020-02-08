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
	gServe *grpc.Server
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

func (serve *RpcServer) StartService(port int64, fn func(s *grpc.Server)) error {
	serve.port = port
	var err error
	address := fmt.Sprintf(":%d", serve.port)
	serve.listen, err = net.Listen("tcp", address)
	if err != nil {
		log.Error(err)
		return err
	}

	// opts := []grpc.ServerOption{grpc.UnaryInterceptor(unaryInterceptor)}
	// serve.serv = grpc.NewServer(opts...)
	serve.gServe = grpc.NewServer()

	fn(serve.gServe)

	reflection.Register(serve.gServe)

	log.Info("Listening and serving RPC on [:%d]", serve.port)

	go serve.run()

	return nil
}

func (serve *RpcServer) run() {
	if err := serve.gServe.Serve(serve.listen); err != nil {
		log.Error("failed to gServe: %v", err)
	}
}
