package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

var parsedUrl = make(map[string]bool)

func isDuplication(url string) bool {
	if exist, ok := parsedUrl[url]; ok && exist {
		return true
	}

	parsedUrl[url] = true
	return false
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	e.Scheduler.Run()

	out := make(chan ParseResult)
	for i := 0; i < e.WorkerCount; i++ {
		CreateWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplication(r.Url) {
			continue
		}

		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			fmt.Printf("Got item: #%d: %v\n",
				itemCount, item)
			itemCount++
		}

		for _, r := range result.Requests {
			if isDuplication(r.Url) {
				continue
			}

			e.Scheduler.Submit(r)
		}
	}
}

func CreateWorker(in chan Request,
	out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
