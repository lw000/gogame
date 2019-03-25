package rpc

import (
	"demo/gogame/cmd/ai/rpc/client"
	"log"
)

type AiRpc struct {
	rpcLoggerMgr *client.RpcLoggerManager
}

func (r *AiRpc) Start() error {
	r.rpcLoggerMgr = &client.RpcLoggerManager{}

	return nil
}

func (r *AiRpc) Stop() error {

	return nil
}

func (r *AiRpc) StartRpcLoggerClient(address string) {
	er := r.rpcLoggerMgr.Start(address)
	if er != nil {
		log.Panic(er)
	}
}
