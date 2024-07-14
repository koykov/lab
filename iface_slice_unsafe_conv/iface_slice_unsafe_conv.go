package iface_slice_unsafe_conv

import (
	"io"
	"unsafe"
)

func conv(src []io.ReadWriter) []io.Reader {
	buf := make([]io.Reader, len(src))
	for i := 0; i < len(src); i++ {
		buf[i] = src[i]
	}
	return buf
}

func convUnsafe(src []io.ReadWriter) []io.Reader {
	return *(*[]io.Reader)(unsafe.Pointer(&src))
}
