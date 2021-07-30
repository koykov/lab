package main

type T struct {
	i0, i1, i2, i3, i4, i5, i6, i7 int32
	u0, u1, u2, u3, u4, u5, u6, u7 uint32
	f0, f1, f2, f3, f4, f5, f6, f7 float32
	d0, d1, d2, d3, d4, d5, d6, d7 float64
}

func w0(t T) float64 {
	return t.d0 + t.d1 + t.d2 + t.d3 + t.d4 + t.d5 + t.d6 + t.d7
}

func w1(t *T) float64 {
	return t.d0 + t.d1 + t.d2 + t.d3 + t.d4 + t.d5 + t.d6 + t.d7
}

func w1Iface(x interface{}) float64 {
	t := x.(*T)
	return t.d0 + t.d1 + t.d2 + t.d3 + t.d4 + t.d5 + t.d6 + t.d7
}

func w2(t *T) float64

//go:noescape
func w3(t *T) float64
