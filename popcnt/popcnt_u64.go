package popcnt

import (
	"math/bits"
	"unsafe"
)

func popcntU64(p []byte) (r uint64) {
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
		_ = buf64[n8-1]
		for i := 0; i < n8; i++ {
			r += uint64(bits.OnesCount64(buf64[i]))
		}
		p = p[n8*8:]
	}
	for i := 0; i < len(p); i++ {
		r += uint64(bits.OnesCount8(p[i]))
	}
	return
}
