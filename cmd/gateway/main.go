package main

import (
	"demo/gogame/cmd/gateway/global"
	"demo/gogame/common/utilty"
	"demo/gogame/proto/router"
	"demo/gogame/rpc/client"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"log"
	"net/http"
	"sync"
	"time"
)

type RecvData struct {
	Uid   string `json:"uid"`
	Input string `json:"input"`
}

var (
	rpcRouterStream *rpcclient.RpcRouterStream
	rpcRouterCli    *rpcclient.RpcRouterClient
	rpcLoggerCli    *rpcclient.RpcLoggerClient

	users sync.Map
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
		//rpcRouterCli := &rpcclient.RpcRouterClient{}
		//if er := rpcRouterCli.Start(fmt.Sprintf("%s:%d", global.Cfg.RouterServ.Host, global.Cfg.RouterServ.Port)); er != nil {
		//	log.Panic(er)
		//}

		uuid := ggutilty.UUID()
		s.Set("uuid", uuid)
		users.Store(uuid, s)
		log.Println("websocket的连接成功", uuid)
	})

	m.HandleMessage(func(s *melody.Session, data []byte) {
		var rdata RecvData
		if er := json.Unmarshal(data, &rdata); er != nil {
			log.Println("解析输入信息失败:", er)
			er = s.CloseWithMsg([]byte("输入信息解析失败"))
			if er != nil {
				log.Println(er)
			}
			return
		}

		//v, ok := s.Get("client")
		//if !ok {
		//	log.Println("error")
		//	return
		//}
		//
		//routerCli := v.(*rpcclient.RpcRouterClient)
		//if !ok {
		//	er := s.CloseWithMsg([]byte("被保存的数据流不是[*client.RpcCenterStream]"))
		//	if er != nil {
		//		log.Println(er)
		//	}
		//	return
		//}
		//
		//resp, er := routerCli.ForwardingData(10000, 1, , string(data))
		//if er != nil {
		//	log.Println("error")
		//}
		//log.Println(resp)
		//
		//er := s.Write(data)
		//if er != nil {
		//	log.Println("error")
		//	return
		//}

		v, ok := s.Get("uuid")
		if !ok {
			log.Println("error")
			return
		}
		uuid := v.(string)

		if er := rpcRouterStream.SendMessage(10000, 1, uuid, string(data)); er != nil {
			er = s.CloseWithMsg([]byte("向gRPC服务端发送消息失败:" + er.Error()))
			return
		}
	})

	m.HandleDisconnect(func(s *melody.Session) {
		//v, ok := s.Get("client")
		//if !ok {
		//	log.Println("error")
		//	return
		//}
		//
		//routerCli := v.(*rpcclient.RpcRouterClient)
		//if er := routerCli.Stop(); er != nil {
		//	log.Println("断开stream错误", er)
		//}

		v, ok := s.Get("uuid")
		if !ok {
			log.Println("error")
			return
		}
		uuid := v.(string)

		users.Delete(uuid)

		log.Println("websocket的断开连接", uuid)
	})

	rpcLoggerCli = &rpcclient.RpcLoggerClient{}
	if er := rpcLoggerCli.Start(fmt.Sprintf("%s:%d", global.Cfg.LoggerServ.Host, global.Cfg.LoggerServ.Port)); er != nil {
		log.Panic(er)
	}

	rpcRouterCli = &rpcclient.RpcRouterClient{}
	if er := rpcRouterCli.Start(fmt.Sprintf("%s:%d", global.Cfg.RouterServ.Host, global.Cfg.RouterServ.Port)); er != nil {
		log.Panic(er)
	}

	var er error
	rpcRouterStream, er = rpcRouterCli.CreateStream(func(response *routersvr.ForwardResponse) {
		switch response.MainId {
		case 10000:
			log.Println(response)

			v, ok := users.Load(response.Uuid)
			if !ok {
				log.Println("session error", response.Uuid)
				return
			}
			s := v.(*melody.Session)

			er = s.Write([]byte(response.Output))
			if er != nil {
				log.Println("error")
				return
			}
		}
	})

	if er != nil {
		log.Panic(er)
	}

	Test()

	log.Panic(r.Run(":10010"))
}
