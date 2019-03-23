package tyws

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/golang/protobuf/proto"
)

type Packet struct {
	ver       uint8
	ccode     uint32
	mid       uint16
	sid       uint16
	requestId uint32
	data      []byte
}

type JSonPacket struct {
	ver       uint8
	ccode     uint32
	mid       uint16
	sid       uint16
	requestId uint32
	data      []byte
}

var (
	pool *sync.Pool
)

func init() {
	pool = &sync.Pool{New: func() interface{} {
		return &Packet{}
	}}
}

func NewPacket(mid, sid uint16, requestId uint32) *Packet {
	p := &Packet{
		mid:       mid,
		sid:       sid,
		requestId: requestId,
	}
	return p
}

func NewPacketWithData(data []byte) (*Packet, error) {
	if len(data) == 0 {
		return nil, errors.New("data item is zero")
	}
	p := &Packet{}
	buf := bytes.NewBuffer(data)
	err := p.readHead(buf)
	if err != nil {
		return nil, err
	}
	p.data = buf.Bytes()
	return p, nil
}

func (p *Packet) writeHead(buf *bytes.Buffer) (err error) {
	if err = binary.Write(buf, binary.LittleEndian, p.ver); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.LittleEndian, p.ccode); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.LittleEndian, p.mid); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.LittleEndian, p.sid); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.LittleEndian, p.requestId); err != nil {
		return err
	}
	return err
}

func (p *Packet) readHead(buf *bytes.Buffer) (err error) {
	if err = binary.Read(buf, binary.LittleEndian, &p.ver); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &p.ccode); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &p.mid); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &p.sid); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &p.requestId); err != nil {
		return err
	}
	return err
}

//Encode 编码数据包
func (p *Packet) encode(data []byte) error {
	buf := &bytes.Buffer{}
	err := p.writeHead(buf)
	if err != nil {
		return err
	}

	if len(data) > 0 {
		n, err := buf.Write(data)
		if err != nil {
			return err
		}
		if n < 0 {

		}
	}

	p.data = buf.Bytes()

	return nil
}

//EncodeProto 编码数据包
func (p *Packet) EncodeProto(pb proto.Message) error {
	if pb == nil {
		err := p.encode([]byte{})
		if err != nil {
			return err
		}
	} else {
		data, err := proto.Marshal(pb)
		if err != nil {
			return err
		}
		err = p.encode(data)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p Packet) Ver() uint8 {
	return p.ver
}

func (p Packet) Ccode() uint32 {
	return p.ccode
}

func (p Packet) Mid() uint16 {
	return p.mid
}

func (p Packet) Sid() uint16 {
	return p.sid
}

func (p Packet) RequestId() uint32 {
	return p.requestId
}

func (p Packet) Data() []byte {
	return p.data
}

func (p Packet) String() string {
	return fmt.Sprintf("{ver:%d ccode:%d mid:%d sid:%d requestId:%d datalen:%d}", p.ver, p.ccode, p.mid, p.sid, p.requestId, len(p.data))
}

func NewJsonPacket(mid, sid uint16, requestId uint32) *JSonPacket {
	return &JSonPacket{
		ver:       0,
		ccode:     0,
		mid:       mid,
		sid:       sid,
		requestId: requestId,
	}
}

func NewJsonPacketWithData(data []byte) *JSonPacket {
	if len(data) == 0 {
		return nil
	}
	return &JSonPacket{
		ver:       0,
		ccode:     0,
		mid:       0,
		sid:       0,
		requestId: 0,
		data:      data,
	}
}

func (p *JSonPacket) Encode(data []byte) ([]byte, error) {
	p.data = data

	newdata, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return newdata, nil
}

func (p *JSonPacket) Decode() error {

	err := json.Unmarshal(p.data, p)
	if err != nil {
		return err
	}

	return nil
}

func (p JSonPacket) Ver() uint8 {
	return p.ver
}

func (p JSonPacket) Ccode() uint32 {
	return p.ccode
}

func (p JSonPacket) Mid() uint16 {
	return p.mid
}

func (p JSonPacket) Sid() uint16 {
	return p.sid
}

func (p JSonPacket) RequestId() uint32 {
	return p.requestId
}

func (p JSonPacket) Bytes() []byte {
	return p.data
}

func (p JSonPacket) String() string {
	return fmt.Sprintf("{ver:%d ccode:%d mid:%d sid:%d requestId:%d datalen:%d}", p.ver, p.ccode, p.mid, p.sid, p.requestId, len(p.data))
}
