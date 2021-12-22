package reset_slice

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/koykov/bytealg"
)

func resetUnsafe(p []byte) []byte {
	h := (*reflect.SliceHeader)(unsafe.Pointer(&p))
	reset1(h.Data, h.Cap)
	return p
}

// func reset1(addr uintptr, l int) {
// 	for i := 0; i < l; i++ {
// 		*(*byte)(unsafe.Pointer(addr + uintptr(i))) = 0
// 	}
// }

func resetSafe(p []byte) []byte {
	n := cap(p)
	if n == 0 {
		return p
	}
	p = bytealg.Grow(p, n)
	_ = p[n-1]

	for i := 0; i < n; i++ {
		p[i] = 0
	}
	return p[:0]
}

func BenchmarkResetSlice(b *testing.B) {
	b.Run("unsafe", func(b *testing.B) {
		x := make([]byte, 10, 10)
		x = bytealg.Map(func(r rune) rune {
			return 124
		}, x)
		x = x[:5]
		for i := 0; i < b.N; i++ {
			_ = resetUnsafe(x)
		}
	})
	// b.Run("amd64", func(b *testing.B) {
	// 	x := make([]byte, 10, 10)
	// 	x = bytealg.Map(func(r rune) rune {
	// 		return 124
	// 	}, x)
	// 	x = x[:5]
	// 	for i := 0; i < b.N; i++ {
	// 		_ = resetAmd64(x)
	// 	}
	// })
	b.Run("safe", func(b *testing.B) {
		x := make([]byte, 10, 10)
		x = bytealg.Map(func(r rune) rune {
			return 124
		}, x)
		x = x[:5]
		for i := 0; i < b.N; i++ {
			_ = resetSafe(x)
		}
	})
}
