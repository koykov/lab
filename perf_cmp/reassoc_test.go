package test

import "testing"

func BenchmarkMultiply0(b *testing.B) {
	var x, y, z, a float32
	x, y, z, a = 25, 30, 10, 5
	for i := 0; i < 10e9; i++ {
		_ = x * y * z * a
	}
}

func BenchmarkMultiply1(b *testing.B) {
	var x, y, z, a float32
	x, y, z, a = 25, 30, 10, 5
	for i := 0; i < 10e9; i++ {
		_ = (x * y) * (z * a)
	}
}
