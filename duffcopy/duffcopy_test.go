package duffcopy

import "testing"

func TestDuffcopy(t *testing.T) {
	t.Run("bce unsafe", func(t *testing.T) {
		var s storage
		i, n := s.getNodeBCE()

		for j := 0; j < 10; j++ {
			k, m := s.getNodeBCE()
			m.typ = 10
			m.key.Len = 100
			s.putNodeUnsafe(k, m)
		}
		n.typ, n.limit = 15, 1

		s.putNodeUnsafe(i, n)
		if s.nodes[0].typ != 15 || s.nodes[i].limit != 1 {
			t.Fail()
		}
	})
}

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
	b.Run("bce unsafe", func(b *testing.B) {
		b.ReportAllocs()
		var s storage
		for i := 0; i < b.N; i++ {
			for j := 0; j < 100; j++ {
				k, n := s.getNodeBCE()
				n.typ = 15
				s.putNodeUnsafe(k, n)
			}
			s.reset()
		}
	})
	b.Run("bce unsafe1", func(b *testing.B) {
		b.ReportAllocs()
		var s storage
		for i := 0; i < b.N; i++ {
			for j := 0; j < 100; j++ {
				k, n := s.getNodeBCE()
				n.typ = 15
				s.putNodeUnsafe1(k, n)
			}
			s.reset()
		}
	})
}
