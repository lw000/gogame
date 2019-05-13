package ggpacket

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type Packet struct {
	ver       uint8
	mid       uint16
	sid       uint16
	checkCode uint32
	clientId  uint32
	data      []byte
}

// var (
// 	pool *sync.Pool
// )
//
// func init() {
// 	pool = &sync.Pool{New: func() interface{} {
// 		return &Packet{}
// 	}}
// }

func NewPacket(mid, sid uint16, clientId uint32) *Packet {
	p := &Packet{
		mid:      mid,
		sid:      sid,
		clientId: clientId,
	}
	return p
}

func NewPacketWithData(data []byte) (*Packet, error) {
	if len(data) == 0 {
		return nil, errors.New("data item is zero")
	}
	p := &Packet{}
	buf := bytes.NewBuffer(data)
	er := p.readHead(buf)
	if er != nil {
		return nil, er
	}
	p.data = buf.Bytes()
	return p, nil
}

func (p *Packet) writeHead(buf *bytes.Buffer) (er error) {
	if er = binary.Write(buf, binary.LittleEndian, p.ver); er != nil {
		return er
	}
	if er = binary.Write(buf, binary.LittleEndian, p.checkCode); er != nil {
		return er
	}
	if er = binary.Write(buf, binary.LittleEndian, p.mid); er != nil {
		return er
	}
	if er = binary.Write(buf, binary.LittleEndian, p.sid); er != nil {
		return er
	}
	if er = binary.Write(buf, binary.LittleEndian, p.clientId); er != nil {
		return er
	}
	return er
}

func (p *Packet) readHead(buf *bytes.Buffer) (er error) {
	if er = binary.Read(buf, binary.LittleEndian, &p.ver); er != nil {
		return er
	}
	if er = binary.Read(buf, binary.LittleEndian, &p.checkCode); er != nil {
		return er
	}
	if er = binary.Read(buf, binary.LittleEndian, &p.mid); er != nil {
		return er
	}
	if er = binary.Read(buf, binary.LittleEndian, &p.sid); er != nil {
		return er
	}
	if er = binary.Read(buf, binary.LittleEndian, &p.clientId); er != nil {
		return er
	}
	return er
}

// Encode 编码数据包
func (p *Packet) Encode(data []byte) error {
	buf := &bytes.Buffer{}
	er := p.writeHead(buf)
	if er != nil {
		return er
	}

	if len(data) > 0 {
		var n int
		n, er = buf.Write(data)
		if er != nil {
			return er
		}
		if n < 0 {

		}
	}

	p.data = buf.Bytes()

	return nil
}

// EncodeProto 编码数据包
// func (p *Packet) EncodeProto(pb proto.Message) error {
// 	if pb == nil {
// 		er := p.Encode([]byte{})
// 		if er != nil {
// 			return er
// 		}
// 	} else {
// 		data, er := proto.Marshal(pb)
// 		if er != nil {
// 			return er
// 		}
// 		er = p.Encode(data)
// 		if er != nil {
// 			return er
// 		}
// 	}
// 	return nil
// }

func (p Packet) Ver() uint8 {
	return p.ver
}

func (p *Packet) CheckCode() uint32 {
	return p.checkCode
}

func (p Packet) Mid() uint16 {
	return p.mid
}

func (p Packet) Sid() uint16 {
	return p.sid
}

func (p *Packet) ClientId() uint32 {
	return p.clientId
}

func (p Packet) Data() []byte {
	return p.data
}

func (p Packet) String() string {
	return fmt.Sprintf("{ver:%d ccode:%d mid:%d sid:%d clientId:%d datalen:%d}", p.ver, p.checkCode, p.mid, p.sid, p.clientId, len(p.data))
}
