package index_runes

import (
	"strings"
	"testing"

	"github.com/koykov/byteconv"
)

func BenchmarkIndexRunes(b *testing.B) {
	s := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean non lectus dui. Nullam gravida purus libero, sit amet interdum massa pretium viverra. Praesent cursus eu mauris nec rhoncus. Fusce dignissim justo et lorem fermentum eleifend. Sed nisi orci, hendrerit quis mauris vitae, scelerisque blandit risus. Ut imperdiet fermentum diam, vel dapibus velit ornare a. Vivamus non mattis ante. Morbi semper, tortor a convallis rutrum, augue diam ullamcorper ligula, ut luctus nisi orci vitae nunc. Suspendisse dictum porttitor est, id lacinia neque bibendum vel. Suspendisse tristique scelerisque nisi quis consequat. Sed sit amet pulvinar nulla, nec lacinia elit.`
	target := "pulvinar nulla"
	sr, targetr := []rune(s), []rune(target)
	b.Run("native", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			strings.Index(string(sr), string(targetr))
		}
	})
	b.Run("buffer", func(b *testing.B) {
		b.ReportAllocs()
		var bufs, buft []byte
		for i := 0; i < b.N; i++ {
			var s1, t1 string
			bufs, s1 = byteconv.AppendR2S(bufs[:0], sr)
			buft, t1 = byteconv.AppendR2S(buft[:0], targetr)
			strings.Index(s1, t1)
		}
	})
}

func BenchmarkConvertSingleRune(b *testing.B) {
	b.ReportAllocs()
	r := []rune("А")
	for i := 0; i < b.N; i++ {
		x := string(r[0])
		_ = x
	}
}
