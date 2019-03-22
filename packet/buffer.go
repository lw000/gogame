package packet

import (
	"bytes"
	log "github.com/alecthomas/log4go"
)

type NetBuffer struct {
	buf *bytes.Buffer
}

func NewNetBuffer() *NetBuffer {
	return &NetBuffer{
		buf: bytes.NewBuffer(nil),
	}
}

func (nb *NetBuffer) Add(buf []byte) int {
	n, er := nb.buf.Write(buf)
	if er != nil {
		log.Error(er)
	}

	return n
}

func (nb *NetBuffer) Read(n int) []byte {
	buf := make([]byte, n)
	n, er := nb.buf.Read(buf)
	if er != nil {
		log.Error(er)
	}
	if n > 0 {
		return buf
	}
	return nil
}
