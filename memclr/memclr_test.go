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
	testfn = []func([]uint64){
		memclr64generic,
		memclr64SSE2,
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
	for i := 0; i < 10; i++ {
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

func tmemclr(p []byte, clearfn func([]uint64)) {
	n := len(p)
	if n == 0 {
		return
	}
	if n >= 32 {
		n64 := (n - n%32) / 8
		type sh struct {
			p    uintptr
			l, c int
		}
		h := sh{p: uintptr(unsafe.Pointer(&p[0])), l: n64, c: n64}
		p64 := *(*[]uint64)(unsafe.Pointer(&h))
		clearfn(p64)
		n = n - n%32
	}
	_ = p[n-1]
	for i := 0; i < len(p); i++ {
		p[i] = 0
	}
}

func TestMemclr(t *testing.T) {
	for _, fn := range testfn {
		for _, st := range stages {
			t.Run(fmt.Sprintf("%s/%d", funcName(fn), len(st)), func(t *testing.T) {
				tmemclr(st, fn)
				if testsum(st) != 0 {
					t.Errorf("sum is not zero")
				}
			})
		}
	}
}

func BenchmarkMemclr(b *testing.B) {
	for _, fn := range testfn {
		for _, st := range stages {
			b.Run(fmt.Sprintf("%s/%d", funcName(fn), len(st)), func(b *testing.B) {
				b.ReportAllocs()
				b.SetBytes(int64(len(st)))
				for i := 0; i < b.N; i++ {
					tmemclr(st, fn)
				}
			})
		}
	}
}

func funcName(i interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return name[strings.LastIndex(name, ".")+1:]
}
