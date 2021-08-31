package combine

func U32To64(lo, hi uint32) uint64 {
	return uint64(lo)<<32 | uint64(hi)
}

func U64To32(x uint64) (lo uint32, hi uint32) {
	lo = uint32(x >> 32)
	hi = uint32(x & 0xffffffff)
	return
}
