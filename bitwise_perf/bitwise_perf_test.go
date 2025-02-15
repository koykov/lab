package bitwise_perf

import "testing"

func BenchmarkBitwise(b *testing.B) {
	b.Run("", func(b *testing.B) {
		x := uint32(15)
		for i := 0; i < b.N; i++ {
			y := uint32(i % 32)
			z := uint32(1 << y)
			n := x &^ z
			_ = n
		}
	})
	b.Run("", func(b *testing.B) {
		x := uint32(15)
		for i := 0; i < b.N; i++ {
			n := x &^ 1 << uint32(i%32)
			_ = n
		}
	})
}
