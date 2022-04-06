package test

import "testing"

type Doer interface {
	Do(int32) int32
}

type base struct{}

type T1 struct {
	base
}
type T2 struct {
	base
	Doer
}

func (t T1) Do(x int32) int32 { return x }

func BenchmarkNested(b *testing.B) {
	b.Run("duck", func(b *testing.B) {
		var x Doer
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			x = &T1{}
		}
		_ = x
	})
	b.Run("nested", func(b *testing.B) {
		var x Doer
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			x = &T2{}
		}
		_ = x
	})
}
