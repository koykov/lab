package call_with_options

import (
	"testing"

	"github.com/koykov/inspector"
)

type T struct {
	l, r string
	buf  []byte
}

type O struct {
	Flag bool
	Ins  inspector.Inspector
}

func (t *T) F(s string, o O) {
	t.r = s
	o.Ins.DeepEqual(&t.l, &t.r)
	t.buf = append(t.buf, s...)
}

func (t *T) reset() {
	t.buf = t.buf[:0]
}

func BenchmarkCallWithOptions(b *testing.B) {
	t := T{l: "foobar"}
	s := "foobar"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t.F(s, O{Ins: &inspector.StaticInspector{}})
		t.reset()
	}
}
