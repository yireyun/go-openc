// openChan.go
package openc

import (
	"reflect"
	"unsafe"
)

type sudog struct {
	// The following fields are protected by the hchan.lock of the
	// channel this sudog is blocking on. shrinkstack depends on
	// this.

	g          *unsafe.Pointer
	selectdone *uint32 // CAS to 1 to win select race (may point to stack)
	next       *sudog
	prev       *sudog
	elem       unsafe.Pointer // data element (may point to stack)

	// The following fields are never accessed concurrently.
	// waitlink is only accessed by g.

	releasetime int64
	ticket      uint32
	waitlink    *sudog // g.waiting list
	c           *hchan // channel
}

type waitq struct {
	first *sudog
	last  *sudog
}

type mutex struct {
	// Futex-based impl treats it as uint32 key,
	// while sema-based impl as M* waitm.
	// Used to be a union, but unions break precise GC.
	key uintptr
}

type hchan struct {
	qcount   uint    // total data in the queue
	dataqsiz uint    // size of the circular queue
	buf      uintptr // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	elemtype *int  // element type
	sendx    uint  // send index
	recvx    uint  // receive index
	recvq    waitq // list of recv waiters
	sendq    waitq // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}

func Open(c interface{}) {
	if c == nil {
		return
	}
	if v := reflect.ValueOf(c); v.Kind() == reflect.Chan {
		c := (*hchan)(unsafe.Pointer(v.Pointer()))
		c.closed = 0
	}
}

func init() {
	c := make(chan struct{})
	p := (*unsafe.Pointer)(unsafe.Pointer(&c))
	h := (*hchan)(*p)
	if h.qcount != 0 ||
		h.dataqsiz != 0 ||
		h.elemsize != 0 ||
		h.closed != 0 ||
		*h.elemtype != 0 ||
		h.sendx != 0 ||
		h.recvx != 0 ||
		h.recvq.first != nil ||
		h.recvq.last != nil ||
		h.sendq.first != nil ||
		h.sendq.last != nil ||
		h.lock.key != 0 {
		panic("After checks found the Go kernel is changed, pleace check openc and update")
	}
	close(c)
	if h.qcount != 0 ||
		h.dataqsiz != 0 ||
		h.elemsize != 0 ||
		h.closed != 1 ||
		*h.elemtype != 0 ||
		h.sendx != 0 ||
		h.recvx != 0 ||
		h.recvq.first != nil ||
		h.recvq.last != nil ||
		h.sendq.first != nil ||
		h.sendq.last != nil ||
		h.lock.key != 0 {
		panic("After checks found the Go kernel is changed, pleace check openc and update")
	}
}
