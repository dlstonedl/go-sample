package main

import (
	"github.com/dlstonedl/go-sample/crawler/engine"
	"github.com/dlstonedl/go-sample/crawler/persist"
	"github.com/dlstonedl/go-sample/crawler/scheduler"
	"github.com/dlstonedl/go-sample/crawler/zhenai/parser"
)

func main() {

	itemChan, err := persist.ItemSaver("crawler")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}

	//e := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.QueueScheduler{},
	//	WorkerCount: 100,
	//	SaverCount:  1,
	//	Saver: &persist.ElasticSaver{
	//		ClientCount: 1,
	//		Index:       "crawler",
	//	},
	//}

	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
