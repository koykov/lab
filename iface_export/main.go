package main

import "github.com/koykov/lab/iface_export/pkg"

type myX int

func (x myX) My() int {
	return int(x)
}

func main() {
	x := pkg.X{New: func() any {
		var x_ myX
		x_ = 5
		return &x_
	}}
	_ = x.Do()
}
