package main

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestParallelRead(t *testing.T) {
	x := map[int32]int32{
		1: 1,
		2: 1,
		3: 1,
		4: 1,
		5: 1,
		6: 1,
		7: 1,
		8: 1,
		9: 1,
		0: 1,
	}

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(ctx context.Context) {
			defer wg.Done()
			var s int32
			for {
				select {
				case <-ctx.Done():
					return
				default:
					s += x[rand.Int31n(10)]
				}
			}
		}(ctx)
	}

	time.Sleep(time.Second * 10)
	cancel()
	wg.Wait()
}
