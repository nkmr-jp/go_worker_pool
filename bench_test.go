package main

import (
	"testing"

	"github.com/nkmr-jp/go_worker_pool/pool"
	"github.com/nkmr-jp/go_worker_pool/work"
)

func BenchmarkConcurrent(b *testing.B) {
	collector := pool.StartDispatcher(WorkerCount) // start up worker pool

	for n := 0; n < b.N; n++ {
		for i, job := range work.CreateJobs(20) {
			collector.Work <- pool.Work{Job: job, ID: i}
		}
	}
}

func BenchmarkNonConcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, job := range work.CreateJobs(20) {
			work.DoWork(job, 1)
		}
	}
}
