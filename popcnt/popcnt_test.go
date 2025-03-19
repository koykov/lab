package popcnt

import (
	"encoding/binary"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

var (
	stages = []struct {
		n uint64
		c uint64
	}{
		{0xFFFFFFFFFFFFFFFF, 64},
		{0x0000000000000000, 0},
		{0x5555555555555555, 32},
	}
	stagesBig = [][]byte{
		make([]byte, 1e3),
		make([]byte, 1e4),
		make([]byte, 1e5),
		make([]byte, 1e6),
		make([]byte, 1e7),
		make([]byte, 1e8),
		make([]byte, 1e9),
		make([]byte, 1e10),
	}
	funcs = []func([]byte) uint64{
		popcntScalar,
		popcntU64,
		popcntTable,
		popcntU64AVX2,
	}
)

func TestPopcnt(t *testing.T) {
	for _, f := range funcs {
		t.Run(funcName(f), func(t *testing.T) {
			for _, s := range stages {
				t.Run("", func(t *testing.T) {
					var buf [8]byte
					binary.LittleEndian.PutUint64(buf[:], s.n)
					r := f(buf[:])
					if r != s.c {
						t.Errorf("popcnt(%x) = %d, want %d", buf[:], r, s.c)
					}
				})
			}
		})
	}
}

func BenchmarkPopcnt(b *testing.B) {
	for _, f := range funcs {
		b.Run(funcName(f), func(b *testing.B) {
			for _, s := range stages {
				b.Run("", func(b *testing.B) {
					var buf [8]byte
					binary.LittleEndian.PutUint64(buf[:], s.n)
					for i := 0; i < b.N; i++ {
						f(buf[:])
					}
				})
			}
		})
	}
}

func BenchmarkPopcntBig(b *testing.B) {
	for _, f := range funcs {
		b.Run(funcName(f), func(b *testing.B) {
			for _, s := range stagesBig {
				b.Run(strconv.Itoa(len(s)), func(b *testing.B) {
					b.SetBytes(int64(len(s)))
					for i := 0; i < b.N; i++ {
						f(s)
					}
				})
			}
		})
	}
}

func funcName(i interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return name[strings.LastIndex(name, ".")+1:]
}
