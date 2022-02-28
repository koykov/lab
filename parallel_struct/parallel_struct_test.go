package main

import (
	"sync"
	"testing"
)

func TestParallelStruct(t *testing.T) {
	var (
		x  X
		wg sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1e9; i++ {
			x.A = i
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 1e9; i++ {
			x.B = i
		}
	}()
	wg.Wait()
}
