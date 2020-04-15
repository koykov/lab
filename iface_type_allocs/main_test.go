package test

import "testing"

type IObj interface {
	Get(buf *interface{})
}

type Obj struct {
	f string
}

func (o *Obj) Get(buf *interface{}) {
	*buf = o.f
}

func getIface(o IObj, t testing.TB, buf interface{}) {
	o.Get(&buf)
	if buf.(string) != "foo" {
		t.Error("mismatch")
	}
}

func getExactType(o *Obj, t testing.TB, buf interface{}) {
	o.Get(&buf)
	if buf.(string) != "foo" {
		t.Error("mismatch")
	}
}

func BenchmarkGetIface(b *testing.B) {
	var (
		o   = &Obj{"foo"}
		buf interface{}
	)
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		getIface(o, b, buf)
	}
}

func BenchmarkGetExactType(b *testing.B) {
	var (
		o   = &Obj{"foo"}
		buf interface{}
	)
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		getExactType(o, b, buf)
	}
}
