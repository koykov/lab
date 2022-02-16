package main

import (
	"github.com/koykov/entry"
	"github.com/koykov/hash"
)

type repo struct {
	hsh hash.BHasher
	lng string
	idx map[uint64]entry.Entry64
	buf []byte
	out []uint64
}

func newRepo(hsh hash.BHasher) *repo {
	r := repo{
		hsh: hsh,
		idx: make(map[uint64]entry.Entry64),
	}
	return &r
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
	r.lng = ""
	r.buf = r.buf[:0]
	r.out = r.out[:0]
	for h := range r.idx {
		delete(r.idx, h)
	}
}
