package network

import (
	"demo/gogame/packet"
	"demo/gogame/user"
	log "github.com/alecthomas/log4go"
	"net"
)

type SrvHandler interface {
	Open() error
	Close() error
	OnMessage([]byte)
	OnError(errCode int, errText string)
}

type Srv struct {
	listener *net.TCPListener
	handler  SrvHandler
}

func (s *Srv) Start(address string, handler SrvHandler) error {
	var (
		er      error
		tcpAddr *net.TCPAddr
	)

	tcpAddr, er = net.ResolveTCPAddr("tcp4", address)
	if er != nil {
		log.Error(er)
	}

	s.listener, er = net.ListenTCP("tcp", tcpAddr)
	if er != nil {
		log.Error(er)
	}

	s.handler = handler

	return nil
}

func (s *Srv) run() {
	for {
		conn, er := s.listener.Accept()
		if er != nil {
			log.Error(er)
			continue
		}

		log.Info("%s", conn.RemoteAddr().String())

		go s.handleConnect(conn)
	}
}

func (s *Srv) handleConnect(conn net.Conn) {
	u := user.NewUser()
	u.AttachClient(NewClientWith(conn))
	u.SetIp(conn.RemoteAddr().String())

	user.UserMgr().AddWith(u)

	buf := make([]byte, 4096)
	buffer := packet.NewNetBuffer()
	for {
		n, er := conn.Read(buf)
		if er != nil {
			log.Info("%s, connect error:%s", conn.RemoteAddr().String(), er.Error())
			return
		}

		if n > 0 {
			buffer.Add(buf[0:n])
		}
	}
}
