package main

import (
	"log"

	"github.com/nkmr-jp/go_worker_pool/pool"
	"github.com/nkmr-jp/go_worker_pool/work"
)

const WorkerCount = 5
const JobCount = 100

func main() {
	log.Println("starting application...")
	collector := pool.StartDispatcher(WorkerCount) // start up worker pool

	for i, job := range work.CreateJobs(JobCount) {
		collector.Work <- pool.Work{Job: job, ID: i}
	}
}
