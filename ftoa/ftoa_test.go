package main

import (
	"bytes"
	"strconv"
	"testing"
)

var (
	buf    = make([]byte, 0)
	expect = []byte("446.15625")
)

func BenchmarkF64toa(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf = f64toa(buf, 446.15625)
		if !bytes.Equal(buf, expect) {
			// b.Error("f64toa mismatch")
		}
	}
}

func BenchmarkAppendFloat(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf = strconv.AppendFloat(buf, 446.15625, 'f', -1, 64)
		if !bytes.Equal(buf, expect) {
			// b.Error("AppendFloat mismatch")
		}
	}
}
