package gghub

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
	"tuyue/tuyue_common/network/ws/packet"
)

type HandlerFunc func(conn *websocket.Conn, pk *typacket.Packet)

var ErrNotFound = errors.New("hub: not found handler")
var ErrProto = errors.New("hub: protos error")
var EmptyMsg = errors.New("hub: empty msg")

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
	_, ok := h.events.Load(k)
	if !ok {
		h.events.Store(k, fn)
	}
}

func (h *Hub) RemoveHandle(mid, sid uint16) {
	key := hubKey{mid: mid, sid: sid}
	_, ok := h.events.Load(key)
	if ok {
		h.events.Delete(key)
	}
}

func (h *Hub) Query(mid, sid uint16) HandlerFunc {
	key := hubKey{mid: mid, sid: sid}
	if v, ok := h.events.Load(key); ok {
		return v.(HandlerFunc)
	}
	return nil
}

func (h *Hub) Close() {

}

func (h *Hub) DispatchMessage(conn *websocket.Conn, msg []byte) error {
	if len(msg) == 0 {
		return EmptyMsg
	}

	pk, er := typacket.NewPacketWithData(msg)
	if er != nil {
		return ErrProto
	}

	fn := h.Query(pk.Mid(), pk.Sid())
	if fn == nil {
		return ErrNotFound
	}

	fn(conn, pk)

	return nil
}
