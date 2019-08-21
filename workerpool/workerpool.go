// workerpool is a general purpose package for parallelising work among many
// worker go routines.
//
// The package should be typically used as follows:
//
//  // Create a worker pool, to initialise the worker go routines
//  wp := workerpool.New(numberOfWorkerGoRoutines)
//  // Add work generators to the pool; the passed in functions should generate
//  // work by writing tasks to the RequestChannel of the SubmitterContext
//  wp.AddWork(function1(context *workerpool.SubmitterContext))
//  wp.AddWork(function2(context *workerpool.SubmitterContext))
//  wp.AddWork(function3(context *workerpool.SubmitterContext))
//  // Signal that wp.AddWork(...) will not be called again
//  wp.Done()
//  // Mandatory: provide callback for task completion (e.g. do nothing / log
//  // results / generate work for another worker pool)
//  wp.OnComplete(func(result workerpool.Result) {...})
package workerpool

import "sync"

func newWorkPipeline() *workPipeline {
	requestChannel := make(chan Work)
	processedChannel := make(chan Result)
	return &workPipeline{
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

type workPipeline struct {
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
	workers  []*worker
	pipeline *workPipeline
}

type worker struct {
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
	wp.workers = make([]*worker, capacity, capacity)
	for i := 0; i < capacity; i++ {
		wp.workers[i] = &worker{
			context: wp.pipeline.workExecutorContext,
			id:      i,
		}
		go func(i int) {
			wp.workers[i].workUntilDone()
		}(i)
	}
	go func() {
		wp.pipeline.workExecutorContext.waitGroup.Wait()
		close(wp.pipeline.workExecutorContext.ProcessedChannel)
	}()
	return wp
}

func (w *worker) workUntilDone() {
	for work := range w.context.RequestChannel {
		w.context.ProcessedChannel <- work(w.id)
	}
	w.context.waitGroup.Done()
}
