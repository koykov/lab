package main

import "math"

type N struct {
	Check func(uint64) bool
}

func nCheck(x uint64) bool {
	return x%uint64(math.MaxUint32) == 0
}

var (
	n = N{Check: nCheck}
)
