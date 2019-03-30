package main

import (
	"demo/gogame/cmd/test/config"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type FastWsClient struct {
	uid  int
	conn *websocket.Conn
}

var (
	cfg *config.JsonConfig
)

func loginMsg(uid int) ([]byte, error) {
	m := make(map[string]string)
	m["uid"] = strconv.Itoa(uid)
	m["mainId"] = "10000"
	m["subId"] = "0"
	m["msg"] = ""

	data, er := json.Marshal(m)
	if er != nil {
		log.Println(er)
		return nil, er
	}

	return data, nil
}

func chatMsg(uid int) ([]byte, error) {
	suid := strconv.Itoa(uid)
	m := make(map[string]string)
	m["uid"] = suid
	m["input"] = strings.Repeat(suid, 2)

	data, er := json.Marshal(m)
	if er != nil {
		log.Println(er)
		return nil, er
	}

	return data, nil
}

func (f *FastWsClient) Create(host, path string) error {
	u := url.URL{Scheme: "ws", Host: host, Path: path}
	var er error
	f.conn, _, er = websocket.DefaultDialer.Dial(u.String(), nil)
	if er != nil {
		//log.Println(er)
		return er
	}
	log.Printf("连接成功[uid=%d]", f.uid)
	return nil
}

func (f *FastWsClient) TestMessage() {
	tickHeartBeat := time.NewTicker(time.Second * time.Duration(45))
	tickSend := time.NewTicker(time.Millisecond * time.Duration(cfg.Millisecond))
	for {
		select {
		case <-tickHeartBeat.C:
			er := f.conn.WriteMessage(websocket.PingMessage, []byte(""))
			if er != nil {
				log.Println(er)
				return
			}
		case <-tickSend.C:
			data, er := chatMsg(f.uid)
			if er != nil {
				log.Println(er)
				return
			}

			er = f.conn.WriteMessage(websocket.TextMessage, data)
			if er != nil {
				log.Println(er)
				return
			}
		}
	}
}

func (f *FastWsClient) Run() {
	for {
		_, message, err := f.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(message))
	}
}

func main() {
	var er error
	cfg, er = config.LoadJsonConfig("./conf/conf.json")
	if er != nil {
		log.Panic(er)
	}

	for i := 1; i <= cfg.Count; i++ {
		ws := FastWsClient{uid: i}
		er = ws.Create(cfg.Host, cfg.Path)
		if er != nil {
			log.Println(er)
		} else {
			if cfg.Send {
				go ws.TestMessage()
				go ws.Run()
			}
		}
		time.Sleep(time.Microsecond * time.Duration(20))
	}
	select {}
}
