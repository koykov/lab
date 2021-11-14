package entry_search

import (
	"sort"
)

const (
	entryLen     = 1e7
	entryLenHalf = entryLen / 2
	lrLimit      = 256
)

type entry struct {
	hash   uint64
	offset uint32
	length uint16
	expire uint32
	aidptr *uint32
}

var n uint32

func searchCustom(a []entry) uint32 {
	el := uint32(len(a))
	var i, z uint32
	if el < lrLimit {
		for i = 0; i < el; i++ {
			if a[i].expire <= n {
				z = i
				break
			}
		}
	} else {
		var x, y, step uint32
		step = el / 256
		for y = step; y < el; y += step {
			if a[y].expire >= n {
				break
			}
			x = y
		}
		if step < lrLimit {
			for i = x; i < y; i++ {
				if a[i].expire == n {
					z = i
					break
				}
			}
		} else {
			h := step/2 + 1
			for i = 0; i < h; i++ {
				if a[x+i].expire >= n {
					z = x + i
					break
				}
				if a[x+h+i].expire > n {
					z = x + h + i - 1
					break
				}
			}
		}
	}
	return z
}

func searchNative(a []entry) uint32 {
	z := sort.Search(len(a), func(i int) bool { return n <= a[i].expire })
	return uint32(z)
}
