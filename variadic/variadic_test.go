package variadic

import (
	"fmt"
	"testing"
)

func varf(f string, args ...interface{}) string {
	return fmt.Sprintf(f, args...)
}

func BenchmarkVariadic(b *testing.B) {
	b.Run("default", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			varf("%#v %#v %#v", 1, 3.1415, "qwerty")
		}
	})
	b.Run("predefined", func(b *testing.B) {
		var args = []interface{}{1, 3.1415, "qwerty"}
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			varf("%#v %#v %#v", args...)
		}
	})
}
