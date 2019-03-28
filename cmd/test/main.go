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

func CreateWs(uid int) {
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:10010", Path: "ws"}
	conn, _, er := websocket.DefaultDialer.Dial(u.String(), nil)
	if er != nil {
		log.Panic(er)
	}

	log.Println("连接成功", uid)

	go func() {
		m := make(map[string]string)
		suid := strconv.Itoa(uid)
		m["uid"] = suid
		m["input"] = strings.Repeat(suid, 2)

		tickHeartBeat := time.NewTicker(time.Second * time.Duration(30))
		t := rand.Intn(5) + 1
		tickSend := time.NewTicker(time.Second * time.Duration(t))
		for {
			select {
			case <-tickHeartBeat.C:
				er = conn.WriteMessage(websocket.PingMessage, []byte(""))
				if er != nil {
					log.Println(er)
				}
			case <-tickSend.C:
				var data []byte
				data, er = json.Marshal(m)
				if er != nil {
					log.Println(er)
				}

				er = conn.WriteMessage(websocket.TextMessage, data)
				if er != nil {
					log.Println(er)
					return
				}
			}
		}
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(message))
	}
}

func main() {
	for i := 10000; i < 10010; i++ {
		go CreateWs(i)
		time.Sleep(time.Microsecond * time.Duration(10))
	}
	select {}
}
