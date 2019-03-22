package network

import (
	"fmt"
	log "github.com/alecthomas/log4go"
	"net"
	"sync"
)

type ClientHandler interface {
	Open() error
	Close() error
	OnMessage([]byte)
	OnError(errCode int, errText string)
}

type Client struct {
	addr string
	port int
	stop bool
	conn net.Conn
	m    sync.Mutex
}

func NewClient() *Client {
	return &Client{}
}

func NewClientWith(conn net.Conn) *Client {
	return &Client{conn: conn}
}

func (c *Client) Connect(addr string, port int) error {
	address := fmt.Sprintf("%s:%d", c.addr, c.port)
	var (
		er      error
		tcpAddr *net.TCPAddr
	)
	tcpAddr, er = net.ResolveTCPAddr("tcp4", address)
	if er != nil {
		log.Error(er)
		return er
	}

	c.conn, er = net.DialTCP("tcp", nil, tcpAddr)
	if er != nil {
		log.Error(er)
		return er
	}

	go c.run()

	return nil
}

func (c *Client) Send(data []byte) (int, error) {
	c.m.Lock()
	defer c.m.Unlock()
	n, er := c.conn.Write(data)
	return n, er
}

func (c *Client) run() {
	defer func() {
		er := c.conn.Close()
		if er != nil {
			log.Error(er)
		}
	}()

	buf := make([]byte, 4096)
	for !c.stop {
		n, er := c.conn.Read(buf)
		if er != nil {
			log.Error(er)
			c.stop = true
			return
		}

		if n > 0 {
			log.Info(string(buf))
		}
	}
}
