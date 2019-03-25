package rpc

import (
	"demo/gogame/cmd/gamesrv/rpc/client"
	"log"
)

type GameRpc struct {
	rpcLoggerMgr *client.RpcLoggerManager
}

func (r *GameRpc) RpcLoggerMgr() *client.RpcLoggerManager {
	return r.rpcLoggerMgr
}

func (r *GameRpc) Start() error {
	r.rpcLoggerMgr = &client.RpcLoggerManager{}

	return nil
}

func (r *GameRpc) Stop() error {

	return nil
}

func (r *GameRpc) StartRpcLoggerClient(address string) {
	er := r.rpcLoggerMgr.Start(address)
	if er != nil {
		log.Panic(er)
	}
}
