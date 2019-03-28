package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"strconv"
	"strings"
	"syncthing/lib/rand"
	"time"
)

type FastWsClient struct {
	uid  int
	conn *websocket.Conn
}

func login(uid int) ([]byte, error) {
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

func chat(uid int) ([]byte, error) {
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

func (f *FastWsClient) Create() error {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:10010", Path: "ws"}
	var er error
	f.conn, _, er = websocket.DefaultDialer.Dial(u.String(), nil)
	if er != nil {
		log.Panic(er)
		return er
	}
	log.Printf("连接成功[uid=%d]", f.uid)
	return nil
}

func (f *FastWsClient) SendMessage() {
	tickHeartBeat := time.NewTicker(time.Second * time.Duration(30))
	t := rand.Intn(5) + 1
	tickSend := time.NewTicker(time.Microsecond * time.Duration(t))
	for {
		select {
		case <-tickHeartBeat.C:
			er := f.conn.WriteMessage(websocket.PingMessage, []byte(""))
			if er != nil {
				log.Println(er)
			}
		case <-tickSend.C:
			{
				data, er := chat(f.uid)
				if er != nil {
					log.Println(er)
				}
				er = f.conn.WriteMessage(websocket.TextMessage, data)
				if er != nil {
					log.Println(er)
					return
				}
			}

			{

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
	for i := 10000; i < 20000; i++ {
		ws := FastWsClient{uid: i}
		er := ws.Create()
		if er != nil {
			log.Println(er)
		} else {
			go ws.SendMessage()
			go ws.Run()
		}
		time.Sleep(time.Microsecond * time.Duration(10))
	}
	select {}
}
