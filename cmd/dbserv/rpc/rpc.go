package rpc

import (
	"demo/gogame/cmd/dbserv/rpc/client"
	"log"
)

type DbRpc struct {
	rpcLoggerMgr *client.RpcLoggerManager
}

func (r *DbRpc) Start() error {
	r.rpcLoggerMgr = &client.RpcLoggerManager{}

	return nil
}

func (r *DbRpc) Stop() error {

	return nil
}

func (r *DbRpc) StartRpcLoggerClient(address string) {
	er := r.rpcLoggerMgr.Start(address)
	if er != nil {
		log.Panic(er)
	}
}
