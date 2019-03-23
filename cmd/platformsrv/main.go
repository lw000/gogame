package main

import (
	"demo/gogame/cmd/platformsrv/global"
	"demo/gogame/cmd/platformsrv/htp"
	"demo/gogame/cmd/platformsrv/platform"
	"demo/gogame/cmd/platformsrv/rpc"
	"demo/gogame/common/sys"
	"fmt"
	"log"
)

var (
	plat *platform.Platform
)

func main() {
	ggsys.RegisterOnInterrupt(func() {
		if er := plat.Stop(); er != nil {

		}
	})

	plat = platform.NewPlatform(1, "棋牌游戏")
	if er := plat.Start(); er != nil {
		log.Panic(er)
	}

	if er := global.LoadGlobalConfig(); er != nil {
		log.Panic(er)
	}

	pfh := htp.PlatformHtp{}
	if er := pfh.Start(global.Cfg.HTTPPort); er != nil {
		log.Panic(er)
	}

	pfc := rpc.PlatformRpc{}
	if er := pfc.Start(); er != nil {
		log.Panic(er)
	}
	pfc.StartRpcService(global.Cfg.RPCPort)
	pfc.StartRpcPlatformClient(fmt.Sprintf("%s:%d", global.Cfg.GateWay.Host, global.Cfg.GateWay.Port))
	pfc.StartRpcLoggerClient(fmt.Sprintf("%s:%d", global.Cfg.LoggerServ.Host, global.Cfg.LoggerServ.Port))

	select {}
}
