package label_build_perf

import (
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/koykov/bytebuf"
)

func BenchmarkLabelBuild(b *testing.B) {
	var reg = map[string]int{}
	applyfn := func(label string) {
		reg[label]++
	}
	b.Run("sprintf", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			label := fmt.Sprintf(`myservice_component_feature{metric="%s"}`, "foobar")
			applyfn(label)
		}
	})
	b.Run("string builder", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var sb strings.Builder
			sb.Grow(100)
			sb.WriteString(`myservice_component_feature{metric="`)
			sb.WriteString("foobar")
			sb.WriteString(`"}"`)
			applyfn(sb.String())
		}
	})
	b.Run("string builder pool", func(b *testing.B) {
		b.ReportAllocs()
		var pool = sync.Pool{New: func() any {
			sb := strings.Builder{}
			return &sb
		}}
		for i := 0; i < b.N; i++ {
			sb := pool.Get().(*strings.Builder)
			sb.Grow(100)
			sb.WriteString(`myservice_component_feature{metric="`)
			sb.WriteString("foobar")
			sb.WriteString(`"}"`)
			applyfn(sb.String())
			sb.Reset()
			pool.Put(sb)
		}
	})
	b.Run("chain buffer", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			var buf bytebuf.Chain
			buf.Grow(100)
			buf.WriteString(`myservice_component_feature{metric="`).
				WriteString("foobar").
				WriteString(`"}"`)
			applyfn(buf.String())
		}
	})
	b.Run("chain buffer pool", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			buf := bytebuf.AcquireChain()
			buf.Reset().Grow(100).
				WriteString(`myservice_component_feature{metric="`).
				WriteString("foobar").
				WriteString(`"}"`)
			applyfn(buf.String())
			bytebuf.ReleaseChain(buf)
		}
	})
}
