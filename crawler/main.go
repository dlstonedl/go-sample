package main

import (
	"github.com/dlstonedl/go-sample/crawler/engine"
	"github.com/dlstonedl/go-sample/crawler/persist"
	"github.com/dlstonedl/go-sample/crawler/scheduler"
	"github.com/dlstonedl/go-sample/crawler/zhenai/parser"
)

func main() {
	pool := persist.ClientPool{
		ClientCount: 1,
	}
	pool.Init()

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		SaverCount:  1,
		Repo: &persist.EsRepo{
			ClientPool: pool,
		},
	}

	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
