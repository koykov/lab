package stalin_sort

import (
	"testing"
)

var stagesMF = []struct {
	a, r []int
}{
	{
		a: []int{1, 2, 10, 3, 2, 4, 15, 6, 30, 20},
		r: []int{1, 2, 2, 3, 4, 6, 10, 15, 20, 30},
	},
	{
		a: []int{78, 33, 100, 61, 65, 72, 11, 66, 89, 3},
		r: []int{3, 11, 33, 61, 65, 66, 72, 78, 89, 100},
	},
	{
		a: []int{2, 2, 3, 1, 10},
		r: []int{1, 2, 2, 3, 10},
	},
	{
		a: []int{1, 2, 10, 3, 2, 4, 15, 6, 30, 20},
		r: []int{1, 2, 2, 3, 4, 6, 10, 15, 20, 30},
	},
	{
		a: []int{1, 2, 2, 3, 2, 5},
		r: []int{1, 2, 2, 2, 3, 5},
	},
}

func TestMercifulStalinSort(t *testing.T) {
	eq := func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	for i := 0; i < len(stagesMF); i++ {
		t.Run("", func(t *testing.T) {
			var buf BufMF[int]
			r := MercifulStalinSort(&buf, stagesMF[i].a)
			if !eq(r, stagesMF[i].r) {
				t.Fail()
			}
		})
	}
}

func BenchmarkMercifulStalinSort(b *testing.B) {
	for i := 0; i < len(stagesMF); i++ {
		b.Run("", func(b *testing.B) {
			var buf BufMF[int]
			b.ReportAllocs()
			for n := 0; n < b.N; n++ {
				buf.Reset()
				MercifulStalinSort(&buf, stagesMF[i].a)
			}
		})
	}
}
