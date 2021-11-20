package clock

import (
	"testing"
	"time"
)

func TestClock(t *testing.T) {
	t.Run("now", func(t *testing.T) {
		c := clock{delta: time.Millisecond}
		c.Start()
		for i := 0; i < 500; i++ {
			t.Log(time.Now().Sub(c.Now()))
			time.Sleep(time.Millisecond)
		}
		c.Stop()
	})
}
