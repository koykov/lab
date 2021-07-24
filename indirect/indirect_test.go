package indirect

import (
	"testing"
	"unsafe"
)

var (
	t = uint32(15)
	p = unsafe.Pointer(&t)
	a = uintptr(p)
)

func TestUintptrIndirectAsm(t *testing.T) {
	if x := uptrIndirAsm(a); x != 15 {
		t.Error("uintptr indirect (asm) failed: need", 15, "got", x)
	}
}

func TestUintptrIndirectNative(t *testing.T) {
	if x := uptrIndirNative(a); x != 15 {
		t.Error("uintptr indirect (native) failed: need", 15, "got", x)
	}
}

func BenchmarkUintptrIndirectAsm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if x := uptrIndirAsm(a); x != 15 {
			b.Error("uintptr indirect (asm) failed: need", 15, "got", x)
		}
	}
}

func BenchmarkUintptrIndirectNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if x := uptrIndirNative(a); x != 15 {
			b.Error("uintptr indirect (native) failed: need", 15, "got", x)
		}
	}
}
