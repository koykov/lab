package loop_opt

import "testing"

const limit = 1e3

var (
	a = make([]int32, limit)
)

func init() {
	for i := 0; i < limit; i++ {
		a[i] = int32(i)
	}
}

func BenchmarkLoopBoundsCheckOn(b *testing.B) {
	b.ResetTimer()
	var x int32
	for i := 0; i < b.N; i++ {
		for j := 0; j < limit; j++ {
			x += a[j]
		}
	}
}

func BenchmarkLoopBoundsCheckOff(b *testing.B) {
	b.ResetTimer()
	var x int32
	for i := 0; i < b.N; i++ {
		_ = a[limit-1]
		for j := 0; j < limit; j++ {
			x += a[j]
		}
	}
}
