package main

import (
	"fmt"
)

type I interface {
	Foo(int)
	Bar(float64)
}

type T0 struct{}

func (t T0) Foo(_ int)      {}
func (t *T0) Bar(_ float64) {}

type T1 struct{}

func (t *T1) Foo(_ int)     {}
func (t *T1) Bar(_ float64) {}

func calc(x any) error {
	if _, ok := x.(I); ok {
		return nil
	}
	return fmt.Errorf("incompatible")
}

func main() {
	t0 := T0{}
	t1 := T1{}
	var err error
	err = calc(t0)
	fmt.Println(err)
	err = calc(t1)
	fmt.Println(err)
	err = calc(&t1)
	fmt.Println(err)
}
