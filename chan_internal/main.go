package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type stream struct {
	data  chan interface{}
	clbuf []byte
}

func newStream(size uint64) *stream {
	s := &stream{}
	s.data = make(chan interface{}, size)
	var i interface{}
	i = s.data

	chptr := *(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&i)) + unsafe.Sizeof(uint(0))))
	fmt.Println("m ptr", chptr)
	chptr += unsafe.Sizeof(uint(0)) * 2
	chptr += unsafe.Sizeof(0)
	chptr += unsafe.Sizeof(uint16(0))
	h := reflect.SliceHeader{
		Data: chptr,
		Len:  4,
		Cap:  4,
	}
	s.clbuf = *(*[]byte)(unsafe.Pointer(&h))

	return s
}

func (s *stream) push(x interface{}) {
	s.data <- x
}

func (s *stream) pull() interface{} {
	return <-s.data
}

func (s *stream) closed() bool {
	if len(s.clbuf) == 0 {
		return true
	}
	return s.clbuf[2] == 1
}

func (s *stream) close() {
	close(s.data)
}

func main() {
	ch := newStream(100000)
	fmt.Println(ch.closed())
	ch.push(1)
	ch.push("foo")
	fmt.Println(ch.closed())
	ch.close()
	fmt.Println(ch.closed())
}
