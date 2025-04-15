package skipfmt4

//go:noescape
func skipfmt4SSE2(p []byte) int

func skipFmt4SSE2(p []byte, n, off int) (int, bool) {
	r := skipfmt4SSE2(p[off:])
	if r == -1 {
		return n, true
	}
	return r + off, false
}
