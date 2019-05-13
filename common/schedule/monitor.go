package ggschedule

import (
	"errors"
	log "github.com/alecthomas/log4go"
	"sync"
	"time"
)

type Monitor struct {
	done          chan struct{}
	once          sync.Once
	maxRetryCount int64
	second        int64
	onEvent       func(currentRetryCount int64, maxRetryCount int64) (ok bool)
}

func NewMonitor() *Monitor {
	return &Monitor{
		done: make(chan struct{}, 1),
	}
}

func (w *Monitor) Start(second int64, maxRetryCount int64, fn func(currentCount int64, maxRetryCount int64) (ok bool)) error {
	if second <= 0 {
		return errors.New("间隔时间错误")
	}

	w.second = second
	w.maxRetryCount = maxRetryCount
	w.onEvent = fn

	w.once.Do(func() {
		go w.run()
	})

	return nil
}

func (w *Monitor) run() {
	if w == nil {
		return
	}

	defer func() {
		log.Info("monitor server exit")
	}()

	ticker := time.NewTicker(time.Second * time.Duration(w.second))

	var currentCount int64 = 1

loop:
	for {
		select {
		case <-ticker.C:
			if currentCount >= w.maxRetryCount {
				currentCount = 0
			}

			if w.onEvent(currentCount, w.maxRetryCount) {
				currentCount = 0
			} else {
				currentCount++
			}

		case <-w.done:
			ticker.Stop()
			close(w.done)
			break loop
		}
	}
}

func (w *Monitor) Stop() error {
	if w == nil {
		return errors.New("object instance is empty")
	}

	w.done <- struct{}{}

	return nil
}
