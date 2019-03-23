package ggsys

import (
	"os"
	"os/signal"
	"syscall"
)

func RegisterOnInterrupt(f func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Kill)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		<-c
		signal.Stop(c)
		f()
		close(c)
		os.Exit(1)
	}()
}
