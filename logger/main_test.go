package logger

import (
	"testing"
)

type Logger interface {
	Printf(format string, v ...interface{})
}

type Log struct{}

func (l *Log) Printf(string, ...interface{}) {}

var (
	li Logger
	ld Log
)

func BenchmarkLoggerIface_Print(b *testing.B) {
	x, y, z := 1, "foo", "bar"
	li = &Log{}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		li.Printf("x", x, "y", y, "z", z)
	}
}

func BenchmarkLoggerDirect_Print(b *testing.B) {
	x, y, z := 1, "foo", "bar"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ld.Printf("x", x, "y", y, "z", z)
	}
}
