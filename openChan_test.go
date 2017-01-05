// openChan_test.go
package openc

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
	"unsafe"
)

var (
	c1 chan struct{}
	c2 chan int
	c3 chan interface{}
)

func TestCloseOpen(t *testing.T) {

	var wg sync.WaitGroup
	c1 = make(chan struct{}, 10)
	c2 = make(chan int, 10)
	c3 = make(chan interface{}, 10)
	c1 <- struct{}{}
	<-c1
	c2 <- 1
	<-c2
	c3 <- 1
	<-c3
	close(c1)
	close(c2)
	close(c3)

	p := (*unsafe.Pointer)(unsafe.Pointer(&c1))
	c := (*hchan)(*p)
	t.Logf("CloseChan,%+v\n", c)
	c.closed = 0
	t.Logf("InitChan, %+v\n", c)
	c1 <- struct{}{}
	t.Logf("SendChan, %+v\n", c)
	<-c1
	t.Logf("RecvChan, %+v\n", c)
	wg.Add(1)
	go func() { d, ok := <-c1; t.Logf("RecvChan, %+v, %v, %v\n", c, d, ok); wg.Done() }()
	t.Logf("StatChan, %+v\n", c)
	c1 <- struct{}{}
	t.Logf("SendChan, %+v\n", c)
	wg.Wait()
	fmt.Println()

	Open(c2)
	t.Logf("CloseChan,%+v\n", c)
	c.closed = 0
	t.Logf("InitChan, %+v\n", c)
	c2 <- 1
	t.Logf("SendChan, %+v, %v\n", c, 1)
	d := <-c2
	t.Logf("RecvChan, %+v, %v\n", c, d)
	wg.Add(1)
	go func() {
		d, ok := <-c2
		t.Logf("RecvChan, %+v, %v, %v\n", c, d, ok)
		if ok {
			wg.Done()
		}
	}()
	go func() {
		d, ok := <-c2
		t.Logf("RecvChan, %+v, %v, %v\n", c, d, ok)
		if ok {
			wg.Done()
		}
	}()
	go func() {
		d, ok := <-c2
		t.Logf("RecvChan, %+v, %v, %v\n", c, d, ok)
		if ok {
			wg.Done()
		}
	}()
	t.Logf("StatChan, %+v\n", c)
	c2 <- 2
	t.Logf("SendChan, %+v, %v\n", c, 2)
	wg.Wait()
	close(c2)
	t.Logf("StatChan, %+v\n", c)
}

func BenchmarkMakeCloseChan(b *testing.B) {
	start := time.Now()
	for i := 0; i < b.N; i++ {
		c1 = make(chan struct{}, 0)
		close(c1)
	}
	end := time.Now()
	use := end.Sub(start)
	op := use / time.Duration(b.N)
	b.Logf("%v, Times:%10v, use: %14v %10v/op\n", runtime.Version(), b.N, use, op)
}

func BenchmarkCloseOpenChan(b *testing.B) {
	c2 = make(chan int, 1)
	p := (*unsafe.Pointer)(unsafe.Pointer(&c2))
	c := (*hchan)(*p)
	start := time.Now()
	for i := 0; i < b.N; i++ {
		close(c2)
		c.closed = 0
	}
	close(c2)
	end := time.Now()
	use := end.Sub(start)
	op := use / time.Duration(b.N)
	b.Logf("%v, Times:%10v, use: %14v %10v/op\n", runtime.Version(), b.N, use, op)
}

func BenchmarkCloseOpenChanSync(b *testing.B) {
	c2 = make(chan int, 1)
	start := time.Now()
	for i := 0; i < b.N; i++ {
		close(c2)
		Open(c2)
	}
	close(c2)
	end := time.Now()
	use := end.Sub(start)
	op := use / time.Duration(b.N)
	b.Logf("%v, Times:%10v, use: %14v %10v/op\n", runtime.Version(), b.N, use, op)
}
