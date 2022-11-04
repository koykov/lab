package main

func main() {
	for i := 1; i < 30; i++ {
		println(fib(i), fibm(i), fibO(i))
	}
}

func fib(n int) int {
	if n < 3 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

var mem [100]int

func fibm(n int) int {
	if n < 3 {
		return 1
	}
	if mem[n] > 0 {
		return mem[n]
	}
	mem[n] = fibm(n-1) + fibm(n-2)
	return mem[n]
}

func fibO(n int) int {
	_, x := fib2(n)
	return x
}

func fib2(n int) (int, int) {
	switch n {
	case 0:
		return 0, 1
	case 1, 2:
		return 1, 1
	default:
		prev, next := fib2(n - 1)
		return next, prev + next
	}
}
