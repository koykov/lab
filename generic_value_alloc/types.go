package generic_value_alloc

import "github.com/koykov/byteseq"

type I[T byteseq.Byteseq] interface {
	Foo()
}

type A[T byteseq.Byteseq] struct {
	m uint32
}

func (a A[T]) Foo() {
	a.m = 5 + 6
}

type B[T byteseq.Byteseq] struct {
	s string
}

func (b B[T]) Foo() {
	_ = b.s
}
