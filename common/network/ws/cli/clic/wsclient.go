package ggclic

import (
	"crypto/tls"
	"errors"
	log "github.com/alecthomas/log4go"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"sync"
	"time"
	"tuyue/tuyue_common/network/ws/packet"
)

type envelope struct {
	mt  int
	msg []byte
}

var (
	TlsDialer = &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
		TLSClientConfig:  &tls.Config{InsecureSkipVerify: true},
	}
)

type WSClient struct {
	userData            interface{}
	mode                CallMode
	conn                *websocket.Conn
	hub                 *Hub
	rwmutex             sync.RWMutex
	done                chan struct{}
	output              chan *envelope
	heartTimeSecond     int
	cancelTimeoutSecond int
	open                bool
	heartBeatHandler    func() error
}

func DefaultClient(heartTimeSecond int, cancelTimeoutSecond int) *WSClient {
	if heartTimeSecond < 0 {
		heartTimeSecond = 0
	}

	if cancelTimeoutSecond < 0 {
		cancelTimeoutSecond = 0
	}

	return &WSClient{
		mode:                SYNC,
		hub:                 NewHub(SYNC),
		done:                make(chan struct{}, 1),
		heartTimeSecond:     heartTimeSecond,
		cancelTimeoutSecond: cancelTimeoutSecond,
	}
}

func DefaultAyncClient(heartTimeSecond int, queueSize int) *WSClient {
	if heartTimeSecond < 0 {
		heartTimeSecond = 0
	}

	return &WSClient{
		mode:                ASYNC,
		hub:                 NewHub(ASYNC),
		done:                make(chan struct{}, 1),
		heartTimeSecond:     heartTimeSecond,
		cancelTimeoutSecond: -1,
		output:              make(chan *envelope, queueSize),
	}
}

func (w *WSClient) Open(scheme string, host string, path string) (er error) {
	u := url.URL{Scheme: scheme, Host: host, Path: path}

	log.Info("connecting to %s", u.String())

	if scheme == "wss" {
		w.conn, _, er = TlsDialer.Dial(u.String(), nil)
	} else if scheme == "ws" {
		w.conn, _, er = websocket.DefaultDialer.Dial(u.String(), nil)
	} else {
		return errors.New("未知Scheme")
	}

	if er != nil {
		return er
	}

	w.open = true

	go func(heartTimeSecond int) {
		defer func() {
			log.Error("ws control exit")
		}()

		if w.Closed() {
			log.Error(errors.New("ws is closed"))
		}

		heartTicker := time.NewTicker(time.Second * time.Duration(heartTimeSecond))
		defer heartTicker.Stop()

	loop:
		for {
			select {
			case <-heartTicker.C: // TODO:心跳处理
				if w.mode == ASYNC {
					// TODO:有bug暂未修复
					w.output <- &envelope{mt: websocket.PingMessage, msg: []byte{}}
					break
				}

				if w.mode == SYNC {
					er = w.heartBeatHandler()
					// er := w.write(websocket.PingMessage, []byte{})
					if er != nil {
						log.Error(er)
					}
					break
				}
			case msg := <-w.output:
				if er = w.conn.WriteMessage(msg.mt, msg.msg); er != nil {
					w.open = false
					w.done <- struct{}{}
					log.Error(er)
				}
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

func (w *WSClient) UserData() interface{} {
	return w.userData
}

func (w *WSClient) SetUserData(userData interface{}) {
	w.userData = userData
}

func (w *WSClient) Closed() bool {
	w.rwmutex.Lock()
	defer w.rwmutex.Unlock()
	return !w.open
}

func (w *WSClient) HandleHeartbeat(h func() error) {
	if h == nil {
		h = func() error { return nil }
	}
	w.heartBeatHandler = h
}

func (w *WSClient) write(messageType int, buf []byte) error {
	w.rwmutex.Lock()
	defer w.rwmutex.Unlock()
	er := w.conn.WriteMessage(messageType, buf)
	if er != nil {
		w.open = false
		w.done <- struct{}{}
		return errors.New("error")
	}

	return nil
}

func (w *WSClient) RegisterMessage(mid, sid uint16) uint32 {
	_, rquestId := w.hub.AddHandle(mid, sid, nil)
	return rquestId
}

func (w *WSClient) UnregisterMessage(mid, sid uint16, clientId uint32) {
	w.hub.RemoveHandle(mid, sid, clientId)
}

func (w *WSClient) WriteProtoMessage(mid, sid uint16, pb proto.Message) (resp interface{}, er error) {
	if w.Closed() {
		return nil, errors.New("ws is closed")
	}

	if w.mode != SYNC {
		return nil, errors.New("CallMode error")
	}

	receiver, clientId := w.hub.AddHandle(mid, sid, nil)
	pk := typacket.NewPacket(mid, sid, clientId)
	data, er := proto.Marshal(pb)
	if er != nil {
		return nil, er
	}
	er = pk.Encode(data)
	if er != nil {
		return nil, errors.New("pb error")
	}

	er = w.write(websocket.BinaryMessage, pk.Data())
	if er != nil {
		return nil, errors.New(er.Error())
	}

	timeout := time.NewTimer(time.Second * time.Duration(w.cancelTimeoutSecond))
	select {
	case resp = <-receiver:
	case <-timeout.C: // TODO: 超时取消执行任务
		w.hub.CancelHandle(mid, sid, clientId)
		timeout.Stop()
		return nil, errors.New("accept websocket message timeout")
	}

	return resp, er
}

func (w *WSClient) AsynWriteProtoMessage(mid, sid uint16, pb proto.Message, f HandlerFunc) error {
	if w.Closed() {
		return errors.New("ws is closed")
	}

	if w.mode != ASYNC {
		return errors.New("调用模式错误")
	}

	_, clientId := w.hub.AddHandle(mid, sid, f)
	pk := typacket.NewPacket(mid, sid, clientId)
	data, er := proto.Marshal(pb)
	if er != nil {
		return er
	}
	er = pk.Encode(data)
	if er != nil {
		return er
	}
	w.output <- &envelope{mt: websocket.BinaryMessage, msg: pk.Data()}

	return nil
}

func (w *WSClient) Run() error {
	if w.Closed() {
		return errors.New("ws is closed")
	}

	go w.readMessage()

	return nil
}

func (w *WSClient) readMessage() {
	defer func() {
		log.Error("ws read exit")
	}()

	for {
		mt, message, er := w.conn.ReadMessage()
		if er != nil {
			w.open = false
			w.done <- struct{}{}
			log.Error(er)
			return
		}

		if mt != websocket.BinaryMessage {
			log.Warn("protocol error")
			return
		}

		er = w.hub.DispatchMessage(message)
		if er != nil {
			log.Warn(er)
		}

		ln := w.hub.Len()
		if ln > 0 {
			log.Warn("unconsumed message length:[%d]", ln)
		}
	}
}

func (w *WSClient) Stop() {
	w.done <- struct{}{}
}

func (w *WSClient) Close() {
	w.hub.Close()
}
