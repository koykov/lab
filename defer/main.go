package main

import "math"

func main() {
	println(foobar())
}

func foobar() (n int) {
	var n1 int
	defer func() { n = n1 }()
	n1 = math.MaxInt - 123
	return
}
