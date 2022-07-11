package byteconv

import (
	"bytes"
	"testing"
)

var (
	origin    = []byte("foobar")
	ibytes    = interface{}(origin)
	ibytesptr = interface{}(&origin)
	ibytes8   = interface{}([]uint8(origin))
	sbyte     = string(rune(56))
)

func TestByteconv(t *testing.T) {
	t.Run("bytes", func(t *testing.T) {
		x, _ := byteconv(ibytes)
		if !bytes.Equal(x, origin) {
			t.FailNow()
		}
	})
	t.Run("bytesptr", func(t *testing.T) {
		x, _ := byteconv(ibytesptr)
		if !bytes.Equal(x, origin) {
			t.FailNow()
		}
	})
	t.Run("uint8", func(t *testing.T) {
		x, _ := byteconv(ibytes8)
		if !bytes.Equal(x, origin) {
			t.FailNow()
		}
	})
	t.Run("single byte", func(t *testing.T) {
		x := byte2str(56)
		if x != sbyte {
			t.FailNow()
		}
	})
}

func BenchmarkByteconv(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		x := byte2str(56)
		if x != sbyte {
			b.FailNow()
		}
	}
}
