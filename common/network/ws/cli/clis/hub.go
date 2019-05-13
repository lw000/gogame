package ggclis

import (
	"errors"
	"sync"
	"tuyue/tuyue_common/network/ws/packet"
)

type HandlerFunc func(pk *typacket.Packet)

type hubKey struct {
	mid uint16
	sid uint16
}

type Hub struct {
	events sync.Map
}

func NewHub() *Hub {
	return &Hub{}
}

func (h *Hub) AddHandle(mid, sid uint16, fn HandlerFunc) {
	k := hubKey{mid: mid, sid: sid}
	if _, ok := h.events.Load(k); !ok {
		h.events.Store(k, fn)
	}
}

func (h *Hub) RemoveHandle(mid, sid uint16) {
	k := hubKey{mid: mid, sid: sid}
	if _, ok := h.events.Load(k); !ok {
		h.events.Delete(k)
	}
}

func (h *Hub) Query(mid, sid uint16) HandlerFunc {
	k := hubKey{mid: mid, sid: sid}
	if f, ok := h.events.Load(k); ok {
		return f.(HandlerFunc)
	}
	return nil
}

func (h *Hub) DispatchMessage(message []byte) error {
	if len(message) == 0 {
		return errors.New("message is empty")
	}

	pk, err := typacket.NewPacketWithData(message)
	if err != nil {
		return err
	}

	fn := h.Query(pk.Mid(), pk.Sid())
	if fn == nil {
		return errors.New("protocol missing")
	} else {
		fn(pk)
	}

	return nil
}

func (h *Hub) Close() {

}
