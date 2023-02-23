package htmlbench

import (
	"html"
	"testing"
)

func BenchmarkEscape(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		s := "<a href='test'>Test</a>"
		e := "&lt;a href=&#39;test&#39;&gt;Test&lt;/a&gt;"
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := html.EscapeString(s)
			if r != e {
				b.Fatalf("'%s' vs '%s'", r, e)
			}
		}
	})
}

func BenchmarkUnescape(b *testing.B) {
	b.Run("simple", func(b *testing.B) {
		s := "&lt;a href=&#39;test&#39;&gt;Test&lt;/a&gt;"
		e := "<a href='test'>Test</a>"
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			r := html.UnescapeString(s)
			if r != e {
				b.Fatalf("'%s' vs '%s'", r, s)
			}
		}
	})
}
