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
	d time.Duration
	q chan item
	e chan struct{}
}

type item struct {
	ts int64
	pl any
}

var (
	qi *pq
	c  uint32
)

func newq() *pq {
	q := &pq{
		d: time.Second * 10,
		q: make(chan item, size),
		e: make(chan struct{}, size),
	}
	return q
}

func (q *pq) enqueue(x any) {
	t := time.Now()
	q.q <- item{
		ts: t.UnixNano(),
		pl: x,
	}
	fmt.Printf("enqueue %d at %s\n", x, t.Format(time.RFC3339Nano))
}

func (q *pq) do(wg *sync.WaitGroup, fn func(any) error) {
	defer wg.Done()
	for {
		select {
		case itm := <-q.q:
			un := time.Now().UnixNano()
			if d := q.d - (time.Duration(un) - time.Duration(itm.ts)); d > 0 {
				fmt.Printf("wait %s\n", d)
				time.Sleep(d)
			}
			_ = fn(itm.pl)
		case <-q.e:
			return
		}
	}
}

func (q *pq) close() {
	close(q.q)
	close(q.e)
}

func init() {
	qi = newq()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go qi.do(&wg, func(x any) error {
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

	for i := 0; i < wc; i++ {
		qi.e <- struct{}{}
	}
	wg.Wait()

	qi.close()
}
