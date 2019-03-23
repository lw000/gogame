package ggtest

import (
	log "github.com/alecthomas/log4go"
	"time"
)

func TExec(info string, f func() (interface{}, error)) (interface{}, error) {
	start := time.Now()
	data, err := f()
	end := time.Now()
	log.Info("%s %v", info, end.Sub(start))
	return data, err
}
