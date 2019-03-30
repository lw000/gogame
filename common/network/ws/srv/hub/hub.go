package gghub

import (
	"demo/gogame/common/network/ws"
	"errors"
	"github.com/gorilla/websocket"
	"sync"
)

type HandleFunc func(c *websocket.Conn, pk *ggwspk.Packet)

type Key struct {
	mid uint16
	sid uint16
}

type Hub struct {
	mf sync.Map
}

func NewWSHub() *Hub {
	return &Hub{}
}

func (h *Hub) Handle(mid, sid uint16, f HandleFunc) {
	k := Key{mid: mid, sid: sid}
	if _, ok := h.mf.Load(k); !ok {
		h.mf.Store(k, f)
	} else {
	}
}

func (h *Hub) RemoveHandle(mid, sid uint16) {
	k := Key{mid: mid, sid: sid}
	if _, ok := h.mf.Load(k); !ok {
		h.mf.Delete(k)
	} else {
	}
}

func (h *Hub) Get(mid, sid uint16) HandleFunc {
	k := Key{mid: mid, sid: sid}
	if f, ok := h.mf.Load(k); ok {
		return f.(HandleFunc)
	}
	return nil
}

func (h *Hub) DispatchMessage(c *websocket.Conn, message []byte) error {
	if len(message) == 0 {
		return errors.New("message is empty")
	}

	pk, err := ggwspk.NewPacketWithData(message)
	if err != nil {
		return err
	}

	f := h.Get(pk.Mid(), pk.Sid())
	if f == nil {
		return errors.New("protocol missing")
	} else {
		f(c, pk)
	}

	return nil
}
