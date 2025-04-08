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
	stages [][]byte
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
		stages = append(stages, make([]byte, pow(i)))
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
		for _, st := range stages {
			b.Run(fmt.Sprintf("%s/%d", funcName(fn.fn), len(st)), func(b *testing.B) {
				b.SetBytes(int64(len(st)))
				for i := 0; i < b.N; i++ {
					tmemclr(st, fn.blocksz, fn.fn)
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
