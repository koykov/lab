package stalin_sort

import (
	"cmp"
	"slices"
)

func MercifulStalinSort[T cmp.Ordered](buf *BufMF[T], a []T) []T {
	n := len(a)
	if n <= 1 {
		return a
	}
	buf.prealloc(uint32(n))
	fw, rfw, bw, rbw, r, r1 := buf.fw(), buf.rfw(), buf.bw(), buf.rbw(), buf.r(), buf.r1()

	// forward sort
	mx := a[0]
	fw = append(fw, mx)
	_ = a[n-1]
	for i := 1; i < n; i++ {
		if a[i] >= mx {
			mx = a[i]
			fw = append(fw, mx)
		} else {
			rfw = append(rfw, a[i])
		}
	}

	// backward sort
	if m := len(rfw); m > 0 {
		mn := rfw[m-1]
		bw = append(bw, mn)
		_ = rfw[:m-1]
		for i := m - 2; i >= 0; i-- {
			if rfw[i] <= mn {
				mn = rfw[i]
				bw = append(bw, mn)
			} else {
				rbw = append(rbw, rfw[i])
			}
		}
		slices.Reverse(bw)
	}

	r = merge(r, fw, bw)
	if len(rbw) == 0 {
		return r
	}

	if buf.n == nil {
		buf.n = &BufMF[T]{}
	}
	rem := MercifulStalinSort[T](buf.n, rbw)
	r1 = merge(r1, r, rem)
	return r1
}

func merge[T cmp.Ordered](dst, a, b []T) []T {
	n, m := len(a), len(b)
	var i, j int
	for i < n && j < m {
		if a[i] <= b[j] {
			dst = append(dst, a[i])
			i++
		} else {
			dst = append(dst, b[j])
			j++
		}
	}
	dst = append(dst, a[i:]...)
	dst = append(dst, b[j:]...)
	return dst
}
