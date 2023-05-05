package byteconv

import (
	"bytes"
	"testing"
)

var (
	s = "foobar"
	p = []byte("foobar")
)

func BenchmarkByteconv(b *testing.B) {
	b.Run("b2s/old", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			x := b2s(p)
			if x != s {
				b.FailNow()
			}
		}
	})
	b.Run("b2s/new", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			x := b2s1(p)
			if x != s {
				b.FailNow()
			}
		}
	})
	b.Run("s2b/old", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			x := s2b(s)
			if !bytes.Equal(x, p) {
				b.FailNow()
			}
		}
	})
	b.Run("s2b/new", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			x := s2b1(s)
			if !bytes.Equal(x, p) {
				b.FailNow()
			}
		}
	})
}
