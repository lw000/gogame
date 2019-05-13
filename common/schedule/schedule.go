package ggschedule

import (
	"errors"
	"time"
	"tuyue/tuyue_common/auth"
	"tuyue/tuyue_common/utils"

	"github.com/ouqiang/timewheel"
)

type Schedule struct {
	tw *timewheel.TimeWheel
	f  func(data interface{})
}

func NewSchedule() *Schedule {
	return &Schedule{}
}

func (s *Schedule) AddTask(second int, data interface{}) string {
	taskId, err := tyauth.MD5([]byte(tyutils.UUID()))
	if err != nil {
		return ""
	}
	s.tw.AddTimer(time.Second*time.Duration(second), taskId, map[string]interface{}{taskId: data})
	return taskId
}

func (s *Schedule) RemoveTask(taskId interface{}) {
	if taskId == nil {
		return
	}

	s.tw.RemoveTimer(taskId)
}

func (s *Schedule) Start(f func(data interface{})) error {
	s.f = f
	s.tw = timewheel.New(time.Second*time.Duration(1), 3600, func(taskData interface{}) {
		if s.f != nil {
			s.f(taskData)
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
