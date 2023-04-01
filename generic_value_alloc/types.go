package generic_value_alloc

import "github.com/koykov/byteseq"

type IA[T byteseq.Byteseq] interface {
	Foo()
}

type A[T byteseq.Byteseq] struct {
	m uint32
}

func (a A[T]) Foo() {
	a.m = 5 + 6
}
