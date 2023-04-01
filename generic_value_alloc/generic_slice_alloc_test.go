package generic_value_alloc

import "testing"

func BenchmarkGVAlloc(b *testing.B) {
	b.Run("uint32/255", func(b *testing.B) {
		b.ReportAllocs()
		ctx := ctx[string]{}
		for i := 0; i < b.N; i++ {
			ctx.setA(A[string]{m: 255})
			ctx.a.Foo()
			ctx.reset()
		}
	})
	b.Run("uint32/256", func(b *testing.B) {
		b.ReportAllocs()
		ctx := ctx[string]{}
		for i := 0; i < b.N; i++ {
			ctx.setA(A[string]{m: 256})
			ctx.a.Foo()
			ctx.reset()
		}
	})
	b.Run("string/alloc", func(b *testing.B) {
		const ss = "foobar"
		b.ReportAllocs()
		ctx := ctx[string]{}
		for i := 0; i < b.N; i++ {
			ctx.setB(B[string]{s: ss})
			ctx.b.Foo()
			ctx.reset()
		}
	})
	b.Run("string/alloc-free", func(b *testing.B) {
		const ss = "foobar"
		b.ReportAllocs()
		ctx := ctx[string]{}
		x := B[string]{s: ss}
		for i := 0; i < b.N; i++ {
			ctx.setB(&x)
			ctx.b.Foo()
			ctx.reset()
		}
	})
}
