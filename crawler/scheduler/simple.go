package scheduler

import "github.com/dlstonedl/go-sample/crawler/engine"

type SimpleSchedule struct {
	workerChan chan engine.Request
	itemChan   chan engine.Item
}

func (s *SimpleSchedule) Save(item engine.Item) {
	go func() { s.itemChan <- item }()
}

func (s *SimpleSchedule) ItemChan() chan engine.Item {
	return s.itemChan
}

func (s *SimpleSchedule) Submit(request engine.Request) {
	go func() { s.workerChan <- request }()
}

func (s *SimpleSchedule) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleSchedule) WorkerReady(chan engine.Request) {
}

func (s *SimpleSchedule) Run() {
	s.workerChan = make(chan engine.Request)
	s.itemChan = make(chan engine.Item)
}
