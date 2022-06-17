package main

type Foo struct {
	x int32
}

func (f *Foo) Bar() {
	f.baz()
}

func (f *Foo) baz() {
	f.x = 1
}

func main() {
	var f *Foo
	f.Bar()
}
