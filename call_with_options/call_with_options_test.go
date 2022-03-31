package call_with_options

import (
	"testing"
)

func BenchmarkCallWithOptions(b *testing.B) {
	t := T{l: "foobar"}
	s := "foobar"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		t.F(s, Options{})
		t.reset()
	}
}
