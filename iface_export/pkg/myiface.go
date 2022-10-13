package pkg

import "errors"

type myiface interface {
	My() int
}

type X struct {
	New func() any
}

func (x X) Do() error {
	if x.New == nil {
		return errors.New("no New() func provided")
	}
	x_ := x.New()
	x__, ok := x_.(myiface)
	if ok {
		_ = x__.My()
		return nil
	}
	return errors.New("incompatible item provided by New()")
}
