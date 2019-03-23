package rpc

import "demo/gogame/cmd/loggersrv/rpc/service"

type LoggerRpc struct {
}

func (r *LoggerRpc) Start(port int64) error {
	go service.StartRpcService(port)
	return nil
}

func (r *LoggerRpc) Stop() error {

	return nil
}
