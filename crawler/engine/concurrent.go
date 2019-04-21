package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	SaverCount  int
	Save        Save
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Save(Item)
	ItemChan() chan Item
	Run()
}

type Save interface {
	ItemSave(Item) error
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

	for i := 0; i < e.SaverCount; i++ {
		CreateSaver(e.Scheduler.ItemChan(), e.Save)
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
			e.Scheduler.Save(item)
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

func CreateSaver(out chan Item, save Save) {
	go func() {
		for {
			item := <-out
			err := save.ItemSave(item)
			if err != nil {
				continue
			}
		}
	}()
}
