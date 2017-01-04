# go-openc可以将被close的chan重新打开，这样可以重复使用已申请的内存，比起每次的创建新的chan性能提高很多。


		BenchmarkMakeCloseChan-4   	20000000	        64.5 ns/op
		--- BENCH: BenchmarkMakeCloseChan-4
		openChan_test.go:99: BenchmarkMakeCloseChan, go1.7.4, Times:         1, use:             0s         0s/op
		openChan_test.go:99: BenchmarkMakeCloseChan, go1.7.4, Times:       100, use:             0s         0s/op
		openChan_test.go:99: BenchmarkMakeCloseChan, go1.7.4, Times:     10000, use:        997.9µs       99ns/op
		openChan_test.go:99: BenchmarkMakeCloseChan, go1.7.4, Times:   1000000, use:      91.4044ms       91ns/op
		openChan_test.go:99: BenchmarkMakeCloseChan, go1.7.4, Times:  20000000, use:     1.2906058s       64ns/op
		BenchmarkCloseOpenChan-4   	50000000	        29.4 ns/op
		--- BENCH: BenchmarkCloseOpenChan-4
		openChan_test.go:115: BenchmarkCloseOpenChan, go1.7.4, Times:         1, use:             0s         0s/op
		openChan_test.go:115: BenchmarkCloseOpenChan, go1.7.4, Times:       100, use:             0s         0s/op
		openChan_test.go:115: BenchmarkCloseOpenChan, go1.7.4, Times:     10000, use:             0s         0s/op
		openChan_test.go:115: BenchmarkCloseOpenChan, go1.7.4, Times:   1000000, use:      29.5253ms       29ns/op
		openChan_test.go:115: BenchmarkCloseOpenChan, go1.7.4, Times:  50000000, use:     1.4721707s       29ns/op
		PASS
		ok  	github.com/yireyun/go-openc	3.364s
···