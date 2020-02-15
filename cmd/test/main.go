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

	data, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return data, nil
}

func chatMsg(uid int) ([]byte, error) {
	suid := strconv.Itoa(uid)
	m := make(map[string]string)
	m["uid"] = suid
	m["input"] = strings.Repeat(suid, 2)

	data, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return data, nil
}

func (fc *FastWsClient) Create(host, path string) error {
	u := url.URL{Scheme: "ws", Host: host, Path: path}
	var err error
	fc.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
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
	err := fc.conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (fc *FastWsClient) Ping() error {
	err := fc.conn.WriteMessage(websocket.PingMessage, []byte(""))
	if err != nil {
		log.Println(err)
		return err
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
			err := fc.Ping()
			if err != nil {
				log.Println(err)
				return
			}
		case <-tickSend.C:
			data, err := chatMsg(uid)
			if err != nil {
				log.Println(err)
				return
			}

			err = fc.SendMessage(data)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func main() {
	var err error
	cfg, err = config.LoadJsonConfig("./conf/conf.json")
	if err != nil {
		log.Panic(err)
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
		err = ws.Create(cfg.Host, cfg.Path)
		if err != nil {
			log.Println(err)
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
