package worker

import (
	"runtime/debug"
	"sync"

	"github.com/mu-box/microbox-server/config"
)

// structs
type (

	//
	Worker struct {
		sync.WaitGroup

		Blocking   bool
		Concurrent bool
		doTex      sync.Mutex
		queueTex   sync.Mutex
		queue      []Job
	}
)

// interfaces
type (

	//
	Job interface {
		Process()
	}
)

func New() *Worker {
	return &Worker{
		Blocking:   false,
		Concurrent: false,
		doTex:      sync.Mutex{},
		queueTex:   sync.Mutex{},
		queue:      []Job{},
	}
}

func (w *Worker) Count() int {
	return len(w.queue)
}

func (w *Worker) Queue(job Job) {
	w.queueTex.Lock()
	defer w.queueTex.Unlock()

	w.queue = append(w.queue, job)
}

func (w *Worker) QueueAndProcess(job Job) {
	w.Queue(job)
	w.Process()
}

func (w *Worker) Process() {
	if w.Blocking {
		w.execute()
	} else {
		go w.execute()
	}
}

// private

func (w *Worker) execute() {

	w.doTex.Lock()
	defer w.doTex.Unlock()

	//
	for {
		if job, ok := w.nextJob(); ok {
			w.Add(1)
			if w.Concurrent {
				go w.processJob(job)
			} else {
				w.processJob(job)
			}
		} else {
			break
		}
	}

	w.Wait()
}

func (w *Worker) nextJob() (Job, bool) {
	w.queueTex.Lock()
	defer w.queueTex.Unlock()

	var job Job

	if len(w.queue) >= 1 {
		job, w.queue = w.queue[0], w.queue[1:]
		return job, true
	}

	return nil, false
}

func (w *Worker) processJob(job Job) {
	defer w.Done()
	//
	defer func() {
		if err := recover(); err != nil {
			config.Log.Error("%s: %s", err, debug.Stack())
			config.Log.Error("[MICROBOX :: WORKER] Job failed: %+v\n", err)
		}
	}()

	//
	job.Process()
}
