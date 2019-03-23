package ggclic

import (
	"errors"
	log "github.com/alecthomas/log4go"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"net/url"
	"sync"
	"time"
	"tuyue/tuyue_common/network/ws"
)

type QueueItem struct {
	mt   int
	data []byte
}

type WSClient struct {
	mode                CallMode
	conn                *websocket.Conn
	hub                 *Hub
	sendMutex           sync.Mutex
	done                chan struct{}
	senderQueue         chan QueueItem
	heartTimeSecond     int
	cancelTimeoutSecond int
	isConnected         bool
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
		senderQueue:         make(chan QueueItem, queueSize),
	}
}

func (w *WSClient) Open(scheme string, host string, path string) (err error) {
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
				if w.mode == ASYNC {
					//TODO:有bug暂未修复
					w.senderQueue <- QueueItem{mt: websocket.PingMessage, data: []byte{}}
					break
				}

				if w.mode == SYNC {
					er := w.heartBeatHandler()
					//er := w.send(websocket.PingMessage, []byte{})
					if er != nil {
						log.Error(er)
					}
					break
				}
			case d := <-w.senderQueue:
				if er := w.conn.WriteMessage(d.mt, d.data); er != nil {
					w.isConnected = false
					w.done <- struct{}{}
					log.Error(er)
				}
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

func (w *WSClient) IsConnected() bool {
	return w.isConnected
}

func (w *WSClient) SetHeartBeatHandler(h func() error) {
	if h == nil {
		h = func() error { return nil }
	}
	w.heartBeatHandler = h
}

func (w *WSClient) send(messageType int, buf []byte) error {
	if w == nil {
		return errors.New("object instance is empty")
	}

	w.sendMutex.Lock()
	defer w.sendMutex.Unlock()
	err := w.conn.WriteMessage(messageType, buf)
	if err != nil {
		w.isConnected = false
		w.done <- struct{}{}
		return errors.New("error")
	}

	return nil
}

func (w *WSClient) RegisterMessage(mid, sid uint16) uint32 {
	_, rquestId := w.hub.Handle(mid, sid, nil)
	return rquestId
}

func (w *WSClient) UnregisterMessage(mid, sid uint16, rquestId uint32) {
	w.hub.RemoveHandle(mid, sid, rquestId)
}

func (w *WSClient) SendMessage(mid, sid uint16, pb proto.Message) (resp interface{}, err error) {
	if w == nil {
		return nil, errors.New("object instance is empty")
	}

	if !w.isConnected {
		return nil, errors.New("websocket server is not connected")
	}

	if w.mode != SYNC {
		return nil, errors.New("调用模式错误")
	}

	receiver, rquestId := w.hub.Handle(mid, sid, nil)
	pk := tyws.NewPacket(mid, sid, rquestId)
	err = pk.EncodeProto(pb)
	if err != nil {
		return nil, errors.New("pb error")
	}

	err = w.send(websocket.BinaryMessage, pk.Data())
	if err != nil {
		return nil, errors.New(err.Error())
	}

	timeout := time.NewTimer(time.Second * time.Duration(w.cancelTimeoutSecond))
	select {
	case resp = <-receiver:
	case <-timeout.C: // TODO: 超时取消执行任务
		w.hub.CancelHandle(mid, sid, rquestId)
		timeout.Stop()
		return nil, errors.New("accept websocket message timeout")
	}

	return resp, err
}

func (w *WSClient) AsynSendMessage(mid, sid uint16, pb proto.Message, f HandlerFunc) error {
	if w == nil {
		return errors.New("object instance is empty")
	}

	if !w.isConnected {
		return errors.New("websocket server is not connected")
	}

	if w.mode != ASYNC {
		return errors.New("调用模式错误")
	}

	_, rquestId := w.hub.Handle(mid, sid, f)
	pk := tyws.NewPacket(mid, sid, rquestId)
	err := pk.EncodeProto(pb)
	if err != nil {
		return err
	}
	w.senderQueue <- QueueItem{mt: websocket.BinaryMessage, data: pk.Data()}

	return nil
}

func (w *WSClient) Run() error {
	if w == nil {
		return errors.New("object instance is empty")
	}

	if !w.isConnected {
		return errors.New("websocket server is not connected")
	}

	go w.readMessage()

	return nil
}

func (w *WSClient) readMessage() {
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

		ln := w.hub.Len()
		if ln > 0 {
			log.Warn("unconsumed message length:[%d]", ln)
		}
	}
}

func (w *WSClient) Stop() error {
	if w == nil {
		return errors.New("object instance is empty")
	}

	w.done <- struct{}{}

	return nil
}

func (w *WSClient) Close() error {
	if w == nil {
		return errors.New("object instance is empty")
	}
	w.hub.Close()

	return nil
}
