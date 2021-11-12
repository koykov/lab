package loop_opt

import "testing"

func BenchmarkLoop(b *testing.B) {
	a := 0
	for i := 0; i < b.N; i++ {
		a = floop(5)
	}
	_ = a
}

func BenchmarkGoto(b *testing.B) {
	a := 0
	for i := 0; i < b.N; i++ {
		a = fgoto(5)
	}
	_ = a
}

func floop(size int) int {
	f := 1
	for i := 1; i <= size; i++ {
		f *= i
	}
	return f
}

func fgoto(size int) int {
	i, f := 0, 1
loop:
	i++
	f *= i
	if i < size {
		goto loop
	}
	return f
}
