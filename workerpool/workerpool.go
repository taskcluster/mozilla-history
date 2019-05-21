package workerpool

import "sync"

func newWorkPipeline() *WorkPipeline {
	requestChannel := make(chan Work)
	processedChannel := make(chan Result)
	return &WorkPipeline{
		// Context for submitting work
		SubmitterContext: &SubmitterContext{
			RequestChannel: requestChannel,
			waitGroup:      sync.WaitGroup{},
		},
		// Context for executing work
		workExecutorContext: &workExecutorContext{
			RequestChannel:   requestChannel,
			ProcessedChannel: processedChannel,
			waitGroup:        sync.WaitGroup{},
		},
		// Context for listening to work
		workListenerContext: &workListenerContext{
			ProcessedChannel: processedChannel,
		},
	}
}

type WorkPipeline struct {
	SubmitterContext    *SubmitterContext
	workExecutorContext *workExecutorContext
	workListenerContext *workListenerContext
}

type Work func(workerId int) Result
type Result interface{}

type SubmitterContext struct {
	RequestChannel chan<- Work
	// WaitGroup is done when all work submitters have submitted their work
	waitGroup sync.WaitGroup
}

type workExecutorContext struct {
	RequestChannel   <-chan Work
	ProcessedChannel chan<- Result
	// WaitGroup is done when all workers have completed their work
	waitGroup sync.WaitGroup
}

type workListenerContext struct {
	ProcessedChannel <-chan Result
}

type WorkSubmitter func(wsc *SubmitterContext)

type WorkerPool struct {
	Workers  []*Worker
	pipeline *WorkPipeline
}

type Worker struct {
	context *workExecutorContext
	id      int
}

func (wp *WorkerPool) AddWork(workSubmitter WorkSubmitter) {
	wp.pipeline.SubmitterContext.waitGroup.Add(1)
	go func() {
		workSubmitter(wp.pipeline.SubmitterContext)
		wp.pipeline.SubmitterContext.waitGroup.Done()
	}()
}

func (wp *WorkerPool) Done() {
	go func() {
		wp.pipeline.SubmitterContext.waitGroup.Wait()
		close(wp.pipeline.SubmitterContext.RequestChannel)
	}()
}

func (wp *WorkerPool) OnComplete(f func(result Result)) {
	for result := range wp.pipeline.workListenerContext.ProcessedChannel {
		f(result)
	}
}

func New(capacity int) *WorkerPool {
	wp := &WorkerPool{}
	wp.pipeline = newWorkPipeline()
	wp.pipeline.workExecutorContext.waitGroup.Add(capacity)
	wp.Workers = make([]*Worker, capacity, capacity)
	for i := 0; i < capacity; i++ {
		wp.Workers[i] = &Worker{
			context: wp.pipeline.workExecutorContext,
			id:      i,
		}
		go func(i int) {
			wp.Workers[i].WorkUntilDone()
		}(i)
	}
	go func() {
		wp.pipeline.workExecutorContext.waitGroup.Wait()
		close(wp.pipeline.workExecutorContext.ProcessedChannel)
	}()
	return wp
}

func (worker *Worker) WorkUntilDone() {
	for work := range worker.context.RequestChannel {
		worker.context.ProcessedChannel <- work(worker.id)
	}
	worker.context.waitGroup.Done()
}
