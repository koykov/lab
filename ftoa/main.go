package main

import (
	"fmt"
	"math"
)

var (
	pw2    []uint64
	digits = []byte("0123456789")
)

func init() {
	pw2 = make([]uint64, 0, 64)
	for i := 0; i < 64; i++ {
		pw2 = append(pw2, uint64(math.Pow(2, float64(i))))
	}
}

func f64toa(dst []byte, f float64) []byte {
	bits := math.Float64bits(f)
	neg := pw2[63]&bits != 0
	if neg {
		dst = append(dst, '-')
	}
	order := (bits<<1)>>53 - 1023
	mantissa := (bits << 12) >> 12
	exMantissa := mantissa | pw2[52]
	up := exMantissa >> (52 - order)
	for up > 10 {
		dst = append(dst, digits[up%10])
		up = up / 10
	}
	dst = append(dst, digits[up])
	// todo improve reverse
	// rev := dst[:]
	// if neg {
	// 	rev = dst[1:]
	// }
	// _ = rev
	// rl := len(rev)/2
	// for i := 0; i < rl; i++ {
	// 	// rev[i], rev[len(rev)-i-1] = rev[len(rev)-i-1], rev[i]
	// }

	lo := exMantissa << (order + 12)
	sh := uint64(0)
	for i := 0; i < 64; i++ {
		if lo&pw2[i] != 0 {
			sh = uint64(i)
			break
		}
	}
	lo = lo >> sh
	if lo > 0 {
		dst = append(dst, '.')
		flo := float64(lo) / float64(pw2[lo])
		for flo > 0 {
			flo = flo * 10
			dst = append(dst, digits[int(flo)])
			flo = flo - float64(int(flo))
		}
	}
	return dst
}

func main() {
	buf := make([]byte, 0, 24)
	buf = f64toa(buf, 446.15625)
	fmt.Println(string(buf))
}

// origin 446.15625
// S P           M
// 0 10000000111 1011111000101000000000000000000000000000000000000000
// 1 10000000111 1011111000101000000000000000000000000000000000000000
