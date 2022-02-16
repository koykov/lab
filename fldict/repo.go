package main

import (
	"github.com/koykov/entry"
	"github.com/koykov/hash"
)

type repo struct {
	hsh hash.BHasher
	idx map[uint64]entry.Entry64
	buf []byte
}

func (r *repo) add(p []byte) {
	h := r.hsh.Sum64(p)
	if _, ok := r.idx[h]; ok {
		return
	}
	lo := uint32(len(r.buf))
	r.buf = append(r.buf, p...)
	hi := uint32(len(r.buf))
	var e entry.Entry64
	e.Encode(lo, hi)
	r.idx[h] = e
}

func (r repo) flush(filename string) (err error) {
	// todo implement me
	return
}

func (r *repo) reset() {
	r.buf = r.buf[:0]
	for h := range r.idx {
		delete(r.idx, h)
	}
}
