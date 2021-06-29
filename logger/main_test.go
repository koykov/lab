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
	a  []interface{}
)

func BenchmarkLoggerIface_Print(b *testing.B) {
	x, y, z := 1, "foo", "bar"
	li = &Log{}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		li.Printf("", x, y, z)
	}
}

func BenchmarkLoggerDirect_Print(b *testing.B) {
	x, y, z := 1, "foo", "bar"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ld.Printf("", x, y, z)
	}
}

func BenchmarkLoggerBufIface_Print(b *testing.B) {
	x, y, z := 1, "foo", "bar"
	li = &Log{}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a = append(a[:0], x, y, z)
		li.Printf("", a...)
	}
}

func BenchmarkLoggerBufDirect_Print(b *testing.B) {
	x, y, z := 1, "foo", "bar"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a = append(a[:0], x, y, z)
		ld.Printf("", a...)
	}
}
