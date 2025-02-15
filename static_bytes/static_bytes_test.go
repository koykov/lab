package static_bytes

import (
	"testing"
	"unsafe"

	"github.com/koykov/x2bytes"
)

func StaticBytesBuffer() []byte {
	const bufsz = 128
	var a [bufsz]byte
	var h struct {
		ptr      uintptr
		len, cap int
	}
	h.ptr, h.cap = uintptr(unsafe.Pointer(&a)), bufsz
	return *(*[]byte)(unsafe.Pointer(&h))
}

func BenchmarkStaticBytesBuffer(b *testing.B) {
	b.ReportAllocs()
	x := 3.1415
	for i := 0; i < b.N; i++ {
		buf := StaticBytesBuffer()
		buf, _ = x2bytes.ToBytes(buf, &x)
		_ = buf
	}
}
