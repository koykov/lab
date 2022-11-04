package main

func main() {
	for i := 1; i < 30; i++ {
		println(fib(i), fibm(i))
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
