package main

import (
	"demo/gogame/cmd/gateway/global"
	"demo/gogame/common/utils"
	"demo/gogame/constant"
	"demo/gogame/pcl"
	"demo/gogame/protos"
	"demo/gogame/protos/router"
	"demo/gogame/rpc/client"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
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
	rpcRouter    *rpcclient.RpcRouterStreamClient
	rpcLoggerCli *rpcclient.RpcLoggerClient

	clients sync.Map
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

func onRouterMessage(resp *routersvr.ReponseMessage) {
	switch resp.MsgType {
	case 0:
		rpcRouter.SetClientUuid(resp.Cuuid)
		log.Printf("客户端注册成功(Clientuuid:%s)", resp.Cuuid)
	case 1:
		session, ok := clients.Load(resp.Uuid)
		if !ok {
			log.Printf("客户端不存在[%s]", resp.Uuid)
			return
		}
		s := session.(*melody.Session)
		if s.IsClosed() {
			log.Println("客户端已经关闭")
			return
		}

		pro := protocol.RequestChat{}
		if er := proto.Unmarshal(resp.Msg, &pro); er != nil {
			log.Println(er)
			return
		}

		log.Printf(`{"uid":%s "input":%s}`, pro.Uid, pro.Msg)

		er := s.Write([]byte(fmt.Sprintf(`{"uid":%s, "input":%s}`, pro.Uid, pro.Msg)))
		if er != nil {
			log.Println("error")
			return
		}
	}
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
			log.Println(er)
		}
		log.Println(c.Query("uid"))
	})

	m.HandleConnect(func(s *melody.Session) {
		uuid := ggutils.UUID()
		s.Set("uuid", uuid)
		clients.Store(uuid, s)
		log.Println("连接成功", uuid)
	})

	m.HandleMessage(func(s *melody.Session, data []byte) {
		var er error
		var rdata RecvData
		if er = json.Unmarshal(data, &rdata); er != nil {
			log.Println(er)
			er = s.CloseWithMsg([]byte(er.Error()))
			if er != nil {
				log.Println(er)
			}
			return
		}

		var pbData []byte
		pro := protocol.RequestChat{Uid: rdata.Uid, Msg: rdata.Input}
		pbData, er = proto.Marshal(&pro)
		if er != nil {
			log.Println(er)
			return
		}

		v, ok := s.Get("uuid")
		if !ok {
			log.Println("error")
			return
		}
		uuid := v.(string)

		if er = rpcRouter.SendMessage(uuid, pbData); er != nil {
			er = s.CloseWithMsg([]byte("路由转发消息失败" + er.Error()))
			return
		}
	})

	m.HandleDisconnect(func(s *melody.Session) {
		v, ok := s.Get("uuid")
		if !ok {
			log.Println("error")
			return
		}
		uuid := v.(string)

		clients.Delete(uuid)

		log.Println("websocket的断开连接", uuid)
	})

	var er error
	rpcLoggerCli = &rpcclient.RpcLoggerClient{}
	if er = rpcLoggerCli.Start(fmt.Sprintf("%s:%d", global.Cfg.LoggerServ.Host, global.Cfg.LoggerServ.Port)); er != nil {
		log.Panic(er)
	}

	rpcRouterCli := &rpcclient.RpcRouterClient{ServiceId: ggconstant.CGatewayServiceId, UUID: ggutils.UUID()}
	if er = rpcRouterCli.Start(fmt.Sprintf("%s:%d", global.Cfg.RouterServ.Host, global.Cfg.RouterServ.Port)); er != nil {
		log.Panic(er)
	}

	rpcRouter, er = rpcRouterCli.CreateStream(onRouterMessage)
	if er != nil {
		log.Panic(er)
	}

	{
		var data []byte
		data, er = ggpcl.LoadPcl("./conf/pcl.json")
		er = rpcRouter.RegisterService(data)
		if er != nil {
			log.Panic(er)
		}
	}

	Test()

	log.Panic(r.Run(fmt.Sprintf(":%d", global.Cfg.Port)))
}
