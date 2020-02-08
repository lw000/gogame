package packet

import (
	"bytes"
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
	n, err := nb.buf.Write(buf)
	if err != nil {
	}

	return n
}

func (nb *NetBuffer) Read(n int) []byte {
	buf := make([]byte, n)
	n, err := nb.buf.Read(buf)
	if err != nil {
	}
	if n > 0 {
		return buf
	}
	return nil
}
