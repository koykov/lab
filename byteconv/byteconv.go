//go:build go1.20
// +build go1.20

package byteconv

import (
	"reflect"
	"unsafe"
)

func b2s(p []byte) string {
	return *(*string)(unsafe.Pointer(&p))
}

func s2b(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	var h reflect.SliceHeader
	h.Data = sh.Data
	h.Len = sh.Len
	h.Cap = sh.Len
	return *(*[]byte)(unsafe.Pointer(&h))
}

func b2s1(p []byte) string {
	return unsafe.String(unsafe.SliceData(p), len(p))
}

func s2b1(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
