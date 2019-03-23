package ggclis

import (
	"errors"
	"net/url"
	"sync"
	"time"
	"tuyue/tuyue_common/network/ws"

	log "github.com/alecthomas/log4go"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

type WSSClient struct {
	conn             *websocket.Conn
	hub              *Hub
	sendMutex        sync.Mutex
	done             chan struct{}
	heartTimeSecond  int
	isConnected      bool
	heartBeatHandler func() error
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

func (w *WSSClient) Open(scheme string, host string, path string) (err error) {
	if w == nil {
		return errors.New("object instance is empty")
	}

	u := url.URL{Scheme: scheme, Host: host, Path: path}

	w.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	w.isConnected = true

	go func(heartTimeSecond int) {
		defer func() {
			log.Error("websocket control exit")
		}()

		heartTicker := time.NewTicker(time.Second * time.Duration(heartTimeSecond))

	loop:
		for {
			select {
			case <-heartTicker.C: //TODO:心跳处理
				er := w.heartBeatHandler()
				//er := w.send(websocket.PingMessage, []byte{})
				if er != nil {
					log.Error(er)
				}
				break
			case <-w.done:
				w.sendMutex.Lock()
				var er error
				er = w.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if er != nil {
					log.Error(er)
				}
				er = w.conn.Close()
				if er != nil {
					log.Error(er)
				}
				w.sendMutex.Unlock()
				heartTicker.Stop()
				w.isConnected = false
				break loop
			}
		}
	}(w.heartTimeSecond)

	return nil
}

func (w *WSSClient) IsConnected() bool {
	return w.isConnected
}

func (w *WSSClient) SetHeartBeatHandler(h func() error) {
	if h == nil {
		h = func() error { return nil }
	}
	w.heartBeatHandler = h
}

func (w *WSSClient) send(mt int, buf []byte) error {
	if w == nil {
		return errors.New("object instance is empty")
	}

	w.sendMutex.Lock()
	defer w.sendMutex.Unlock()
	err := w.conn.WriteMessage(mt, buf)
	if err != nil {
		w.isConnected = false
		w.done <- struct{}{}
		return errors.New("error")
	}

	return nil
}

func (w *WSSClient) RegisterMessage(mid, sid uint16, f HandlerFunc) uint32 {
	w.hub.Handle(mid, sid, f)
	return 0
}

func (w *WSSClient) UnregisterMessage(mid, sid uint16) {
	w.hub.RemoveHandle(mid, sid)
}

func (w *WSSClient) SendMessage(mid, sid uint16, rquestId uint32, pb proto.Message) error {
	if w == nil {
		return errors.New("object instance is empty")
	}

	if !w.isConnected {
		return errors.New("websocket server is not connected")
	}

	pk := tyws.NewPacket(mid, sid, rquestId)
	err := pk.EncodeProto(pb)
	if err != nil {
		return err
	}

	err = w.send(websocket.BinaryMessage, pk.Data())
	if err != nil {
		return err
	}

	return nil
}

func (w *WSSClient) Run() error {
	if w == nil {
		return errors.New("object instance is empty")
	}

	if !w.isConnected {
		return errors.New("websocket server is not connected")
	}

	go func() {
		defer func() {
			log.Error("websocket read exit")
		}()

		for {
			mt, message, err := w.conn.ReadMessage()
			if err != nil {
				w.isConnected = false
				w.done <- struct{}{}
				log.Error(err)
				return
			}

			if mt != websocket.BinaryMessage {
				log.Warn("protocol error")
				return
			}

			err = w.hub.DispatchMessage(message)
			if err != nil {
				log.Warn(err)
			}
		}
	}()

	return nil
}

func (w *WSSClient) Stop() error {
	if w == nil {
		return errors.New("object instance is empty")
	}

	w.done <- struct{}{}

	return nil
}

func (w *WSSClient) Close() error {
	if w == nil {
		return errors.New("object instance is empty")
	}
	w.hub.Close()

	return nil
}
