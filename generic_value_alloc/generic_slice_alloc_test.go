package generic_value_alloc

import "testing"

func BenchmarkGVAlloc(b *testing.B) {
	b.Run("255", func(b *testing.B) {
		b.ReportAllocs()
		ctx := ctx[string]{}
		for i := 0; i < b.N; i++ {
			ctx.setA(A[string]{m: 255})
			ctx.aa.Foo()
			ctx.reset()
		}
	})
	b.Run("256", func(b *testing.B) {
		b.ReportAllocs()
		ctx := ctx[string]{}
		for i := 0; i < b.N; i++ {
			ctx.setA(A[string]{m: 256})
			ctx.aa.Foo()
			ctx.reset()
		}
	})
}
