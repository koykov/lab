package ensure_bool

import "testing"

func BenchmarkEnsureBool(b *testing.B) {
	exampleTrue := []byte("foo: True")
	exampleFalse := []byte("foo: False")
	b.Run("bin/true", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ensureTrueBin(exampleTrue, 4, binSafe)
		}
	})
	b.Run("bin/false", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ensureFalseBin(exampleFalse, 4, binSafe)
		}
	})
	b.Run("bin unsafe/true", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ensureTrueBin(exampleTrue, 4, binUnsafe)
		}
	})
	b.Run("bin unsafe/false", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ensureFalseBin(exampleFalse, 4, binUnsafe)
		}
	})
	b.Run("map/true", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ensureTrueMap(exampleTrue, 4)
		}
	})
	b.Run("map/false", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ensureFalseMap(exampleFalse, 4)
		}
	})
	b.Run("equal/true", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ensureTrueEqual(exampleTrue, 4)
		}
	})
	b.Run("equal/false", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = ensureFalseEqual(exampleFalse, 4)
		}
	})
}
