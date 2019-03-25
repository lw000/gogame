package main

import (
	"demo/gogame/cmd/gamesrv/global"
	"demo/gogame/cmd/gamesrv/rpc"
	"demo/gogame/common/sys"
	"fmt"
	"github.com/labstack/gommon/log"
)

type Game struct {
	Name string
	Id   int64
}

func (g *Game) Start() error {

	return nil
}

func (g *Game) Stop() error {

	return nil
}

func main() {
	ggsys.RegisterOnInterrupt(func() {

	})

	if er := global.LoadGlobalConfig(); er != nil {
		log.Panic(er)
	}

	gr := rpc.GameRpc{}
	if er := gr.Start(); er != nil {
		log.Panic(er)
	}

	gr.StartRpcLoggerClient(fmt.Sprintf("%s:%d", global.Cfg.LoggerServ.Host, global.Cfg.LoggerServ.Port))

	go func() {
		for {
			er := gr.RpcLoggerMgr().WriteLogger("dbserv")
			if er != nil {

			}
			//time.Sleep(time.Second * time.Duration(1))
		}
	}()

	select {}
}
