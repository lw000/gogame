package ggsockt

import (
	"github.com/labstack/gommon/log"
	"os"
)

type Context interface {
}

type context struct {
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}
}

func (c *context) new() error {
	return nil
}
