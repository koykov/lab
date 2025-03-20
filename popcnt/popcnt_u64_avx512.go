package popcnt

import (
	"math/bits"
	"unsafe"
)

func popcntU64AVX512s(data []uint64) uint64

func popcntU64AVX512(p []byte) (r uint64) {
	n := len(p)
	if n == 0 {
		return
	}
	_ = p[n-1]
	if n > 8 {
		type sh struct {
			p    uintptr
			l, c int
		}
		n8 := n / 8
		h := sh{p: uintptr(unsafe.Pointer(&p[0])), l: n8, c: n8}
		buf64 := *(*[]uint64)(unsafe.Pointer(&h))
		r += popcntU64AVX512s(buf64)
		p = p[n8*8:]
	}
	for i := 0; i < len(p); i++ {
		r += uint64(bits.OnesCount8(p[i]))
	}
	return
}
