package main

import (
	"os"
	"regexp"
	"sort"

	"github.com/koykov/bytealg"
	"github.com/koykov/entry"
	"github.com/koykov/fastconv"
	"github.com/koykov/hash"
)

var (
	sdRe = regexp.MustCompile("^[0-9$]+")
)

type repo struct {
	hsh  hash.BHasher
	lng  string
	idx  map[uint64]entry.Entry64
	buf  []byte
	bufB [][]byte
	out  []uint64
}

func newRepo(hsh hash.BHasher) *repo {
	r := repo{
		hsh: hsh,
		idx: make(map[uint64]entry.Entry64),
	}
	return &r
}

func (r *repo) add(p []byte) {
	if len(p) == 0 {
		return
	}
	r.bufB = bytealg.AppendSplit(r.bufB[:0], p, bSep, -1)
	for i := 0; i < len(r.bufB); i++ {
		b := bytealg.Trim(r.bufB[i], bTrim)
		if sdRe.Match(b) {
			continue
		}
		if len(b) == 0 {
			continue
		}
		h := r.hsh.Sum64(b)
		if _, ok := r.idx[h]; ok {
			continue
		}
		lo := uint32(len(r.buf))
		r.buf = append(r.buf, b...)
		hi := uint32(len(r.buf))
		var e entry.Entry64
		e.Encode(lo, hi)
		r.idx[h] = e
	}
}

func (r repo) flush(filename string) error {
	for h := range r.idx {
		r.out = append(r.out, h)
	}
	sort.Sort(&r)
	_ = os.Remove(filename)
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	for i := 0; i < len(r.out); i++ {
		lo, hi := r.idx[r.out[i]].Decode()
		_, _ = f.Write(r.buf[lo:hi])
		_, _ = f.Write(bNl)
	}
	return f.Close()
}

func (r *repo) reset() {
	r.lng = ""
	r.buf = r.buf[:0]
	r.out = r.out[:0]
	for h := range r.idx {
		delete(r.idx, h)
	}
}

func (r repo) Len() int {
	return len(r.out)
}

func (r repo) Less(i, j int) bool {
	var lo, hi uint32
	lo, hi = r.idx[r.out[i]].Decode()
	a := fastconv.B2S(r.buf[lo:hi])
	lo, hi = r.idx[r.out[j]].Decode()
	b := fastconv.B2S(r.buf[lo:hi])
	return a < b
}

func (r *repo) Swap(i, j int) {
	r.out[i], r.out[j] = r.out[j], r.out[i]
}
