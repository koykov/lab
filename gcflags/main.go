package main

// run me with go tool arguments: -gcflags '-m -m'
func main() {
	var u = uint32(15)
	x := &u
	foo(indirect(x))
}

func indirect(x *uint32) uint32 {
	return *x + 1
}

func foo(y uint32) uint32 {
	return y + 1
}
