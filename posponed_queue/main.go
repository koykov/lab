package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	wc   = 4
	size = 100
)

type pq struct {
	d  time.Duration
	qt chan int64
	qs chan any
}

var (
	qi *pq
	c  uint32
)

func (q *pq) enqueue(x any) {
	q.qs <- x
	t := time.Now()
	q.qt <- t.UnixNano()
	fmt.Printf("enqueue %d at %s\n", x, t.Format(time.RFC3339Nano))
}

func (q *pq) do(fn func(any) error) {
	for {
		select {
		case t := <-q.qt:
			un := time.Now().UnixNano()
			if d := q.d - (time.Duration(un) - time.Duration(t)); d > 0 {
				fmt.Printf("wait %s\n", d)
				time.Sleep(d)
			}
			x := <-q.qs
			_ = fn(x)
		}
	}
}

func init() {
	qi = &pq{
		d:  time.Second * 10,
		qt: make(chan int64, size),
		qs: make(chan any, size),
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go qi.do(func(x any) error {
			fmt.Printf("process %d at %s\n", x, time.Now().Format(time.RFC3339Nano))
			time.Sleep(time.Millisecond * 67)
			return nil
		})
	}

	for i := 0; i < 4; i++ {
		qi.enqueue(atomic.AddUint32(&c, 1))
		time.Sleep(time.Millisecond * 10)
	}
	time.Sleep(time.Second * 2)

	for i := 0; i < 10; i++ {
		qi.enqueue(atomic.AddUint32(&c, 1))
		time.Sleep(time.Millisecond * 40)
	}
	time.Sleep(time.Second * 20)

	wg.Wait()
}
