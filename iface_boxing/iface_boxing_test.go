package main

import "testing"

func BenchmarkIfaceBoxing(b *testing.B) {
	b.Run("native", func(b *testing.B) {
		var x A
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = x.Foo().
				Bar("aaa", "lorem ipsum").
				Bar("bbb", uint32(25678)).
				Bar("ccc", int64(-252435234)).
				Bar("ddd", float32(34534.1234132)).
				Push()
			x.reset()
		}
		_ = x
	})
	b.Run("dummy", func(b *testing.B) {
		var x B
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = x.Foo().
				Bar("aaa", "lorem ipsum").
				Bar("bbb", uint32(25678)).
				Bar("ccc", int64(-252435234)).
				Bar("ddd", float32(34534.1234132)).
				Push()
		}
		_ = x
	})
}
