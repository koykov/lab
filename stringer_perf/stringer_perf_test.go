package stringer_perf

import "testing"

type T int

const (
	t0 T = iota
	t1
	t2
	t3
)

func (t T) String() string {
	switch t {
	case t0:
		return "first"
	case t1:
		return "second"
	case t2:
		return "third"
	case t3:
		return "fourth"
	}
	return ""
}

var tmap = map[T]string{
	t0: "first",
	t1: "second",
	t2: "third",
	t3: "fourth",
}

var tarr = [4]string{
	"first",
	"second",
	"third",
	"fourth",
}

func BenchmarkStringerPerf(b *testing.B) {
	b.Run("switch", func(b *testing.B) {
		var s string
		t := t2
		for i := 0; i < b.N; i++ {
			s = t.String()
		}
		_ = s
	})
	b.Run("map", func(b *testing.B) {
		var s string
		t := t2
		for i := 0; i < b.N; i++ {
			s = tmap[t]
		}
		_ = s
	})
	b.Run("array", func(b *testing.B) {
		var s string
		t := t2
		for i := 0; i < b.N; i++ {
			s = tarr[t]
		}
		_ = s
	})
}
