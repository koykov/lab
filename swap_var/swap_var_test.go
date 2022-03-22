package main

import "testing"

func BenchmarkSwapVar(b *testing.B) {
	b.Run("delta", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x, y := 5, 17
			x = x + y
			y = x - y
			x = x - y
			if x != 17 || y != 5 {
				b.FailNow()
			}
		}
	})
	b.Run("xor", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x, y := 5, 17
			x ^= y
			y ^= x
			x ^= y
			if x != 17 || y != 5 {
				b.FailNow()
			}
		}
	})
	b.Run("divmul", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x, y := 5, 17
			x = x * y
			y = x / y
			x = x / y
			if x != 17 || y != 5 {
				b.FailNow()
			}
		}
	})
	b.Run("sugar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x, y := 5, 17
			x, y = y, x
			if x != 17 || y != 5 {
				b.FailNow()
			}
		}
	})
}
