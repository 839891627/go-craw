package main

import (
	"crawler/douban/parser"
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
)

func main() {
	doubanUrl := "https://movie.douban.com/chart"
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		//Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
		ItemChan: persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        doubanUrl,
		ParserFunc: parser.ParseMovieList,
		//Url:        "https://movie.douban.com/subject/26897885/",
		//ParserFunc: parser.ParseMovie,
	})
}
