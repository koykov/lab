package generic_value_alloc

import "github.com/koykov/byteseq"

type ctx[T byteseq.Byteseq] struct {
	aa IA[T]
}

func (ctx *ctx[T]) setA(a IA[T]) {
	ctx.aa = a
}

func (ctx *ctx[T]) reset() {
	ctx.aa = nil
}
