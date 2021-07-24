package indirect

import "unsafe"

func uptrIndirNative(addr uintptr) uint32 {
	return *(*uint32)(unsafe.Pointer(addr))
}
