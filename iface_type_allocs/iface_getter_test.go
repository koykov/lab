package test

import (
	"strconv"
	"testing"
)

type T struct {
	kv   map[string]string
	bufI int64
	bufU uint64
	bufF float64
	bufS string
}

var (
	x = &T{kv: map[string]string{
		"foo": "123456",
		"bar": "qwerty",
	}}
)

func (t T) Get(key string, def interface{}) interface{} {
	raw, ok := t.kv[key]
	if !ok {
		return def
	}
	var (
		val interface{}
		err error
	)
	switch def.(type) {
	case int, int8, int16, int32, int64:
		val, err = strconv.ParseInt(raw, 10, 64)
	case uint, uint8, uint16, uint32, uint64:
		val, err = strconv.ParseUint(raw, 10, 64)
	case float32, float64:
		val, err = strconv.ParseFloat(raw, 64)
	default:
		val = raw
	}

	if err != nil {
		return def
	}
	return val
}

func (t *T) GetBuf(key string, def interface{}) interface{} {
	raw, ok := t.kv[key]
	if !ok {
		return def
	}
	switch def.(type) {
	case int, int8, int16, int32, int64:
		t.bufI, _ = strconv.ParseInt(raw, 10, 64)
		return &t.bufI
	case uint, uint8, uint16, uint32, uint64:
		t.bufU, _ = strconv.ParseUint(raw, 10, 64)
		return &t.bufU
	case float32, float64:
		t.bufF, _ = strconv.ParseFloat(raw, 64)
		return &t.bufF
	default:
		t.bufS = raw
		return &t.bufS
	}
}

func (t T) GetInt(key string, def int) int {
	raw, ok := t.kv[key]
	if !ok {
		return def
	}
	if n, err := strconv.ParseInt(raw, 10, 64); err == nil {
		return int(n)
	}
	return def
}

func (t T) GetString(key, def string) string {
	if raw, ok := t.kv[key]; ok {
		return raw
	}
	return def
}

func BenchmarkIfaceGetInt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		raw := x.Get("foo", int32(0))
		if n, ok := raw.(int64); !ok || n != 123456 {
			b.Error("int mismatch")
		}
	}
}

func BenchmarkIfaceGetString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		raw := x.Get("bar", "N/D")
		if s, ok := raw.(string); !ok || s != "qwerty" {
			b.Error("string mismatch")
		}
	}
}

func BenchmarkIfaceBufferedGetInt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		raw := x.GetBuf("foo", int32(0))
		if n, ok := raw.(*int64); !ok || *n != 123456 {
			b.Error("int mismatch")
		}
	}
}

func BenchmarkIfaceBufferedGetString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		raw := x.GetBuf("bar", "N/D")
		if s, ok := raw.(*string); !ok || *s != "qwerty" {
			b.Error("string mismatch")
		}
	}
}

func BenchmarkExactGetInt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		raw := x.GetInt("foo", 0)
		if raw != 123456 {
			b.Error("int mismatch")
		}
	}
}

func BenchmarkExactGetString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		raw := x.GetString("bar", "N/D")
		if raw != "qwerty" {
			b.Error("string mismatch")
		}
	}
}
