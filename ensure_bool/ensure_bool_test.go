package ensure_bool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnsureBool(t *testing.T) {
	exampleTrue := []byte("foo: True")
	exampleFalse := []byte("foo: False")
	t.Run("bin/true", func(t *testing.T) {
		x := ensureTrueBin(exampleTrue, 5, binSafe)
		assert.True(t, x)
	})
	t.Run("bin/false", func(t *testing.T) {
		x := ensureFalseBin(exampleFalse, 5, binSafe)
		assert.True(t, x)
	})
	t.Run("bin unsafe/true", func(t *testing.T) {
		x := ensureTrueBin(exampleTrue, 5, binUnsafe)
		assert.True(t, x)
	})
	t.Run("bin unsafe/false", func(t *testing.T) {
		x := ensureFalseBin(exampleFalse, 5, binUnsafe)
		assert.True(t, x)
	})
	t.Run("map/true", func(t *testing.T) {
		x := ensureTrueMap(exampleTrue, 5)
		assert.True(t, x)
	})
	t.Run("map/false", func(t *testing.T) {
		x := ensureFalseMap(exampleFalse, 5)
		assert.True(t, x)
	})
	t.Run("equal/true", func(t *testing.T) {
		x := ensureTrueEqual(exampleTrue, 5)
		assert.True(t, x)
	})
	t.Run("equal/false", func(t *testing.T) {
		x := ensureFalseEqual(exampleFalse, 5)
		assert.True(t, x)
	})
}

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
