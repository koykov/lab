package main

import "math"

func F(x uint64) bool {
	return x%uint64(math.MaxUint32) == 0
}
