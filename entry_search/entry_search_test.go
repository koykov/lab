package entry_search

import (
	"testing"
	"time"
)

var a = make([]entry, 0, entryLen)

func init() {
	now := uint32(time.Now().Unix())
	var acc int
	for i := 0; i < entryLen; i++ {
		a = append(a, entry{expire: now})
		if i == entryLenHalf {
			n = now
		}
		acc++
		if acc == 123456 {
			now++
			acc = 0
		}
	}
}

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
