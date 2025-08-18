package main

import (
	"context"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan int)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(ctx context.Context) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					c <- 1 // heizenpanic
				}
			}
		}(ctx)
	}
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				close(c)
				return
			case x := <-c:
				_ = x
			}
		}
	}(ctx)

	cancel()

	wg.Wait()
}
