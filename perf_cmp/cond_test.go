package test

import (
	"bytes"
	"testing"
)

var (
	a, _b int
	c, d  string
	e, f  []byte
)

func BenchmarkComparisonIf(b *testing.B) {
	a, _b = 5, 5
	c, d = "foo0", "foo1"
	e, f = []byte("bar0"), []byte("bar1")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if a != _b {
			b.Error("mismatch a and b integers")
		}
		if c == d {
			b.Error("mismatch c and d strings")
		}
		if bytes.Equal(e, f) {
			b.Error("mismatch e and f bytes")
		}
	}
}

func BenchmarkComparisonSwitch(b *testing.B) {
	a, _b = 5, 5
	c, d = "foo0", "foo1"
	e, f = []byte("bar0"), []byte("bar1")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		switch {
		case a != _b:
			b.Error("mismatch a and b integers")
		case c == d:
			b.Error("mismatch c and d strings")
		case bytes.Equal(e, f):
			b.Error("mismatch e and f bytes")
		}
	}
}
