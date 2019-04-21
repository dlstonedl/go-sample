package main

import (
	"github.com/dlstonedl/go-sample/crawler/engine"
	"github.com/dlstonedl/go-sample/crawler/persist"
	"github.com/dlstonedl/go-sample/crawler/scheduler"
	"github.com/dlstonedl/go-sample/crawler/zhenai/parser"
)

func main() {
	single := persist.SingleClient{}
	single.Init()

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		Saver: &persist.SingleSaver{
			Index:    "crawler",
			EsClient: single.GetSingleClient,
		},
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
