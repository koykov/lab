// +build !appengine
// +build gc
// +build !purego

package main

//go:noescape
func hashstr(addr uintptr) uintptr
