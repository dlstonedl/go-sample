package scheduler

import "github.com/dlstonedl/go-sample/crawler/engine"

type SimpleSchedule struct {
	workChan chan engine.Request
}

func (s *SimpleSchedule) ConfigureWorkerChan(c chan engine.Request) {
	s.workChan = c
}

func (s *SimpleSchedule) Submit(request engine.Request) {
	go func() { s.workChan <- request }()
}
