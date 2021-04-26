package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	url := "http://127.0.0.1:8080/mock/www.zhenai.com/zhenghun"

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 50,
	}

	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
}
