package main

import "testing"

var u = uint32(15)

func BenchmarkEscAnalyze(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		x := &u
		foo(indirect(x))
	}
}
