package skipfmt4

//go:noescape
func skipfmt4SIMD(p []byte) int

func skipFmt4SSE2(p []byte, n, off int) (int, bool) {
	r := skipfmt4SIMD(p[off:])
	if r == -1 {
		return n, true
	}
	return r + off, false
}
