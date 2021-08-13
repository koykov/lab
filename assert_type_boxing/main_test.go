package main

import "testing"

type I interface {
	Foo()
}

type T struct {
	addr string
}

func (t T) Foo() {}

func (t T) Addr() string {
	return t.addr
}

func assertBoxing(x I) *T {
	if t, ok := interface{}(x).(*T); ok {
		return t
	}
	return nil
}

func TestAssertTypeBoxing(t *testing.T) {
	x := &T{addr: "127.0.0.1"}
	if y := assertBoxing(x); y != nil {
		if y.Addr() != "127.0.0.1" {
			t.Error("boxing assert failed")
		}
	}
}

func BenchmarkAssertTypeBoxing(b *testing.B) {
	x := &T{addr: "127.0.0.1"}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if y := assertBoxing(x); y != nil {
			if y.Addr() != "127.0.0.1" {
				b.Error("boxing assert failed")
			}
		}
	}
}
