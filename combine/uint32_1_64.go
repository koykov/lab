package combine

func U321To64(lo, hi uint32, f uint8) uint64 {
	return uint64(lo)<<32 | uint64(hi)<<1 | uint64(f)
}

func U64To321(x uint64) (lo, hi uint32, f uint8) {
	lo = uint32(x >> 32)
	hi = uint32((x << 32) >> 33)
	f = uint8((x << 63) >> 63)
	return
}
