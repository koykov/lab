package test

import (
	"math"
	"testing"
)

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var x, y float32
		x, y = math.Pi, 0.385
		for j := 0; j < 1e6; j++ {
			_ = x * y
		}
	}
}

func BenchmarkGoto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var (
			x, y float32
			j    int
		)
		x, y = math.Pi, 0.385
	loop:
		_ = x * y
		j++
		if j < 1e6 {
			goto loop
		}
	}
}
