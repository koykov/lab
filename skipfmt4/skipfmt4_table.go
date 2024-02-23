package skipfmt4

const maxTable = 255

func skipFmt4Table(src []byte, n, offset int) (int, bool) {
	_ = src[n-1]
	_ = skipTable[maxTable]
	for ; offset < n && skipTable[src[offset]]; offset++ {
	}
	return offset, offset == n
}

var skipTable [maxTable + 1]bool

func init() {
	skipTable[' '] = true
	skipTable['\t'] = true
	skipTable['\n'] = true
	skipTable['\r'] = true
}
