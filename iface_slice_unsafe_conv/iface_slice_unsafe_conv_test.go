package iface_slice_unsafe_conv

import (
	"bytes"
	"io"
	"testing"
)

var buf []io.ReadWriter

func init() {
	for i := 0; i < 10; i++ {
		b := bytes.NewBuffer(nil)
		b.WriteString("foo")
		buf = append(buf, b)
	}
}

func TestConvUnsafe(t *testing.T) {
	t.Run("safe", func(t *testing.T) {
		b := make([]byte, 64)
		rdr := conv(buf)
		for i := 0; i < len(rdr); i++ {
			r := rdr[i]
			n, _ := r.Read(b)
			println(string(b[:n]))
		}
	})
	t.Run("unsafe", func(t *testing.T) {
		b := make([]byte, 64)
		rdr := convUnsafe(buf)
		for i := 0; i < len(rdr); i++ {
			r := rdr[i]
			n, _ := r.Read(b)
			println(string(b[:n]))
		}
	})
}

func BenchmarkConvUnsafe(b *testing.B) {
	b.Run("safe", func(b *testing.B) {
		b.ReportAllocs()
		b1 := make([]byte, 64)
		for i := 0; i < b.N; i++ {
			rdr := conv(buf)
			_, _ = rdr[0].Read(b1)
		}
	})
	b.Run("unsafe", func(b *testing.B) {
		b.ReportAllocs()
		b1 := make([]byte, 64)
		for i := 0; i < b.N; i++ {
			rdr := convUnsafe(buf)
			_, _ = rdr[0].Read(b1)
		}
	})
}
