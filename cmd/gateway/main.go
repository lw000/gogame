package main

import (
	"demo/gogame/cmd/gateway/global"
	"demo/gogame/rpc/client"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"log"
	"net/http"
	"time"
)

type RecvData struct {
	Uid   string `json:"uid"`
	Input string `json:"input"`
}

var (
	rpcLoggerCli *rpcclient.RpcLoggerClient
)

func Test() {
	//测试日志写入服务
	go func() {
		for {
			er := rpcLoggerCli.WriteLogger("gateway-1")
			if er != nil {
				log.Println("error")
			}
			time.Sleep(time.Second * time.Duration(1))
		}
	}()
}

func main() {
	if er := global.LoadGlobalConfig(); er != nil {
		log.Panic(er)
	}

	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "html/index.html")
	})

	r.GET("ws", func(c *gin.Context) {
		er := m.HandleRequest(c.Writer, c.Request)
		if er != nil {

		}
		log.Println(c.Query("uid"))
	})

	m.HandleConnect(func(s *melody.Session) {
		log.Println("新用户连接")

		//centerStream, er := rpcRouterCli.CreateStream(func(resp *routersvr.Response) {
		//	switch resp.MainId {
		//	case 1:
		//		log.Println(resp)
		//	case 2:
		//		log.Println(resp)
		//	}
		//})
		//if er != nil {
		//	log.Println(er)
		//	return
		//}

		rpcRouterCli := &rpcclient.RpcRouterClient{}
		if er := rpcRouterCli.Start(fmt.Sprintf("%s:%d", global.Cfg.RouterServ.Host, global.Cfg.RouterServ.Port)); er != nil {
			log.Panic(er)
		}

		s.Set("client", rpcRouterCli)
	})

	m.HandleMessage(func(s *melody.Session, data []byte) {
		v, ok := s.Get("client")
		if !ok {
			log.Println("error")
			return
		}
		routerCli := v.(*rpcclient.RpcRouterClient)
		if !ok {
			er := s.CloseWithMsg([]byte("被保存的数据流不是[*client.RpcCenterStream]"))
			if er != nil {
				log.Println(er)
			}
			return
		}

		var rdata RecvData
		if er := json.Unmarshal(data, &rdata); er != nil {
			log.Println("解析输入信息失败:", er)
			er = s.CloseWithMsg([]byte("输入信息解析失败"))
			if er != nil {
				log.Println(er)
			}
			return
		}

		resp, er := routerCli.ForwardingData(1, 1, 1, string(data))
		if er != nil {
			log.Println("error")
		}

		er = s.Write(data)
		if er != nil {
			log.Println("error")
			return
		}

		log.Println(resp)

		//if err := centerStream.SendMessage(1, 1, 1, string(data)); err != nil {
		//	er = s.CloseWithMsg([]byte("向gRPC服务端发送消息失败:" + err.Error()))
		//	return
		//}
	})

	m.HandleDisconnect(func(s *melody.Session) {
		log.Println("websocket的断开连接")

		v, ok := s.Get("client")
		if !ok {
			log.Println("error")
			return
		}

		routerCli := v.(*rpcclient.RpcRouterClient)
		if er := routerCli.Stop(); er != nil {
			log.Println("断开stream错误", er)
		}
	})

	rpcLoggerCli = &rpcclient.RpcLoggerClient{}
	if er := rpcLoggerCli.Start(fmt.Sprintf("%s:%d", global.Cfg.LoggerServ.Host, global.Cfg.LoggerServ.Port)); er != nil {
		log.Panic(er)
	}

	Test()

	log.Panic(r.Run(":10010"))
}
