package rpc

import (
	"demo/gogame/cmd/dbserv/rpc/client"
	"demo/gogame/cmd/dbserv/rpc/service"
	"log"
)

type DbRpc struct {
	rpcLoggerMgr *client.RpcLoggerManager
}

func (r *DbRpc) RpcLoggerMgr() *client.RpcLoggerManager {
	return r.rpcLoggerMgr
}

func (r *DbRpc) Start() error {
	r.rpcLoggerMgr = &client.RpcLoggerManager{}

	return nil
}

func (r *DbRpc) Stop() error {

	return nil
}

func (r *DbRpc) StartRpcDbServer(port int64) {
	go service.StartRpcService(port)
}

func (r *DbRpc) StartRpcLoggerClient(address string) {
	er := r.rpcLoggerMgr.Start(address)
	if er != nil {
		log.Panic(er)
	}
}
