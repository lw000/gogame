package main

import (
	"demo/gogame/cmd/platformsrv/global"
	"demo/gogame/cmd/platformsrv/htp"
	"demo/gogame/cmd/platformsrv/platform"
	"demo/gogame/cmd/platformsrv/rpc"
	"demo/gogame/common/sys"
	"fmt"
	"log"
	"sync/atomic"
	"time"
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
	pfc.StartRpcService(global.Cfg.Port)
	//pfc.StartRpcPlatformClient(fmt.Sprintf("%s:%d", global.Cfg.GateWay.Host, global.Cfg.GateWay.Port))
	pfc.StartRpcLoggerClient(fmt.Sprintf("%s:%d", global.Cfg.LoggerServ.Host, global.Cfg.LoggerServ.Port))
	pfc.StartRpcDbClient(fmt.Sprintf("%s:%d", global.Cfg.DBServ.Host, global.Cfg.DBServ.Port))

	//测试日志写入服务
	go func() {
		for {
			er := pfc.RpcLoggerMgr().WriteLogger("dbserv")
			if er != nil {

			}
			time.Sleep(time.Second * time.Duration(1))
		}
	}()

	//go func() {
	//	for {
	//		var requestId int32 = 0
	//		atomic.AddInt32(&requestId, 1)
	//		er := pfc.RpcPlatformMgr().Cli().SendStreamMessage(1, 10000, requestId, "platform-1")
	//		if er != nil {
	//			log.Println(er)
	//			return
	//		}
	//		time.Sleep(time.Second * time.Duration(1))
	//	}
	//}()

	//测试数据库服务
	go func() {
		for {
			var requestId int32 = 0
			atomic.AddInt32(&requestId, 1)
			er := pfc.RpcDbMgr().Cli().SendStreamMessage(1, 10000, requestId, "dbserv-1")
			if er != nil {
				log.Println(er)
			}
			time.Sleep(time.Second * time.Duration(1))
		}
	}()

	select {}
}
