package main

type I interface {
	Foo() I
	Bar(string, interface{}) I
	Push() error
}

type A struct {
	buf []struct {
		a string
		b interface{}
	}
}

func (a *A) Foo() I { return a }

func (a *A) Bar(x string, y interface{}) I {
	a.buf = append(a.buf, struct {
		a string
		b interface{}
	}{x, y})
	return a
}

func (a A) Push() error { return nil }

func (a *A) reset() {
	a.buf = a.buf[:0]
}

type B struct{}

func (b B) Foo() I                    { return b }
func (b B) Bar(string, interface{}) I { return b }
func (b B) Push() error               { return nil }
