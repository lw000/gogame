package rpc

import (
	"demo/gogame/cmd/platformsrv/rpc/client"
	"demo/gogame/cmd/platformsrv/rpc/service"
)

type Rpc struct {
}

func StartRpcService(port int) {
	service.StartRpcService(port)
}

func StartRpcClient(address string) {
	client.StartRpcClient(address)
}
