package regex_alloc

import (
	"regexp"
	"testing"
)

func BenchmarkRegexAlloc(b *testing.B) {
	re := regexp.MustCompile(`Chrome/(\d+\.[\.\d]+).+Sparrow`)
	s := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.31031 Safari/537.36 Sparrow"
	b.Run("match", func(b *testing.B) {
		b.ReportAllocs()
		var x bool
		for i := 0; i < b.N; i++ {
			x = re.MatchString(s)
		}
		_ = x
	})
	b.Run("find submatch", func(b *testing.B) {
		b.ReportAllocs()
		var x []string
		for i := 0; i < b.N; i++ {
			x = re.FindStringSubmatch(s)
		}
		_ = x
	})
	b.Run("find submatch index", func(b *testing.B) {
		b.ReportAllocs()
		var x []int
		for i := 0; i < b.N; i++ {
			x = re.FindStringSubmatchIndex(s)
		}
		_ = x
	})
}
