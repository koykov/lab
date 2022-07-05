package main

import "testing"

func BenchmarkArrayAlloc(b *testing.B) {
	b.Run("array", func(b *testing.B) {
		b.ReportAllocs()
		var x int32
		for i := 0; i < b.N; i++ {
			var a [10]int32
			for j := 0; j < 10; j++ {
				a[j] = int32(j)
			}
			x = a[0]
		}
		_ = x
	})
	b.Run("slice", func(b *testing.B) {
		b.ReportAllocs()
		var x int32
		for i := 0; i < b.N; i++ {
			a := make([]int32, 10)
			for j := 0; j < 10; j++ {
				a[j] = int32(j)
			}
			x = a[0]
		}
		_ = x
	})
}
