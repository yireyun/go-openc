# go-openc可以将被close的chan重新打开，这样可以重复使用已申请的内存，比起每次的创建新的chan性能提高很多。


		BenchmarkMakeCloseChan-4       	20000000	        90.9 ns/op
		--- BENCH: BenchmarkMakeCloseChan-4
		openChan_test.go:99: go1.7.4, Times:         1, use:             0s         0s/op
		openChan_test.go:99: go1.7.4, Times:       100, use:             0s         0s/op
		openChan_test.go:99: go1.7.4, Times:     10000, use:        998.3µs       99ns/op
		openChan_test.go:99: go1.7.4, Times:   1000000, use:      91.0664ms       91ns/op
		openChan_test.go:99: go1.7.4, Times:  20000000, use:     1.8185449s       90ns/op
		BenchmarkCloseOpenChan-4       	50000000	        29.5 ns/op
		--- BENCH: BenchmarkCloseOpenChan-4
		openChan_test.go:115: go1.7.4, Times:         1, use:             0s         0s/op
		openChan_test.go:115: go1.7.4, Times:       100, use:             0s         0s/op
		openChan_test.go:115: go1.7.4, Times:     10000, use:             0s         0s/op
		openChan_test.go:115: go1.7.4, Times:   1000000, use:      28.5199ms       28ns/op
		openChan_test.go:115: go1.7.4, Times:  50000000, use:      1.473497s       29ns/op
		BenchmarkCloseOpenChanSync-4   	30000000	        44.9 ns/op
		--- BENCH: BenchmarkCloseOpenChanSync-4
		openChan_test.go:129: go1.7.4, Times:         1, use:             0s         0s/op
		openChan_test.go:129: go1.7.4, Times:       100, use:             0s         0s/op
		openChan_test.go:129: go1.7.4, Times:     10000, use:        498.2µs       49ns/op
		openChan_test.go:129: go1.7.4, Times:   1000000, use:      47.5278ms       47ns/op
		openChan_test.go:129: go1.7.4, Times:  30000000, use:     1.3479236s       44ns/op
		PASS
		ok  	github.com/yireyun/go-openc	8.235s
