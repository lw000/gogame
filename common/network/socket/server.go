package ggsockt

import (
	"github.com/labstack/gommon/log"
	"net"
	"time"
)

type Server struct {
}

func handleClient(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {

		}
	}()

	var (
		n   int
		err error
		buf []byte
	)

	buf = make([]byte, 1024)

	for {
		n, err = conn.Read(buf)
		if err != nil {
			log.Error("connected closed")
			break
		}
		if n > 0 {

		}

		log.Debug("read:%s", string(buf[0:n]))

		n, err = conn.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
		if err != nil {
			log.Error("connected closed")
			break
		}

		if n > 0 {

		}
	}
}

func (s *Server) Run(addr string) error {
	listen, err := net.Listen("tcp", addr)
	checkError(err)

	defer func() {
		err = listen.Close()
		if err != nil {

		}
	}()

	log.Debug("server start... port:[%s]", addr)

	for {
		var conn net.Conn
		conn, err = listen.Accept()
		if err != nil {
			continue
		}

		log.Debug("[%s]", conn.RemoteAddr().String())

		go handleClient(conn)
	}
}
