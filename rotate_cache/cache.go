package main

import (
	"github.com/koykov/fastconv"
	"github.com/koykov/policy"
)

type entry struct {
	off, len uint
}

type cache struct {
	lock policy.Lock
	idx  map[int32]entry
	buf  []byte
}

func (c *cache) set(id int32, val string) {
	c.lock.Lock()
	if c.idx == nil {
		c.idx = make(map[int32]entry)
	}
	e := entry{
		off: uint(len(c.buf)),
		len: uint(len(val)),
	}
	c.idx[id] = e
	c.buf = append(c.buf, val...)
	c.lock.Unlock()
}

func (c *cache) get(id int32) string {
	c.lock.Lock()
	defer c.lock.Unlock()
	if e, ok := c.idx[id]; ok {
		return fastconv.B2S(c.buf[e.off : e.off+e.len])
	}
	return ""
}

func (c *cache) reset() {
	c.lock.Lock()
	for k := range c.idx {
		delete(c.idx, k)
	}
	c.buf = c.buf[:0]
	c.lock.Unlock()
}
