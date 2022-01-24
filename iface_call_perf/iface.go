package main

import "math"

type I interface {
	Check(x uint64) bool
}

type T struct{}

func (t T) Check(x uint64) bool {
	return x%uint64(math.MaxUint32) == 0
}

var (
	t I
)

func init() {
	t = T{}
}
