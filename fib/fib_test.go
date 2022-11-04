package main

import "testing"

func BenchmarkFib(b *testing.B) {
	b.Run("fib classic", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fib(30)
		}
	})
	b.Run("fib memoized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fibm(30)
		}
	})
}
