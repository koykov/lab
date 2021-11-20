package clock

import (
	"testing"
	"time"
)

func TestClock(t *testing.T) {
	t.Run("now", func(t *testing.T) {
		c := clock{delta: time.Microsecond}
		c.Start()
		for i := 0; i < 50; i++ {
			t.Log(time.Now().Sub(c.Now()))
			time.Sleep(time.Millisecond)
		}
		c.Stop()
	})
}

func BenchmarkClock(b *testing.B) {
	b.Run("now clock", func(b *testing.B) {
		c := clock{delta: time.Microsecond}
		c.Start()
		var n time.Time
		for i := 0; i < b.N; i++ {
			n = c.Now()
		}
		_ = n
		c.Stop()
	})
	b.Run("now native", func(b *testing.B) {
		var n time.Time
		for i := 0; i < b.N; i++ {
			n = time.Now()
		}
		_ = n
	})
}
