package rpc

import (
	"demo/gogame/rpc/client"
	"demo/gogame/rpc/service"
)

type RpcMgr struct {
	rpcPlatformCli *client.RpcPlatformClient
	rpcLoggerCli   *client.RpcLoggerClient
	rpcDbCli       *client.RpcDbClient
	rpcPlatformSvr *service.RpcPlatformServer
}

func (r *RpcMgr) RpcPlatformSvr() *service.RpcPlatformServer {
	return r.rpcPlatformSvr
}

func (r *RpcMgr) SetRpcPlatformSvr(rpcPlatformSvr *service.RpcPlatformServer) {
	r.rpcPlatformSvr = rpcPlatformSvr
}

func (r *RpcMgr) RpcDbCli() *client.RpcDbClient {
	return r.rpcDbCli
}

func (r *RpcMgr) SetRpcDbCli(rpcDbCli *client.RpcDbClient) {
	r.rpcDbCli = rpcDbCli
}

func (r *RpcMgr) SetRpcLoggerCli(rpcLoggerCli *client.RpcLoggerClient) {
	r.rpcLoggerCli = rpcLoggerCli
}

func (r *RpcMgr) SetRpcPlatformCli(rpcPlatformCli *client.RpcPlatformClient) {
	r.rpcPlatformCli = rpcPlatformCli
}

func (r *RpcMgr) RpcPlatformCli() *client.RpcPlatformClient {
	return r.rpcPlatformCli
}

func (r *RpcMgr) RpcLoggerCli() *client.RpcLoggerClient {
	return r.rpcLoggerCli
}
