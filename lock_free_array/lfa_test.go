package lock_free_array

import (
	"math"
	"sync"
	"sync/atomic"
	"testing"
)

// lock version
type la struct {
	mux sync.Mutex
	buf []int
}

func (a *la) add(x int) {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.buf = append(a.buf, x)
}

func (a *la) reset() {
	a.buf = a.buf[:0]
}

// lock-free version
type lfa struct {
	idx  uint32
	buf  []int
	len_ uint32
}

func (a *lfa) reserve(len_ int) {
	a.len_ = uint32(len_)
	a.idx = math.MaxUint32
	a.buf = make([]int, len_)
}

func (a *lfa) add(x int) {
	a.buf[atomic.AddUint32(&a.idx, 1)%a.len_] = x
}

func (a *lfa) reset() {
	a.idx = math.MaxUint32
}

func BenchmarkLFA(b *testing.B) {
	const len_ = 100
	b.Run("lock array", func(b *testing.B) {
		var a la
		for i := 0; i < b.N; i++ {
			var wg sync.WaitGroup
			wg.Add(len_)
			for j := 0; j < len_; j++ {
				go func() {
					defer wg.Done()
					a.add(15)
				}()
			}
			wg.Wait()
			a.reset()
		}
	})
	b.Run("lock-free array", func(b *testing.B) {
		var a lfa
		a.reserve(len_)
		for i := 0; i < b.N; i++ {
			var wg sync.WaitGroup
			wg.Add(len_)
			for j := 0; j < len_; j++ {
				go func() {
					defer wg.Done()
					a.add(15)
				}()
			}
			wg.Wait()
			a.reset()
		}
	})
}
