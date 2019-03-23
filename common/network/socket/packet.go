package ggsockt

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
)

type Packet struct {
	len       int
	version   int
	mainCmd   uint16
	subCmd    uint16
	requestId uint32
	data      []byte
}

func New(mainCmd, subCmd uint16, requestId uint32) *Packet {
	pk := &Packet{
		version:   1,
		mainCmd:   mainCmd,
		subCmd:    subCmd,
		requestId: requestId,
	}
	return pk
}

func NewWithData(data []byte) *Packet {
	pk := &Packet{}
	buf := bytes.NewBuffer(data)
	err := pk.readHead(buf)
	if err != nil {
		return nil
	}
	pk.data = buf.Bytes()

	return pk
}

func (p *Packet) writeHead(buf *bytes.Buffer) (err error) {
	if err = binary.Write(buf, binary.LittleEndian, p.len); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.LittleEndian, p.version); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.LittleEndian, p.mainCmd); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.LittleEndian, p.subCmd); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.LittleEndian, p.requestId); err != nil {
		return err
	}

	return err
}

func (p *Packet) readHead(buf *bytes.Buffer) (err error) {
	if err = binary.Read(buf, binary.LittleEndian, &p.len); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &p.version); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &p.mainCmd); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &p.subCmd); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.LittleEndian, &p.requestId); err != nil {
		return err
	}
	return err
}

//Encode 编码数据包
func (p *Packet) Encode(data []byte) error {
	buf := &bytes.Buffer{}
	p.len = len(data) + 4 + 4 + 2 + 2 + 4
	err := p.writeHead(buf)
	if err != nil {
		return err
	}

	if len(data) > 0 {
		var n int
		n, err = buf.Write(data)
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
		err := p.Encode([]byte{})
		if err != nil {
			return err
		}
	} else {
		data, err := proto.Marshal(pb)
		if err != nil {
			return err
		}
		err = p.Encode(data)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p Packet) Version() int {
	return p.version
}

func (p Packet) Len() int {
	return p.len
}

func (p Packet) MainCmd() uint16 {
	return p.mainCmd
}

func (p Packet) SubMain() uint16 {
	return p.subCmd
}

func (p Packet) RequestId() uint32 {
	return p.requestId
}

func (p Packet) Data() []byte {
	return p.data
}

func (p Packet) String() string {
	return fmt.Sprintf("{len:%d version:%d mainCmd:%d subCmd:%d requestId:%d datalen:%d}", p.len, p.version, p.mainCmd, p.subCmd, p.requestId, len(p.data))
}
