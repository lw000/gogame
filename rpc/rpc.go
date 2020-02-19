package rpc

import (
	"gogame/rpc/client"
	"gogame/rpc/service"
)

type RpcMgr struct {
	rpcPlatformCli *rpcclient.RpcPlatformClient
	rpcLoggerCli   *rpcclient.RpcLoggerClient
	rpcDbCli       *rpcclient.RpcDbClient
	rpcPlatformSvr *rpcservice.RpcPlatformServer
}

func (r *RpcMgr) RpcPlatformSvr() *rpcservice.RpcPlatformServer {
	return r.rpcPlatformSvr
}

func (r *RpcMgr) SetRpcPlatformSvr(rpcPlatformSvr *rpcservice.RpcPlatformServer) {
	r.rpcPlatformSvr = rpcPlatformSvr
}

func (r *RpcMgr) RpcDbCli() *rpcclient.RpcDbClient {
	return r.rpcDbCli
}

func (r *RpcMgr) SetRpcDbCli(rpcDbCli *rpcclient.RpcDbClient) {
	r.rpcDbCli = rpcDbCli
}

func (r *RpcMgr) SetRpcLoggerCli(rpcLoggerCli *rpcclient.RpcLoggerClient) {
	r.rpcLoggerCli = rpcLoggerCli
}

func (r *RpcMgr) SetRpcPlatformCli(rpcPlatformCli *rpcclient.RpcPlatformClient) {
	r.rpcPlatformCli = rpcPlatformCli
}

func (r *RpcMgr) RpcPlatformCli() *rpcclient.RpcPlatformClient {
	return r.rpcPlatformCli
}

func (r *RpcMgr) RpcLoggerCli() *rpcclient.RpcLoggerClient {
	return r.rpcLoggerCli
}
