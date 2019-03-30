package ggschedule

import (
	"demo/gogame/common/auth"
	"demo/gogame/common/utilty"
	"errors"
	"github.com/ouqiang/timewheel"
	"time"
)

type Schedule struct {
	tw *timewheel.TimeWheel
	f  func(taskId, data interface{})
}

func NewSchedule() *Schedule {
	return &Schedule{}
}

func (s *Schedule) AddTask(second int, data interface{}) interface{} {
	taskId, err := ggauth.MD5([]byte(ggutilty.UUID()))
	if err != nil {
		return nil
	}
	s.tw.AddTimer(time.Second*time.Duration(second), taskId, timewheel.TaskData{"t": taskId, "d": data})
	return taskId
}

func (s *Schedule) RemoveTask(taskId interface{}) {
	if taskId == nil {
		return
	}

	s.tw.RemoveTimer(taskId)
}

func (s *Schedule) Start(f func(taskId, data interface{})) error {
	s.f = f
	s.tw = timewheel.New(time.Second*time.Duration(1), 3600, func(taskData timewheel.TaskData) {
		taskId := taskData["t"]
		data := taskData["d"]
		if s.f != nil {
			s.f(taskId, data)
		}
	})

	if s.tw == nil {
		return errors.New("start schedule failed")
	}

	s.tw.Start()

	return nil
}

func (s *Schedule) Stop() {
	s.tw.Stop()
}
