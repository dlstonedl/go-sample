package engine

import "github.com/elastic/go-elasticsearch/v7"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	Saver       Saver
	SaverCount  int
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

type Saver interface {
	Save(*elasticsearch.Client, Item)
	CreateClientPool() chan *elasticsearch.Client
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

	itemChan := make(chan Item)
	clientChan := e.Saver.CreateClientPool()
	for i := 0; i < e.SaverCount; i++ {
		CreateSaver(itemChan, clientChan, e.Saver)
	}

	for _, r := range seeds {
		if isDuplication(r.Url) {
			continue
		}

		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() { itemChan <- item }()
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

func CreateSaver(itemChan chan Item, clientChan chan *elasticsearch.Client, saver Saver) {
	go func() {
		for {
			item := <-itemChan
			client := <-clientChan
			saver.Save(client, item)
		}
	}()
}
