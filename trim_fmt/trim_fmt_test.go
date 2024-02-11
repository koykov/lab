package trim_fmt

import (
	"fmt"
	"strings"
	"testing"
)

type stage struct {
	x *X
	i int
}

var stages []stage

func init() {
	for i := 0; i < 100; i++ {
		stg := stage{
			x: &X{buf: []byte(strings.Repeat(" ", i) + "a")},
			i: i,
		}
		stages = append(stages, stg)
	}
}

func TestTrimFmt(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(fmt.Sprintf("goto v0/%d", i), func(t *testing.T) {
			if j, _ := stg.x.trimFmtGotoV0(0); j != stg.i {
				t.FailNow()
			}
		})
		t.Run(fmt.Sprintf("goto v1/%d", i), func(t *testing.T) {
			if j, _ := stg.x.trimFmtGotoV1(0); j != stg.i {
				t.FailNow()
			}
		})
		t.Run(fmt.Sprintf("for v0/%d", i), func(t *testing.T) {
			if j, _ := stg.x.trimFmtForV0(0); j != stg.i {
				t.FailNow()
			}
		})
		t.Run(fmt.Sprintf("for v1/%d", i), func(t *testing.T) {
			if j, _ := stg.x.trimFmtForV1(0); j != stg.i {
				t.FailNow()
			}
		})
	}
}

func BenchmarkTrimFmt(b *testing.B) {
	b.Run(fmt.Sprintf("goto v0"), func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			stg := &stages[j%len(stages)]
			stg.x.trimFmtGotoV0(0)
		}
	})
	b.Run(fmt.Sprintf("goto v1"), func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			stg := &stages[j%len(stages)]
			stg.x.trimFmtGotoV1(0)
		}
	})
	b.Run(fmt.Sprintf("for v0"), func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			stg := &stages[j%len(stages)]
			stg.x.trimFmtForV0(0)
		}
	})
	b.Run(fmt.Sprintf("for v1"), func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			stg := &stages[j%len(stages)]
			stg.x.trimFmtForV1(0)
		}
	})
	b.Run(fmt.Sprintf("fj v0"), func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			stg := &stages[j%len(stages)]
			stg.x.trimFmtFJV0(0)
		}
	})
	b.Run(fmt.Sprintf("fj v1"), func(b *testing.B) {
		for j := 0; j < b.N; j++ {
			stg := &stages[j%len(stages)]
			stg.x.trimFmtFJV1(0)
		}
	})
}
