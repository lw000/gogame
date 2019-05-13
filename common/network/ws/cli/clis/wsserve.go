package ggclis

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
	"tuyue/tuyue_common/network/ws/packet"

	log "github.com/alecthomas/log4go"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

var (
	TlsDialer = &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
		TLSClientConfig:  &tls.Config{InsecureSkipVerify: true},
	}
)

type WSSClient struct {
	conn            *websocket.Conn
	hub             *Hub
	rwmutex         sync.RWMutex
	done            chan struct{}
	heartTimeSecond int
	open            bool
	heartBeatFn     func() error
}

func DefaultClient(heartTimeSecond int) *WSSClient {
	if heartTimeSecond < 0 {
		heartTimeSecond = 0
	}

	return &WSSClient{
		hub:             NewHub(),
		done:            make(chan struct{}, 1),
		heartTimeSecond: heartTimeSecond,
	}
}

func (w *WSSClient) Closed() bool {
	w.rwmutex.Lock()
	defer w.rwmutex.Unlock()
	return !w.open
}

func (w *WSSClient) Open(scheme string, host string, path string) (er error) {
	u := url.URL{Scheme: scheme, Host: host, Path: path}

	log.Info("connecting to %s", u.String())

	if scheme == "wss" {
		w.conn, _, er = TlsDialer.Dial(u.String(), nil)
	} else if scheme == "ws" {
		w.conn, _, er = websocket.DefaultDialer.Dial(u.String(), nil)
	} else {
		return errors.New(fmt.Sprintf("未知Scheme:%s", scheme))
	}

	if er != nil {
		log.Error(er)
		return er
	}

	w.open = true

	go func(heartTimeSecond int) {
		defer func() {
			log.Error("ws control exit")
		}()

		heartTicker := time.NewTicker(time.Second * time.Duration(heartTimeSecond))
		defer heartTicker.Stop()

	loop:
		for {
			select {
			case <-heartTicker.C: //TODO:心跳处理
				er = w.heartBeatFn()
				//er := w.send(websocket.PingMessage, []byte{})
				if er != nil {
					log.Error(er)
				}
				break
			case <-w.done:

				w.rwmutex.Lock()
				er = w.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if er != nil {
					log.Error(er)
				}
				er = w.conn.Close()
				if er != nil {
					log.Error(er)
				}
				w.open = false
				w.rwmutex.Unlock()
				break loop
			}
		}
	}(w.heartTimeSecond)

	return nil
}

func (w *WSSClient) HanldeHeartbeat(fn func() error) {
	if fn == nil {
		fn = func() error { return nil }
	}
	w.heartBeatFn = fn
}

func (w *WSSClient) write(mt int, buf []byte) error {
	w.rwmutex.Lock()
	defer w.rwmutex.Unlock()
	er := w.conn.WriteMessage(mt, buf)
	if er != nil {
		w.open = false
		w.done <- struct{}{}
		return errors.New("error")
	}

	return nil
}

func (w *WSSClient) AddHandle(mid, sid uint16, f HandlerFunc) {
	w.hub.AddHandle(mid, sid, f)
}

func (w *WSSClient) RemoveHandle(mid, sid uint16) {
	w.hub.RemoveHandle(mid, sid)
}

func (w *WSSClient) WriteBinaryMessage(mid, sid uint16, clientId uint32, data []byte) error {
	if w.Closed() {
		return errors.New("ws is closed")
	}

	pk := typacket.NewPacket(mid, sid, clientId)
	er := pk.Encode(data)
	if er != nil {
		log.Error(er)
		return er
	}

	er = w.write(websocket.BinaryMessage, pk.Data())
	if er != nil {
		log.Error(er)
		return er
	}

	return nil
}

func (w *WSSClient) WriteProtoMessage(mid, sid uint16, clientId uint32, pb proto.Message) error {
	if w.Closed() {
		return errors.New("ws is closed")
	}

	pk := typacket.NewPacket(mid, sid, clientId)
	data, er := proto.Marshal(pb)
	if er != nil {
		return er
	}
	er = pk.Encode(data)
	if er != nil {
		log.Error(er)
		return er
	}

	er = w.write(websocket.BinaryMessage, pk.Data())
	if er != nil {
		log.Error(er)
		return er
	}

	return nil
}

func (w *WSSClient) Run() error {
	if w.Closed() {
		return errors.New("ws is closed")
	}

	go func() {
		defer func() {
			log.Error("ws readMessage exit")
		}()

		for {
			mt, message, er := w.conn.ReadMessage()
			if er != nil {
				w.done <- struct{}{}
				log.Error(er)
				return
			}

			if mt != websocket.BinaryMessage {
				log.Error("protocol error")
				return
			}

			er = w.hub.DispatchMessage(message)
			if er != nil {
				log.Error(er)
			}
		}
	}()

	return nil
}

func (w *WSSClient) Stop() {
	w.done <- struct{}{}
}

func (w *WSSClient) Close() {
	w.hub.Close()
}
