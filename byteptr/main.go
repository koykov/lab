package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := []byte("foobar")
	b, c := a[0:3], a[3:]
	fmt.Printf("%s at %d\n%s at %d\n%s at %d", a, ptr(a), b, ptr(b), c, ptr(c))
	// After append old underlying array becomes a garbage, but b and c points to it.
	a = append(a, []byte("extra")...)
	fmt.Printf("\n%s at %d\n%s at %d\n%s at %d", a, ptr(a), b, ptr(b), c, ptr(c))
}

func ptr(p []byte) uint64 {
	h := *(*reflect.SliceHeader)(unsafe.Pointer(&p))
	return uint64(h.Data)
}
