package ggsockt

import (
	"github.com/labstack/gommon/log"
	"net"
)

type Client struct {
	conn      net.Conn
	connected bool
	done      chan struct{}
	onMessage func(data []byte) error
}

func (c *Client) SetOnMessage(onMessage func(data []byte) error) {
	c.onMessage = onMessage
}

func (c *Client) Connected() bool {
	return c.connected
}

func (c *Client) Open(address string) error {
	var err error
	c.conn, err = net.Dial("tcp", "127.0.0.1:7777")

	checkError(err)

	c.connected = true

	go c.run()

	return nil
}

func (c *Client) Send(data []byte) error {
	var n int
	var err error
	n, err = c.conn.Write([]byte("Hello!"))
	if err != nil {
		log.Error("connected closed")
	}

	if n > 0 {

	}

	return nil
}

func (c *Client) run() {
	var n int
	var err error
	buf := make([]byte, 1024)
	for {
		n, err = c.conn.Read(buf)
		if err != nil {
			log.Error("connected closed")
			break
		}

		if n > 0 {

		}

		if c.onMessage != nil {
			err = c.onMessage(buf[0:n])
		}

		log.Printf("read:%s\n", string(buf[0:n]))
	}
}

func (c *Client) Close() error {
	err := c.conn.Close()
	if err != nil {

	}
	return nil
}
