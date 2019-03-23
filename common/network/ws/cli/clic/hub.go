package ggclic

import (
	"container/list"
	"errors"
	"fmt"
	log "github.com/alecthomas/log4go"
	"sync"
	"time"
	"tuyue/tuyue_common/network/ws"
	"tuyue/tuyue_common/utilty"
)

type HandlerFunc func(buf []byte)

type CallMode int

const (
	SYNC  CallMode = 0 //同步消息模式
	ASYNC CallMode = 1 //异步消息模式
)

type key struct {
	mid uint16
	sid uint16
}

type handlerEvent struct {
	eventType   int         //0:常驻,
	status      int         //0:正常，1:超时
	requestId   uint32      //事件Id
	requestTime int64       //请求时间戳
	reponseTime int64       //回应时间戳
	f           HandlerFunc //回调函数
	receiver    chan interface{}
}

type Handler struct {
	model CallMode // 0：同步 1：异步
	ls    *list.List
	m     sync.RWMutex
}

type Hub struct {
	model CallMode // 0：同步 1：异步
	mfs   sync.Map
}

func newHandler(model CallMode) *Handler {
	return &Handler{
		model: model,
		m:     sync.RWMutex{},
		ls:    list.New(),
	}
}

func (h *Handler) Add(v *handlerEvent) {
	h.m.Lock()
	defer h.m.Unlock()

	h.ls.PushBack(v)
}

func (h *Handler) AddSyncHandler(f HandlerFunc) (receiver <-chan interface{}, requestId uint32) {
	if h.model != SYNC {
		return nil, 0
	}

	h.m.Lock()
	defer h.m.Unlock()

	requestId = tyutilty.HashCode(tyutilty.UUID())
	var hv *handlerEvent
	hv = &handlerEvent{f: f, requestId: requestId, requestTime: time.Now().Unix(), receiver: make(chan interface{}, 1)}
	h.ls.PushBack(hv)

	return hv.receiver, requestId
}

func (h *Handler) AddAsyncHandler(f HandlerFunc) (requestId uint32) {
	if h.model != ASYNC {
		return 0
	}

	h.m.Lock()
	defer h.m.Unlock()

	var ev *handlerEvent
	requestId = tyutilty.HashCode(tyutilty.UUID())
	ev = &handlerEvent{f: f, requestId: requestId, requestTime: time.Now().Unix(), receiver: nil}
	h.ls.PushBack(ev)
	return
}

func (h *Handler) Query(requestId uint32) *handlerEvent {
	h.m.Lock()
	defer h.m.Unlock()

	for e := h.ls.Front(); e != nil; e = e.Next() {
		ev := e.Value.(*handlerEvent)
		if ev.requestId == requestId {
			return ev
		}
	}

	return nil
}

func (h *Handler) Remove(requestId uint32) {
	h.RemoveWith(&handlerEvent{requestId: requestId})
}

func (h *Handler) RemoveWith(v *handlerEvent) {
	h.m.Lock()
	defer h.m.Unlock()
	for e := h.ls.Front(); e != nil; e = e.Next() {
		ev := e.Value.(*handlerEvent)
		if ev.requestId == v.requestId {
			if ev.receiver != nil {
				close(ev.receiver)
			}
			h.ls.Remove(e)
			break
		}
	}
}

func (h *Handler) Cancel(requestId uint32) {
	h.CancelWithOuttime(&handlerEvent{requestId: requestId})
}

func (h *Handler) CancelWithOuttime(v *handlerEvent) {
	h.m.Lock()
	defer h.m.Unlock()

	for e := h.ls.Front(); e != nil; e = e.Next() {
		ev := e.Value.(*handlerEvent)
		if ev.requestId == v.requestId {
			ev.status = 1
			if ev.receiver != nil {
				close(ev.receiver)
			}
			log.Info("%+v", ev)
			h.ls.Remove(e)
			break
		}
	}
}

func (h *Handler) Clear() {
	h.m.Lock()
	defer h.m.Unlock()
	for e := h.ls.Front(); e != nil; {
		ev := e.Value.(*handlerEvent)
		if ev.receiver != nil {
			close(ev.receiver)
		}
		next := e.Next()
		h.ls.Remove(e)
		e = next
	}
}

