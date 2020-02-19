package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/lw000/gocommon/utils"
	"github.com/olahol/melody"
	"gogame/cmd/gateway/global"
	"gogame/constant"
	"gogame/pcl"
	"gogame/protos"
	"gogame/protos/router"
	"gogame/rpc/client"
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
	// 测试日志写入服务
	go func() {
		for {
			err := rpcLoggerCli.WriteLogger("gateway-1")
			if err != nil {
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
		if err := proto.Unmarshal(resp.Msg, &pro); err != nil {
			log.Println(err)
			return
		}

		log.Printf(`{"uid":%s "input":%s}`, pro.Uid, pro.Msg)

		err := s.Write([]byte(fmt.Sprintf(`{"uid":%s, "input":%s}`, pro.Uid, pro.Msg)))
		if err != nil {
			log.Println("error")
			return
		}
	}
}

func main() {
	if err := global.LoadGlobalConfig(); err != nil {
		log.Panic(err)
	}

	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "html/index.html")
	})

	r.GET("ws", func(c *gin.Context) {
		err := m.HandleRequest(c.Writer, c.Request)
		if err != nil {
			log.Println(err)
		}
		log.Println(c.Query("uid"))
	})

	m.HandleConnect(func(s *melody.Session) {
		uuid := tyutils.UUID()
		s.Set("uuid", uuid)
		clients.Store(uuid, s)
		log.Println("连接成功", uuid)
	})

	m.HandleMessage(func(s *melody.Session, data []byte) {
		var err error
		var rdata RecvData
		if err = json.Unmarshal(data, &rdata); err != nil {
			log.Println(err)
			err = s.CloseWithMsg([]byte(err.Error()))
			if err != nil {
				log.Println(err)
			}
			return
		}

		var pbData []byte
		pro := protocol.RequestChat{Uid: rdata.Uid, Msg: rdata.Input}
		pbData, err = proto.Marshal(&pro)
		if err != nil {
			log.Println(err)
			return
		}

		v, ok := s.Get("uuid")
		if !ok {
			log.Println("error")
			return
		}
		uuid := v.(string)

		if err = rpcRouter.SendMessage(uuid, pbData); err != nil {
			err = s.CloseWithMsg([]byte("路由转发消息失败" + err.Error()))
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

	var err error
	rpcLoggerCli = &rpcclient.RpcLoggerClient{}
	if err = rpcLoggerCli.Start(fmt.Sprintf("%s:%d", global.Cfg.LoggerServe.Host, global.Cfg.LoggerServe.Port)); err != nil {
		log.Panic(err)
	}

	rpcRouterCli := &rpcclient.RpcRouterClient{ServiceId: ggconstant.CGatewayServiceId, UUID: tyutils.UUID()}
	if err = rpcRouterCli.Start(fmt.Sprintf("%s:%d", global.Cfg.RouterServe.Host, global.Cfg.RouterServe.Port)); err != nil {
		log.Panic(err)
	}

	rpcRouter, err = rpcRouterCli.CreateStream(onRouterMessage)
	if err != nil {
		log.Panic(err)
	}

	{
		var data []byte
		data, err = ggpcl.LoadPcl("./conf/pcl.json")
		err = rpcRouter.RegisterService(data)
		if err != nil {
			log.Panic(err)
		}
	}

	Test()

	log.Panic(r.Run(fmt.Sprintf(":%d", global.Cfg.Port)))
}
