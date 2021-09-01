package combine

func I32To64(lo, hi int32) int64 {
	return int64(lo)<<32 | int64(hi)
}

func I64To32(x int64) (lo, hi int32) {
	lo = int32(x >> 32)
	hi = int32(x & 0xffffffff)
	return
}
