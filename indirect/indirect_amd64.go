// +build !appengine
// +build gc
// +build !purego

package indirect

// Indirect uint32 value from uintptr pointer, avoiding using unsafe.Pointer and as a result without unsafe.Pointer
// misuse warning.
//
//go:noescape
func uptrIndirAsm(addr uintptr) uint32
