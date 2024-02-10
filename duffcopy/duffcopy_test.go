package duffcopy

import "testing"

func BenchmarkDuffcopy(b *testing.B) {
	b.Run("native", func(b *testing.B) {
		b.ReportAllocs()
		var s storage
		for i := 0; i < b.N; i++ {
			for j := 0; j < 100; j++ {
				k, n := s.getNode()
				n.typ = 15
				s.putNode(k, n)
			}
			s.reset()
		}
	})
}
