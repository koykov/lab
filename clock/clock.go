package clock

import (
	"sync/atomic"
	"time"
)

type clock struct {
	sec, nsec int64
	delta     time.Duration
	done      bool
	// once      sync.Once
}

func (c *clock) init() {
	ts := time.Now().UnixNano()
	atomic.StoreInt64(&c.sec, ts/1e9)
	atomic.StoreInt64(&c.nsec, ts%1e9)
	c.done = false
	go func() {
		t := time.NewTicker(c.delta)
		for {
			select {
			case <-t.C:
				c.tick()
				if c.done {
					return
				}
			}
		}
	}()
}

func (c *clock) Start() {
	c.init()
}

func (c *clock) Stop() {
	c.done = true
}

func (c *clock) Now() time.Time {
	return time.Unix(atomic.LoadInt64(&c.sec), atomic.LoadInt64(&c.nsec))
}

func (c *clock) tick() {
	ts := time.Now().UnixNano()
	atomic.StoreInt64(&c.sec, ts/1e9)
	atomic.StoreInt64(&c.nsec, ts%1e9)
}
