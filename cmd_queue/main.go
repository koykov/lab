package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	bucketsSize    = 64
	bucketsWorkers = 4
)

var (
	buckets = make([]int, bucketsSize)
)

func init() {
	for i := 0; i < bucketsSize; i++ {
		buckets[i] = i
	}
}

func main() {
	queue := make(chan int, bucketsWorkers)

	var (
		wg   sync.WaitGroup
		proc uint32
	)
	for i := 0; i < bucketsWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				if idx, ok := <-queue; ok {
					time.Sleep(time.Duration(200+rand.Int31n(200)) * time.Millisecond)
					processed := atomic.AddUint32(&proc, 1)
					fmt.Println(processed, idx)
					continue
				}
				break
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < bucketsSize; i++ {
			queue <- i
		}
		close(queue)
	}()

	wg.Wait()
}
