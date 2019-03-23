package service

import (
	"demo/gogame/proto/chat"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartRpcService(port int) {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Panic(err)
	}

	serv := grpc.NewServer()

	chat.RegisterChatServer(serv, &RpcChatStreamer{})

	reflection.Register(serv)

	if err = serv.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
