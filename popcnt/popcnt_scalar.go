package popcnt

import "math/bits"

func popcntScalar(p []byte) (r uint64) {
	n := len(p)
	if n == 0 {
		return
	}
	_ = p[n-1]
	for i := 0; i < n; i++ {
		r += uint64(bits.OnesCount8(p[i]))
	}
	return
}
