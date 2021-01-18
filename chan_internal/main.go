package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var ch chan interface{}
	ch = make(chan interface{}, 10)
	chcptr := chClosePtr(ch)
	fmt.Println("closed", chcptr)
	ch <- 1
	ch <- "foo"
	fmt.Println("closed", chcptr)
	close(ch)
	fmt.Println("closed", chcptr)
}

func chClosePtr(ch interface{}) []byte {
	chptr := *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&ch)) + unsafe.Sizeof(uint(0))))
	chptr += unsafe.Sizeof(uint(0)) * 2
	chptr += unsafe.Sizeof(0)
	chptr += unsafe.Sizeof(uint16(0))
	h := reflect.SliceHeader{
		Data: chptr,
		Len:  4,
		Cap:  4,
	}
	b := *(*[]byte)(unsafe.Pointer(&h))
	return b
}
