package ggclic

import (
	"container/list"
	"errors"
	"fmt"
	log "github.com/alecthomas/log4go"
	"sync"
	"tuyue/tuyue_common/network/ws/packet"
	"tuyue/tuyue_common/utils"
)

type HandlerFunc func(buf []byte)

type CallMode int

const (
	SYNC  CallMode = 0 //同步消息模式
	ASYNC CallMode = 1 //异步消息模式
)

type hubKey struct {
	mid uint16
	sid uint16
}

type handlerEvent struct {
	status   int         //0:正常，1:超时
	clientId uint32      //客户端Id
	fn       HandlerFunc //回调函数
	receiver chan interface{}
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

func (h *Handler) AddSyncHandler(fn HandlerFunc) (receiver <-chan interface{}, clientId uint32) {
	if h.model != SYNC {
		return nil, 0
	}

	h.m.Lock()
	defer h.m.Unlock()

	clientId = tyutils.HashCode(tyutils.UUID())
	var hv *handlerEvent
	hv = &handlerEvent{fn: fn, clientId: clientId, receiver: make(chan interface{}, 1)}
	h.ls.PushBack(hv)

	return hv.receiver, clientId
}

func (h *Handler) AddAsyncHandler(fn HandlerFunc) (clientId uint32) {
	if h.model != ASYNC {
		return 0
	}

	h.m.Lock()
	defer h.m.Unlock()

	var ev *handlerEvent
	clientId = tyutils.HashCode(tyutils.UUID())
	ev = &handlerEvent{fn: fn, clientId: clientId, receiver: nil}
	h.ls.PushBack(ev)
	return
}

func (h *Handler) Query(clientId uint32) *handlerEvent {
	h.m.Lock()
	defer h.m.Unlock()

	for e := h.ls.Front(); e != nil; e = e.Next() {
		ev := e.Value.(*handlerEvent)
		if ev.clientId == clientId {
			return ev
		}
	}

	return nil
}

func (h *Handler) Remove(clientId uint32) {
	h.RemoveWith(&handlerEvent{clientId: clientId})
}

func (h *Handler) RemoveWith(v *handlerEvent) {
	h.m.Lock()
	defer h.m.Unlock()
	for e := h.ls.Front(); e != nil; e = e.Next() {
		ev := e.Value.(*handlerEvent)
		if ev.clientId == v.clientId {
			if ev.receiver != nil {
				close(ev.receiver)
			}
			h.ls.Remove(e)
			break
		}
	}
}

func (h *Handler) Cancel(clientId uint32) {
	h.CancelWithOuttime(&handlerEvent{clientId: clientId})
}

func (h *Handler) CancelWithOuttime(v *handlerEvent) {
	h.m.Lock()
	defer h.m.Unlock()

	for e := h.ls.Front(); e != nil; e = e.Next() {
		ev := e.Value.(*handlerEvent)
		if ev.clientId == v.clientId {
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

func (h *Handler) Range(fn func(v *handlerEvent) bool) {
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
			if fn(ev) {
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

func (h *Hub) AddHandle(mid, sid uint16, fn HandlerFunc) (receiver <-chan interface{}, clientId uint32) {
	k := hubKey{mid: mid, sid: sid}
	v, ok := h.mfs.Load(k)
	if !ok {
		v = newHandler(h.model)
		h.mfs.Store(k, v)
	}

	if h.model == SYNC {
		receiver, clientId = v.(*Handler).AddSyncHandler(fn)
		return receiver, clientId
	}

	if h.model == ASYNC {
		clientId = v.(*Handler).AddAsyncHandler(fn)
		return nil, clientId
	}

	return nil, 0
}

func (h *Hub) RemoveHandle(mid, sid uint16, clientId uint32) {
	v, ok := h.mfs.Load(hubKey{mid: mid, sid: sid})
	if ok {
		v.(*Handler).Remove(clientId)
	}
}

func (h *Hub) CancelHandle(mid, sid uint16, clientId uint32) {
	v, ok := h.mfs.Load(hubKey{mid: mid, sid: sid})
	if ok {
		v.(*Handler).Cancel(clientId)
	}
}

func (h *Hub) Get(mid, sid uint16) *Handler {
	if v, ok := h.mfs.Load(hubKey{mid: mid, sid: sid}); ok {
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

	pk, er := typacket.NewPacketWithData(message)
	if er != nil {
		return er
	}

	hd := h.Get(pk.Mid(), pk.Sid())
	if hd == nil {
		return errors.New("protocol missing")
	}

	ev := hd.Query(pk.ClientId())
	if ev == nil {
		return errors.New(fmt.Sprintf("the request was cancelled due to timeout. detail:%+v", pk))
	}

	switch h.model {
	case SYNC:
		ev.receiver <- pk.Data()
	case ASYNC:
		if ev.fn != nil {
			ev.fn(pk.Data())
		} else {
		}
	default:
	}

	hd.RemoveWith(ev)

	return nil
}
