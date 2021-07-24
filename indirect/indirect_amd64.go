// +build !appengine
// +build gc
// +build !purego

package indirect

//go:noescape
func uptrIndirAsm(addr uintptr) uint32
