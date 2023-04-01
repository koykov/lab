package generic_value_alloc

import "github.com/koykov/byteseq"

type ctx[T byteseq.Byteseq] struct {
	a I[T]
	b I[T]
}

func (ctx *ctx[T]) setA(a I[T]) {
	ctx.a = a
}

func (ctx *ctx[T]) setB(b I[T]) {
	ctx.b = b
}

func (ctx *ctx[T]) reset() {
	ctx.a = nil
	ctx.b = nil
}