func (h *Handler) Range(f func(v *handlerEvent) bool) {
	h.m.RLock()
	defer h.m.RUnlock()

	for e := h.ls.Front(); e != nil; {
		ev := e.Value.(*handlerEvent)
		if ev.status == 1 {
			if ev.receiver != nil {
				close(ev.receiver)
			}

			next := e.Next()
			h.ls.Remove(e)
			e = next
		}
	}

	for e := h.ls.Front(); e != nil; e = e.Next() {
		ev := e.Value.(*handlerEvent)
		if ev.status == 0 {
			if f(ev) {
				if ev.receiver != nil {
					close(ev.receiver)
				}
				h.ls.Remove(e)
				break
			}
		}
	}
}

func (h *Handler) Close() {
	h.m.RLock()
	defer h.m.RUnlock()

	for e := h.ls.Front(); e != nil; {
		ev := e.Value.(*handlerEvent)
		if ev.receiver != nil {
			close(ev.receiver)
		}

		next := e.Next()
		h.ls.Remove(e)
		e = next
	}
}

func (h *Handler) Len() int {
	h.m.RLock()
	defer h.m.RUnlock()

	return h.ls.Len()
}

func NewHub(model CallMode) *Hub {
	return &Hub{model: model}
}

func (h *Hub) Handle(mid, sid uint16, f HandlerFunc) (receiver <-chan interface{}, requestId uint32) {
	k := key{mid: mid, sid: sid}
	v, ok := h.mfs.Load(k)
	if !ok {
		v = newHandler(h.model)
		h.mfs.Store(k, v)
	}

	if h.model == SYNC {
		receiver, requestId = v.(*Handler).AddSyncHandler(f)
		return receiver, requestId
	}

	if h.model == ASYNC {
		requestId = v.(*Handler).AddAsyncHandler(f)
		return nil, requestId
	}

	return nil, 0
}

func (h *Hub) RemoveHandle(mid, sid uint16, requestId uint32) {
	v, ok := h.mfs.Load(key{mid: mid, sid: sid})
	if ok {
		v.(*Handler).Remove(requestId)
	}
}

func (h *Hub) CancelHandle(mid, sid uint16, requestId uint32) {
	v, ok := h.mfs.Load(key{mid: mid, sid: sid})
	if ok {
		v.(*Handler).Cancel(requestId)
	}
}

func (h *Hub) Get(mid, sid uint16) *Handler {
	if v, ok := h.mfs.Load(key{mid: mid, sid: sid}); ok {
		return v.(*Handler)
	}
	return nil
}

func (h *Hub) Len() int {
	var msglen int
	h.mfs.Range(func(key interface{}, value interface{}) bool {
		msglen += value.(*Handler).Len()
		return false
	})
	return msglen
}

func (h *Hub) Close() {
	h.mfs.Range(func(key interface{}, value interface{}) bool {
		value.(*Handler).Range(func(e *handlerEvent) bool {
			close(e.receiver)
			return false
		})
		return false
	})
}

func (h *Hub) DispatchMessage(message []byte) error {
	defer func() {
		if err := recover(); err != nil {
		}
	}()

	if len(message) == 0 {
		return errors.New("message is empty")
	}

	pk, err := tyws.NewPacketWithData(message)
	if err != nil {
		return err
	}

	hd := h.Get(pk.Mid(), pk.Sid())
	if hd == nil {
		return errors.New("protocol missing")
	}

	ev := hd.Query(pk.RequestId())
	if ev == nil {
		return errors.New(fmt.Sprintf("the request was cancelled due to timeout. detail:%+v", pk))
	}

	//更新响应时间
	ev.reponseTime = time.Now().Unix()

	switch h.model {
	case SYNC:
		ev.receiver <- pk.Data()
	case ASYNC:
		if ev.f != nil {
			ev.f(pk.Data())
		} else {
		}
	default:
	}

	hd.RemoveWith(ev)

	return nil
}
