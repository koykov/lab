package stalin_sort

import (
	"cmp"
	"slices"
)

type BufMF[T cmp.Ordered] struct {
	fw, rfw, bw, rbw, r, r1 []T

	N *BufMF[T]
}

func (buf *BufMF[T]) Reset() {
	buf.fw, buf.rfw, buf.bw, buf.rbw, buf.r, buf.r1 = buf.fw[:0], buf.rfw[:0], buf.bw[:0], buf.rbw[:0], buf.r[:0], buf.r1[:0]
	if buf.N != nil {
		buf.N.Reset()
	}
}

func MercifulStalinSort[T cmp.Ordered](buf *BufMF[T], a []T) []T {
	n := len(a)
	if n <= 1 {
		return a
	}

	// forward sort
	mx := a[0]
	buf.fw = append(buf.fw, mx)
	_ = a[n-1]
	for i := 1; i < n; i++ {
		if a[i] >= mx {
			mx = a[i]
			buf.fw = append(buf.fw, mx)
		} else {
			buf.rfw = append(buf.rfw, a[i])
		}
	}

	// backward sort
	if m := len(buf.rfw); m > 0 {
		mn := buf.rfw[m-1]
		buf.bw = append(buf.bw, mn)
		_ = buf.rfw[:m-1]
		for i := m - 2; i >= 0; i-- {
			if buf.rfw[i] <= mn {
				mn = buf.rfw[i]
				buf.bw = append(buf.bw, mn)
			} else {
				buf.rbw = append(buf.rbw, buf.rfw[i])
			}
		}
		slices.Reverse(buf.bw)
	}

	buf.r = merge(buf.r[:0], buf.fw, buf.bw)
	if len(buf.rbw) == 0 {
		return buf.r
	}

	if buf.N == nil {
		buf.N = &BufMF[T]{}
	}
	rem := MercifulStalinSort[T](buf.N, buf.rbw)
	buf.r1 = merge(buf.r1, buf.r, rem)
	return buf.r1
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
