package loop_opt

import "testing"

func BenchmarkArrayPtr(b *testing.B) {
	type X struct {
		a, b, c, d, e, f, g, h uint64
		i, j, k, l, m, n, o, p float64
	}

	b.Run("value/10", func(b *testing.B) {
		var stage [10]X
		for i := 0; i < b.N; i++ {
			var x uint64
			for _, v := range stage {
				x += v.a
			}
			_ = x
		}
	})
	b.Run("ptr/10", func(b *testing.B) {
		var stage [10]X
		for i := 0; i < b.N; i++ {
			var x uint64
			for _, v := range &stage {
				x += v.a
			}
			_ = x
		}
	})

	b.Run("value/100", func(b *testing.B) {
		var stage [100]X
		for i := 0; i < b.N; i++ {
			var x uint64
			for _, v := range stage {
				x += v.a
			}
			_ = x
		}
	})
	b.Run("ptr/100", func(b *testing.B) {
		var stage [100]X
		for i := 0; i < b.N; i++ {
			var x uint64
			for _, v := range &stage {
				x += v.a
			}
			_ = x
		}
	})

	b.Run("value/1000", func(b *testing.B) {
		var stage [1000]X
		for i := 0; i < b.N; i++ {
			var x uint64
			for _, v := range stage {
				x += v.a
			}
			_ = x
		}
	})
	b.Run("ptr/1000", func(b *testing.B) {
		var stage [1000]X
		for i := 0; i < b.N; i++ {
			var x uint64
			for _, v := range &stage {
				x += v.a
			}
			_ = x
		}
	})

	b.Run("value/10000", func(b *testing.B) {
		var stage [10000]X
		for i := 0; i < b.N; i++ {
			var x uint64
			for _, v := range stage {
				x += v.a
			}
			_ = x
		}
	})
	b.Run("ptr/10000", func(b *testing.B) {
		var stage [10000]X
		for i := 0; i < b.N; i++ {
			var x uint64
			for _, v := range &stage {
				x += v.a
			}
			_ = x
		}
	})

	b.Run("value/100000", func(b *testing.B) {
		var stage [100000]X
		for i := 0; i < b.N; i++ {
			var x uint64
			for _, v := range stage {
				x += v.a
			}
			_ = x
		}
	})
	b.Run("ptr/100000", func(b *testing.B) {
		var stage [100000]X
		for i := 0; i < b.N; i++ {
			var x uint64
			for _, v := range &stage {
				x += v.a
			}
			_ = x
		}
	})

	b.Run("value/1000000", func(b *testing.B) {
		var stage [1000000]X
		for i := 0; i < b.N; i++ {
			var x uint64
			for _, v := range stage {
				x += v.a
			}
			_ = x
		}
	})
	b.Run("ptr/1000000", func(b *testing.B) {
		var stage [1000000]X
		for i := 0; i < b.N; i++ {
			var x uint64
			for _, v := range &stage {
				x += v.a
			}
			_ = x
		}
	})
}
