package scheduler

import "github.com/dlstonedl/go-sample/crawler/engine"

type SimpleSchedule struct {
	workChan chan engine.Request
}

func (s *SimpleSchedule) Submit(request engine.Request) {
	go func() { s.workChan <- request }()
}

func (s *SimpleSchedule) WorkerChan() chan engine.Request {
	return s.workChan
}

func (s *SimpleSchedule) WorkerReady(chan engine.Request) {
}

func (s *SimpleSchedule) Run() {
	s.workChan = make(chan engine.Request)
}
