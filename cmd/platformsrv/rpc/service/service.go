package service

import (
	"demo/gogame/proto/platform"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartRpcService(port int64) {
	address := fmt.Sprintf(":%d", port)

	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Panic(err)
	}

	serv := grpc.NewServer()

	platform.RegisterPlatformServer(serv, &RpcPlatformServer{})

	reflection.Register(serv)

	log.Printf("Listening and serving RPC on %s", address)

	if err = serv.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
