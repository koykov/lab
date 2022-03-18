package lost_buffer

import (
	"bytes"
	"reflect"
	"testing"
	"unsafe"
)

var (
	p0 = []byte("https://google.com/")
	p1 = []byte("notxsxfz4bnhtmw3yc3jtl4hvowtsunysw25xm46v6d2wuyyzhsl3k7xy3nfeucvmi6hc23sneduqmps4cdmgprbbsvfotxdmct5exhchhk42mhoj2nmnm7m3sbmtghggope4h5zkgshlmrzioathl6csdi23b5lvxdurgjtv7bjbunnq6v23r2i5eyi7o7w4")
	p2 = []byte("?foo=bar&x=qwerty")
	e  = []byte("https://google.com/notxsxfz4bnhtmw3yc3jtl4hvowtsunysw25xm46v6d2wuyyzhsl3k7xy3nfeucvmi6hc23sneduqmps4cdmgprbbsvfotxdmct5exhchhk42mhoj2nmnm7m3sbmtghggope4h5zkgshlmrzioathl6csdi23b5lvxdurgjtv7bjbunnq6v23r2i5eyi7o7w4?foo=bar&x=qwerty")
)

func TestLostBuffer(t *testing.T) {
	t.Run("buffer", func(t *testing.T) {
		var buf bytes.Buffer
		var (
			dst    [][]byte
			lo, hi int
		)
		for i := 0; i < 5; i++ {
			lo = buf.Len()
			_, _ = buf.Write(p0)
			_, _ = buf.Write(p1)
			_, _ = buf.Write(p2)
			hi = buf.Len()
			dst = append(dst, buf.Bytes()[lo:hi])
		}
		for i := 0; i < len(dst); i++ {
			h := *(*reflect.SliceHeader)(unsafe.Pointer(&dst[i]))
			t.Logf("dst%d addr %x", i, h.Data)
		}
	})
	t.Run("slice", func(t *testing.T) {
		var buf []byte
		var (
			dst    [][]byte
			lo, hi int
		)
		for i := 0; i < 5; i++ {
			lo = len(buf)
			buf = append(buf, p0...)
			buf = append(buf, p1...)
			buf = append(buf, p2...)
			hi = len(buf)
			dst = append(dst, buf[lo:hi])
		}
		bh := *(*reflect.SliceHeader)(unsafe.Pointer(&buf))
		t.Logf("buf addr %x", bh.Data)
		for i := 0; i < len(dst); i++ {
			h := *(*reflect.SliceHeader)(unsafe.Pointer(&dst[i]))
			t.Logf("dst%d addr %x", i, h.Data)
		}
	})
}
