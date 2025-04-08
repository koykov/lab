package memclr

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"unsafe"
)

var (
	stages   [][]byte
	stages64 []struct {
		buf64 []uint64
		buf   []byte
		sz    int64
	}
	testfn = []struct {
		blocksz int
		fn      func([]uint64)
	}{
		{8, memclr64generic},
		{8, memclr64SSE2},
		{8, memclr64AVX2},
		{8, memclr64AVX512},
	}
)

func init() {
	pow := func(n int) int {
		r := 1
		for i := 0; i < n; i++ {
			r *= 10
		}
		return r
	}
	for i := 2; i < 10; i++ {
		p := make([]byte, pow(i))
		p64, p8 := tconv64(p)
		stages = append(stages, p)
		stages64 = append(stages64, struct {
			buf64 []uint64
			buf   []byte
			sz    int64
		}{p64, p8, int64(len(p))})
	}
	fillStages()
}

func fillStages() {
	for i := 0; i < len(stages); i++ {
		for j := 0; j < len(stages[i]); j++ {
			stages[i][j] = byte(j % 256)
		}
	}
}

func testsum(p []byte) (r uint64) {
	for i := 0; i < len(p); i++ {
		r += uint64(p[i])
	}
	return
}

func tconv64(p []byte) (p64 []uint64, p8 []byte) {
	n := len(p)
	if n == 0 {
		return
	}
	n64 := (n - n%8) / 8
	p64 = *(*[]uint64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(&p[0])),
		Len:  n64,
		Cap:  n64,
	}))
	p8 = p[n64*8:]
	return
}

func tmemclr(p []byte, blocksz int, clearfn func([]uint64)) {
	n := len(p)
	if n == 0 {
		return
	}
	if n >= blocksz {
		n64 := (n - n%blocksz) / 8
		type sh struct {
			p    uintptr
			l, c int
		}
		h := sh{p: uintptr(unsafe.Pointer(&p[0])), l: n64, c: n64}
		p64 := *(*[]uint64)(unsafe.Pointer(&h))
		clearfn(p64)
		p = p[n64*8:]
		n = len(p)
	}
	if n == 0 {
		return
	}
	_ = p[n-1]
	for i := 0; i < len(p); i++ {
		p[i] = 0
	}
}

func tmemclr64(p64 []uint64, p []byte, clearfn func([]uint64)) {
	n64, n := len(p64), len(p)
	if n64 > 0 {
		clearfn(p64)
	}
	if n > 0 {
		_ = p[n-1]
		for i := 0; i < len(p); i++ {
			p[i] = 0
		}
	}
}

func TestMemclr(t *testing.T) {
	for _, fn := range testfn {
		fillStages()
		for _, st := range stages {
			t.Run(fmt.Sprintf("%s/%d", funcName(fn.fn), len(st)), func(t *testing.T) {
				tmemclr(st, fn.blocksz, fn.fn)
				if testsum(st) != 0 {
					t.Errorf("sum is not zero")
					tmemclr(st, fn.blocksz, fn.fn)
				}
			})
		}
	}
}

func BenchmarkMemclr(b *testing.B) {
	for _, fn := range testfn {
		for _, st := range stages64 {
			b.Run(fmt.Sprintf("%s/%d", funcName(fn.fn), st.sz), func(b *testing.B) {
				b.SetBytes(st.sz)
				for i := 0; i < b.N; i++ {
					tmemclr64(st.buf64, st.buf, fn.fn)
				}
			})
		}
	}
}

func BenchmarkMemclrNative(b *testing.B) {
	for _, st := range stages {
		b.Run(fmt.Sprintf("%s/%d", "native", len(st)), func(b *testing.B) {
			b.SetBytes(int64(len(st)))
			for i := 0; i < b.N; i++ {
				Memclr(st)
			}
		})
	}
}

func funcName(i interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return name[strings.LastIndex(name, ".")+1:]
}
