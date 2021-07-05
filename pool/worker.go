package pool

import (
	"log"

	"github.com/nkmr-jp/go_worker_pool/work"
)

type Work struct {
	ID  int
	Job string
}

type Worker struct {
	ID            int
	WorkerChannel chan chan Work
	Channel       chan Work
	End           chan bool
}

// Start start worker
func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerChannel <- w.Channel
			select {
			case job := <-w.Channel:
				// do work
				work.DoWork(job.Job, w.ID)
			case <-w.End:
				return
			}
		}
	}()
}

// Stop end worker
func (w *Worker) Stop() {
	log.Printf("worker [%d] is stopping", w.ID)
	w.End <- true
}
