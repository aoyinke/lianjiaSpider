package scheduler

import "gengycSrc/lianjiaSpider/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (q *QueuedScheduler) WorkerChan() chan engine.Request  {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) Submit(r engine.Request)  {
	q.requestChan <-r
}

func (q *QueuedScheduler) WorkerReady(w chan engine.Request)  {
	q.workerChan <-w
}

func (q *QueuedScheduler) Run()  {
	q.workerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for{
			var activeWorker chan engine.Request
			var activeRequest engine.Request
			if len(requestQ) > 0 && len(workerQ) >0{
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r:=<-q.requestChan:
				requestQ = append(requestQ,r)
			case w:= <-q.workerChan:
				workerQ = append(workerQ,w)
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}

		}
	}()
}