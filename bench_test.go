package main

import "testing"
import "tutorials/concurrent-limiter/pool"
import "tutorials/concurrent-limiter/work"

func BenchmarkConcurrent(b *testing.B) {
	collector := pool.StartDispatcher(WorkerCount) // start up worker pool

	for n := 0; n < b.N; n++ {
		for i, job := range work.CreateJobs(20) {
			collector.Work <- pool.Work{Job: job, ID: i}
		}
	}
}

func BenchmarkNonconcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, work := range work.CreateJobs(20) {
			job.DoWork(work, 1)
		}
	}
}
