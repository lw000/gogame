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
	tag            int
	conn           *websocket.Conn
	onMessage      func([]byte)
	onConnected    func()
	onDisConnected func()
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

func (fc *FastWsClient) Create(host, path string) error {
	u := url.URL{Scheme: "ws", Host: host, Path: path}
	var er error
	fc.conn, _, er = websocket.DefaultDialer.Dial(u.String(), nil)
	if er != nil {
		return er
	}

	fc.onConnected()

	return nil
}

func (fc *FastWsClient) HandleConnected(f func()) {
	fc.onConnected = f
}

func (fc *FastWsClient) HandleDisConnected(f func()) {
	fc.onDisConnected = f
}

func (fc *FastWsClient) HandleMessage(f func(data []byte)) {
	fc.onMessage = f
}

func (fc *FastWsClient) SendMessage(data []byte) error {
	er := fc.conn.WriteMessage(websocket.TextMessage, data)
	if er != nil {
		log.Println(er)
		return er
	}
	return nil
}

func (fc *FastWsClient) Ping() error {
	er := fc.conn.WriteMessage(websocket.PingMessage, []byte(""))
	if er != nil {
		log.Println(er)
		return er
	}
	return nil
}

func (fc *FastWsClient) Run() {
	for {
		_, message, err := fc.conn.ReadMessage()
		if err != nil {
			log.Println(err)

			fc.onDisConnected()

			return
		}
		fc.onMessage(message)
	}
}

func TestMessage(fc *FastWsClient, uid int) {
	tickHeartBeat := time.NewTicker(time.Second * time.Duration(45))
	tickSend := time.NewTicker(time.Millisecond * time.Duration(cfg.Millisecond))
	for {
		select {
		case <-tickHeartBeat.C:
			er := fc.Ping()
			if er != nil {
				log.Println(er)
				return
			}
		case <-tickSend.C:
			data, er := chatMsg(uid)
			if er != nil {
				log.Println(er)
				return
			}

			er = fc.SendMessage(data)
			if er != nil {
				log.Println(er)
				return
			}
		}
	}
}

func main() {
	var er error
	cfg, er = config.LoadJsonConfig("./conf/conf.json")
	if er != nil {
		log.Panic(er)
	}

	for i := 1; i <= cfg.Count; i++ {
		ws := &FastWsClient{}
		ws.HandleConnected(func() {
			log.Printf("connected [uid=%d]", i)
		})

		ws.HandleDisConnected(func() {
			log.Println("disconnected")
		})

		ws.HandleMessage(func(data []byte) {
			log.Println(string(data))
		})
		er = ws.Create(cfg.Host, cfg.Path)
		if er != nil {
			log.Println(er)
			continue
		}

		if cfg.Send {
			go TestMessage(ws, i)
			go ws.Run()
		}
		time.Sleep(time.Microsecond * time.Duration(20))
	}
	select {}
}
