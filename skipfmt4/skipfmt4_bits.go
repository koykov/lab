package skipfmt4

import (
	"encoding/binary"
	"unsafe"
)

const maxBits = 538976288

func skipFmt4Bits(src []byte, n, offset int) (int, bool) {
	_ = src[n-1]
	_ = skipBits[maxBits]
	n4, n2 := n-n%4, n-n%2
	for offset < n4 {
		if b := *(*uint32)(unsafe.Pointer(&src[offset])); b <= maxBits && skipBits[b] {
			offset += 4
			continue
		}
		break
	}
	for offset < n2 {
		if b := *(*uint16)(unsafe.Pointer(&src[offset])); int(b) <= maxBits && skipBits[b] {
			offset += 2
			continue
		}
		break
	}
	for ; offset < n && skipBits[src[offset]]; offset++ {
	}
	return offset, offset == n
}

var skipBits [maxBits + 1]bool

func init() {
	skipBits[' '] = true
	skipBits['\t'] = true
	skipBits['\n'] = true
	skipBits['\r'] = true

	skipBits[binary.LittleEndian.Uint16([]byte("\n "))] = true
	skipBits[binary.LittleEndian.Uint16([]byte("\r "))] = true
	skipBits[binary.LittleEndian.Uint16([]byte("\n\t"))] = true
	skipBits[binary.LittleEndian.Uint16([]byte("\r\t"))] = true
	skipBits[binary.LittleEndian.Uint16([]byte("  "))] = true
	skipBits[binary.LittleEndian.Uint16([]byte("\t\t"))] = true

	skipBits[binary.LittleEndian.Uint32([]byte("\n   "))] = true
	skipBits[binary.LittleEndian.Uint32([]byte("\r   "))] = true
	skipBits[binary.LittleEndian.Uint32([]byte("\n\t\t\t"))] = true
	skipBits[binary.LittleEndian.Uint32([]byte("\r\t\t\t"))] = true
	skipBits[binary.LittleEndian.Uint32([]byte("    "))] = true
	skipBits[binary.LittleEndian.Uint32([]byte("\t\t\t\t"))] = true

	fmt4 := [4]byte{' ', '\t', '\n', '\r'}
	var buf [4]byte
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				for l := 0; l < 4; l++ {
					buf[0], buf[1], buf[2], buf[3] = fmt4[i], fmt4[j], fmt4[k], fmt4[l]
					h := binary.LittleEndian.Uint32(buf[:])
					skipBits[h] = true
				}
			}
		}
	}
}
