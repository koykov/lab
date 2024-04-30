package hexlen

import (
	"math"
	"strconv"
	"testing"
)

const n = 123456

func BenchmarkHexLen(b *testing.B) {
	b.Run("strconv", func(b *testing.B) {
		b.ReportAllocs()
		var buf []byte
		for i := 0; i < b.N; i++ {
			buf = strconv.AppendInt(buf[:0], n, 16)
			// len(buf) == 5
		}
	})
	b.Run("logN", func(b *testing.B) {
		b.ReportAllocs()
		ln16 := math.Log(16)
		for i := 0; i < b.N; i++ {
			l := math.Ceil(math.Log(n) / ln16)
			_ = l
			// int(l) == 5
		}
	})
}
