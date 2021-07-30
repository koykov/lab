package main

import "testing"

var a interface{}

func BenchmarkTNoPtr(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := T{d0: 0, d1: 1, d2: 2, d3: 3, d4: 4, d5: 5, d6: 6, d7: 7}
		x := w0(t)
		_ = x
	}
}

func BenchmarkTWithPtr(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := T{d0: 0, d1: 1, d2: 2, d3: 3, d4: 4, d5: 5, d6: 6, d7: 7}
		x := w1(&t)
		_ = x
	}
}

func BenchmarkTWithPtrIfaceGlobal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := T{d0: 0, d1: 1, d2: 2, d3: 3, d4: 4, d5: 5, d6: 6, d7: 7}
		a = &t
		x := w1Iface(a)
		_ = x
	}
}

func BenchmarkTWithPtrIfaceLocal(b *testing.B) {
	b.ReportAllocs()
	var a1 interface{}
	for i := 0; i < b.N; i++ {
		t := T{d0: 0, d1: 1, d2: 2, d3: 3, d4: 4, d5: 5, d6: 6, d7: 7}
		a1 = &t
		x := w1Iface(a1)
		_ = x
	}
}

func BenchmarkTWithPtrIfaceNested(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := T{d0: 0, d1: 1, d2: 2, d3: 3, d4: 4, d5: 5, d6: 6, d7: 7}
		var a1 interface{}
		a1 = &t
		x := w1Iface(a1)
		_ = x
	}
}

func BenchmarkTWithPtrAsmEscape(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := T{d0: 0, d1: 1, d2: 2, d3: 3, d4: 4, d5: 5, d6: 6, d7: 7}
		x := w2(&t)
		_ = x
	}
}

func BenchmarkTWithPtrAsmNoEscape(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t := T{d0: 0, d1: 1, d2: 2, d3: 3, d4: 4, d5: 5, d6: 6, d7: 7}
		x := w3(&t)
		_ = x
	}
}
