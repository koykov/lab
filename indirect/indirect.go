package indirect

import "unsafe"

// Indirect uint32 value from uintptr pointer using unsafe.Pointer. go vet warning about possible misuse of unsafe.Pointer.
func uptrIndirNative(addr uintptr) uint32 {
	return *(*uint32)(unsafe.Pointer(addr))
}
