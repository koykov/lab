package regex_vs_index

import (
	"regexp"
	"strings"
	"testing"
)

func BenchmarkPerf(b *testing.B) {
	b.Run("regex", func(b *testing.B) {
		var (
			re = regexp.MustCompile(`BashPodder`)
			m  bool
		)
		s := "This is a BashPodder feed reder."
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			m = re.MatchString(s)
		}
		_ = m
	})
	b.Run("index", func(b *testing.B) {
		var m bool
		s := "This is a BashPodder feed reder."
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			m = strings.Index(s, "BashPodder") != -1
		}
		_ = m
	})
}
