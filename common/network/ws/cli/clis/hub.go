package tyclis

import (
	"errors"
	"sync"
	"tuyue/tuyue_common/network/ws"
)

type HandlerFunc func(pk *tyws.Packet)

type Key struct {
	mid uint16
	sid uint16
}

type Hub struct {
	mf sync.Map
}

func NewHub() *Hub {
	return &Hub{}
}

func (h *Hub) Handle(mid, sid uint16, f HandlerFunc) {
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

func (h *Hub) Get(mid, sid uint16) HandlerFunc {
	k := Key{mid: mid, sid: sid}
	if f, ok := h.mf.Load(k); ok {
		return f.(HandlerFunc)
	}
	return nil
}

func (h *Hub) DispatchMessage(message []byte) error {
	if len(message) == 0 {
		return errors.New("message is empty")
	}

	pk, err := tyws.NewPacketWithData(message)
	if err != nil {
		return err
	}

	f := h.Get(pk.Mid(), pk.Sid())
	if f == nil {
		return errors.New("protocol missing")
	} else {
		f(pk)
	}

	return nil
}

func (h *Hub) Close() {

}
