package main

import "testing"

func BenchmarkCallPerf(b *testing.B) {
	b.Run("interface", func(b *testing.B) {
		var ok bool
		for i := 0; i < b.N; i++ {
			ok = t.Check(999999999)
		}
		_ = ok
	})
	b.Run("nested", func(b *testing.B) {
		var ok bool
		for i := 0; i < b.N; i++ {
			ok = n.Check(999999999)
		}
		_ = ok
	})
	b.Run("func", func(b *testing.B) {
		var ok bool
		for i := 0; i < b.N; i++ {
			ok = F(999999999)
		}
		_ = ok
	})
}
