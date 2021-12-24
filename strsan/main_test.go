package main

import (
	"sync/atomic"
	"testing"

	"github.com/koykov/fastconv"
)

func sanitizeString(s string) {
	b := fastconv.S2B(s)
	_ = b[len(b)-1]
	for i := 0; i < len(b); i++ {
		b[i] = sanitizeByte(b[i])
	}
}

func sanitizeStringRolling(s string) {
	b := fastconv.S2B(s)
	_ = b[len(b)-1]

	i, chunks := 0, len(b)/5
	for ; i < chunks; i += 5 {
		b[i] = sanitizeByte(b[i])
		b[i+1] = sanitizeByte(b[i+1])
		b[i+2] = sanitizeByte(b[i+2])
		b[i+3] = sanitizeByte(b[i+3])
		b[i+4] = sanitizeByte(b[i+4])
	}
	for ; i < len(b); i++ {
		b[i] = sanitizeByte(b[i])
	}
}

func sanitizeByte(b byte) byte {
	if b < ' ' || b == 0x7f {
		return ' '
	}
	return b
}

func makeAString(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = 'a'
	}
	return fastconv.B2S(b)
}

func BenchmarkSanitize(b *testing.B) {
	sm := makeAString(1 << 4)
	md := makeAString(1 << 5)
	lg := makeAString(1 << 6)
	xl := makeAString(1 << 7)
	strings := []string{sm, md, lg, xl}

	it := new(int64)
	nextString := func() string {
		i := atomic.AddInt64(it, 1) % int64(len(strings))
		return strings[i]
	}

	b.Run("no rolling", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sanitizeString(nextString())
		}
	})
	b.Run("rolling", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sanitizeStringRolling(nextString())
		}
	})
}
