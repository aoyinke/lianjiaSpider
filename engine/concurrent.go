package engine

var visitedUrl = make(map[string]bool)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
	RequestProcessor Processor

}

type Processor func(Request) (ParseResults,error)

type Scheduler interface {
	ReadNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request)  {

	out := make(chan ParseResults)
	e.Scheduler.Run()
	for i:= 0;i<e.WorkerCount;i++{
		e.createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)
	}

	// 去重
	for _,r := range seeds{
		if isDuplicate(r.Url){
			continue
		}
		// do something
		e.Scheduler.Submit(r)
	}

	for{
		result :=<- out
		for _,item := range result.Items{
			go func() {
				e.ItemChan <- item
			}()
		}
		for _,request := range result.Requests{
			if isDuplicate(request.Url){
				continue
			}
			// do something
			e.Scheduler.Submit(request)
		}


	}

}

func (e *ConcurrentEngine) createWorker(in chan Request,out chan ParseResults,ready ReadNotifier)  {
	go func() {
		for  {
			ready.WorkerReady(in)
			request := <-in

			result,e := Worker(request)
			if e!=nil{
				continue
			}
			out <-result
		}
	}()
}

func isDuplicate(url string) bool  {
	if visitedUrl[url]{
		return true
	}
	visitedUrl[url] = true
	return false
}