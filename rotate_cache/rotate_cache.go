package main

import "sync/atomic"

type rotateCache struct {
	idx uint32
	buf [2]cache
}

func (c *rotateCache) set(id int32, val string) {
	var i uint32
	if atomic.LoadUint32(&c.idx) == 0 {
		i = 1
	}
	c.buf[i].set(id, val)
}

func (c *rotateCache) get(id int32) string {
	return c.buf[atomic.LoadUint32(&c.idx)].get(id)
}

func (c *rotateCache) rotate() {
	switch atomic.LoadUint32(&c.idx) {
	case 0:
		atomic.StoreUint32(&c.idx, 1)
	default:
		atomic.StoreUint32(&c.idx, 0)
	}
}

func (c *rotateCache) resetBuf() {
	var i uint32
	if atomic.LoadUint32(&c.idx) == 0 {
		i = 1
	}
	c.buf[i].reset()
}
