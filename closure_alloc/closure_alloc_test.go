package closure_alloc

import "testing"

type ctx struct {
	cls []func()
}

func (ctx *ctx) add(fn func()) {
	ctx.cls = append(ctx.cls, fn)
}

func (ctx *ctx) bulkExec() {
	for i := 0; i < len(ctx.cls); i++ {
		ctx.cls[i]()
	}
}

func (ctx *ctx) reset() {
	ctx.cls = ctx.cls[:0]
}

func BenchmarkClosureAlloc(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		b.ReportAllocs()
		ctx := ctx{}
		for i := 0; i < b.N; i++ {
			var c int
			ctx.add(func() { c++ })
			ctx.bulkExec()
			if c != 1 {
				b.FailNow()
			}
		}
	})
	b.Run("simple 2", func(b *testing.B) {
		b.ReportAllocs()
		ctx := ctx{}
		for i := 0; i < b.N; i++ {
			var c, d int
			ctx.add(func() {
				c++
				d++
			})
			ctx.bulkExec()
			if c != 1 {
				b.FailNow()
			}
		}
	})
	b.Run("prealloc external var", func(b *testing.B) {
		b.ReportAllocs()
		ctx := ctx{}
		var c int
		for i := 0; i < b.N; i++ {
			ctx.reset()
			c = 0
			ctx.add(func() { c++ })
			ctx.bulkExec()
			if c != 1 {
				b.FailNow()
			}
		}
	})
}
