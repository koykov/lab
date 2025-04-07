package memclr

import "unsafe"

func memclr(p []byte) {
	if len(p) == 0 {
		return
	}
	_ = p[len(p)-1]
	for i := 0; i < len(p); i++ {
		p[i] = 0
	}
}

func memclrBlock(p []byte, memclr64fn func([]uint64)) {
	n := len(p)
	if n == 0 {
		return
	}
	if n >= 32 {
		n64 := (n - n%32) / 8
		type sh struct {
			p    uintptr
			l, c int
		}
		h := sh{p: uintptr(unsafe.Pointer(&p[0])), l: n64, c: n64}
		p64 := *(*[]uint64)(unsafe.Pointer(&h))
		memclr64fn(p64)
		n = n - n%32
	}
	if n == 0 {
		return
	}
	_ = p[n-1]
	for i := 0; i < len(p); i += 8 {
		p[i] = 0
	}
}

func memclr64generic(p []uint64) {
	_ = p[len(p)-1]
	for i := 0; i < len(p); i++ {
		p[i] = 0
	}
}
