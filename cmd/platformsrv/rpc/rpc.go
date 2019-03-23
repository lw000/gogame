package rpc

import (
	"demo/gogame/cmd/platformsrv/rpc/client"
	"demo/gogame/cmd/platformsrv/rpc/service"
	"log"
)

type PlatformRpc struct {
	rpcPlatformMgr *client.RpcPlatformManager
	rpcLoggerMgr   *client.RpcLoggerManager
}

func (r *PlatformRpc) Start() error {
	r.rpcPlatformMgr = &client.RpcPlatformManager{}
	r.rpcLoggerMgr = &client.RpcLoggerManager{}

	return nil
}

func (r *PlatformRpc) Stop() error {

	return nil
}

func (r *PlatformRpc) StartRpcService(port int64) {
	go service.StartRpcService(port)
}

func (r *PlatformRpc) StartRpcPlatformClient(address string) {
	er := r.rpcPlatformMgr.Start(address)
	if er != nil {
		log.Panic(er)
	}
}

func (r *PlatformRpc) StartRpcLoggerClient(address string) {
	er := r.rpcLoggerMgr.Start(address)
	if er != nil {
		log.Panic(er)
	}
}
