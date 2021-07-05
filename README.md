# go_worker_pool

See: [Write a Go Worker Pool in 15 minutes | by Joseph Livni | Medium](https://medium.com/@j.d.livni/write-a-go-worker-pool-in-15-minutes-c9b42f640923)

```sh
DEBUG=true go run main.go
#> 2021/07/05 12:15:09 starting application...
#> starting worker:  1
#> starting worker:  2
#> starting worker:  3
#> starting worker:  4
#> starting worker:  5
#> worker [3] - created hash [3193224551] from word [XVlBzgba]
#> worker [4] - created hash [121297580] from word [xhxKQFDa]
#> worker [2] - created hash [1481401259] from word [hTHctcuA]
#> ...
```

WorkerCount = 5
```sh
# See: https://qiita.com/marnie_ms4/items/7014563083ca1d824905
go test -bench . -benchmem
#> starting worker:  1
#> starting worker:  2
#> starting worker:  3
#> starting worker:  4
#> starting worker:  5
#> goos: darwin
#> goarch: amd64
#> pkg: github.com/nkmr-jp/go_worker_pool
#> cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
#> BenchmarkConcurrent-12                 1        3008696766 ns/op           10248 B/op        111 allocs/op
#> BenchmarkNonConcurrent-12              1        20047928833 ns/op           2048 B/op         47 allocs/op
#> PASS
#> ok      github.com/nkmr-jp/go_worker_pool       23.572s
```

WorkerCount = 10
```sh
go test -bench . -benchmem
#> starting worker:  1
#> starting worker:  2
#> starting worker:  3
#> starting worker:  4
#> starting worker:  5
#> starting worker:  6
#> starting worker:  7
#> starting worker:  8
#> starting worker:  9
#> starting worker:  10
#> goos: darwin
#> goarch: amd64
#> pkg: github.com/nkmr-jp/go_worker_pool
#> cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
#> BenchmarkConcurrent-12                 1        1000470827 ns/op           14496 B/op        138 allocs/op
#> BenchmarkNonConcurrent-12              1        20043945895 ns/op           2656 B/op         54 allocs/op
#> PASS
#> ok      github.com/nkmr-jp/go_worker_pool       21.148s
```
