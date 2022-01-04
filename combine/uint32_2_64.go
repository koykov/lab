package combine

func U322To64(lo, hi uint32, a, b uint8) uint64 {
	return uint64(a)<<63 | uint64(b)<<56 | uint64(lo)<<28 | uint64(hi)
}

func U64To322(x uint64) (lo, hi uint32, a, b uint8) {
	a = uint8(x >> 63)
	b = uint8((x >> 56) & 0b01111111)
	lo = uint32((x >> 28) & 0x0fffffff)
	hi = uint32(x & 0x0fffffff)
	return
}
