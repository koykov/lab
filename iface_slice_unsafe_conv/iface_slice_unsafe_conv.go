package iface_slice_unsafe_conv

import (
	"io"
	"unsafe"
)

func convUnsafe(src []io.ReadWriter) []io.Reader {
	return *(*[]io.Reader)(unsafe.Pointer(&src))
}
