package main

import "github.com/koykov/fastconv"

type entry struct {
	off, len uint
}

type cache struct {
	idx map[int32]entry
	buf []byte
}

func (c *cache) set(id int32, val string) {
	if c.idx == nil {
		c.idx = make(map[int32]entry)
	}
	e := entry{
		off: uint(len(c.buf)),
		len: uint(len(val)),
	}
	c.idx[id] = e
	c.buf = append(c.buf, val...)
}

func (c *cache) get(id int32) string {
	if e, ok := c.idx[id]; ok {
		return fastconv.B2S(c.buf[e.off : e.off+e.len])
	}
	return ""
}

func (c *cache) reset() {
	for k := range c.idx {
		delete(c.idx, k)
	}
	c.buf = c.buf[:0]
}
