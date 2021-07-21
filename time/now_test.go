package time

import (
	"sync/atomic"
	"testing"
	"time"
)

var (
	now uint32
)

func init() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				atomic.StoreUint32(&now, uint32(time.Now().Unix()))
			}
		}
	}()
}

func BenchmarkTimeNowNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := time.Now()
		_ = x
	}
}

func BenchmarkTimeNowNativeParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x := time.Now()
			_ = x
		}
	})
}

func BenchmarkTimeNowAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := atomic.LoadUint32(&now)
		_ = x
	}
}

func BenchmarkTimeNowAtomicParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x := atomic.LoadUint32(&now)
			_ = x
		}
	})
}
