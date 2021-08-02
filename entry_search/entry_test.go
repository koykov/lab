package entry_search

import "testing"

func TestSearchCustom(t *testing.T) {
	x := searchCustom(a)
	if a[x].expire != a[x-1].expire+1 {
		t.Error("custom search: wrong expired entry bound")
	}
}

func TestSearchNative(t *testing.T) {
	x := searchNative(a)
	if a[x].expire != a[x-1].expire+1 {
		t.Error("native search: wrong expired entry bound")
	}
}

func BenchmarkSearchCustom(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		x := searchCustom(a)
		if a[x].expire != a[x-1].expire+1 {
			b.Error("custom search: wrong expired entry bound")
		}
	}
}

func BenchmarkSearchNative(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		x := searchNative(a)
		if a[x].expire != a[x-1].expire+1 {
			b.Error("custom search: wrong expired entry bound")
		}
	}
}
