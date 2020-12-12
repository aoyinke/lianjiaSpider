package main

import (
	"gengycSrc/lianjiaSpider/engine"
	"gengycSrc/lianjiaSpider/handlers"
	"gengycSrc/lianjiaSpider/persist"
	"gengycSrc/lianjiaSpider/scheduler"
)

func main() {
	itemChan, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(
		engine.Request{
			Url: "https://changzhou.lianjia.com/ershoufang/pg1/",
			Parser: engine.NewFuncParser(
				handlers.Parser,
				"Parser",
			),
		},

	)
}
