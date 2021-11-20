package iface_dymmy_log

import (
	"log"
	"os"
	"testing"
)

type T struct {
	L Logger
}

func BenchmarkDummyLog(b *testing.B) {
	b.Run("dummy", func(b *testing.B) {
		x := T{L: &DummyLog{}}
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			x.L.Printf("dummy log call with args %d, %s, %f", 10, "foo", 3.1515)
		}
	})
	b.Run("log", func(b *testing.B) {
		x := T{L: log.New(os.Stderr, "", log.LstdFlags)}
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			x.L.Printf("dummy log call with args %d, %s, %f", 10, "foo", 3.1515)
		}
	})
}
