package generic_interface

import (
	"hash/crc32"
	"testing"

	"github.com/koykov/fastconv"
)

type Bytes interface {
	~string | ~[]byte
}

type I[T Bytes] interface {
	Ier(text T) uint32
}

type X[T Bytes] struct{}

func (x X[T]) Ier(text T) uint32 {
	var b []byte
	switch any(text).(type) {
	case string:
		b = fastconv.S2B(string(text))
	case []byte:
		b = []byte(text)
	}
	return crc32.ChecksumIEEE(b)
}

func BenchmarkIer(b *testing.B) {
	var (
		s0 = "foobar"
		s1 = []byte("foobar")
		e  = uint32(2666930069)
	)
	b.Run("string", func(b *testing.B) {
		var x I[string]
		x = &X[string]{}
		x.Ier("foobar")
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			y := x.Ier(s0)
			if y != e {
				b.Fail()
			}
		}
	})
	b.Run("bytes", func(b *testing.B) {
		var x I[[]byte]
		x = &X[[]byte]{}
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			y := x.Ier(s1)
			if y != e {
				b.Fail()
			}
		}
	})
}
