package loop_opt

import (
	"fmt"
	"testing"
)

func BenchmarkGoto(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000}
	for i := 0; i < len(sizes); i++ {
		size := sizes[i]
		b.Run(fmt.Sprintf("loop/%d", size), func(b *testing.B) {
			var x int
			for j := 0; j < b.N; j++ {
				x = floop(size)
			}
			_ = x
		})
		b.Run(fmt.Sprintf("goto/%d", size), func(b *testing.B) {
			var x int
			for j := 0; j < b.N; j++ {
				x = fgoto(size)
			}
			_ = x
		})
	}
}

func floop(size int) int {
	f := 1
	for i := 1; i <= size; i++ {
		f *= i
	}
	return f
}

func fgoto(size int) int {
	i, f := 0, 1
loop:
	i++
	f *= i
	if i < size {
		goto loop
	}
	return f
}
