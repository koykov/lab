package popcnt

import (
	"math"
	"math/bits"
)

var tpopcnt [math.MaxUint8 + 1]uint64

func init() {
	for i := 0; i <= math.MaxUint8; i++ {
		tpopcnt[i] = uint64(bits.OnesCount8(byte(i)))
	}
}

func popcntTable(p []byte) (r uint64) {
	n := len(p)
	if n == 0 {
		return
	}
	_, _ = p[n-1], tpopcnt[math.MaxUint8]
	for len(p) > 8 {
		r += tpopcnt[p[0]] + tpopcnt[p[1]] + tpopcnt[p[2]] + tpopcnt[p[3]] +
			tpopcnt[p[4]] + tpopcnt[p[5]] + tpopcnt[p[6]] + tpopcnt[p[7]]
		p = p[8:]
	}
	for i := 0; i < len(p); i++ {
		r += tpopcnt[p[i]]
	}
	return
}
