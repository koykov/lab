package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int, 100)
	for i := 0; i < 100; i++ {
		m[i] = i
	}

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1e6; j++ {
				v := m[j%1000]
				_ = v
			}
		}()
	}
	wg.Wait()
	fmt.Println("done")
}
