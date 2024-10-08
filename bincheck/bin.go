package bincheck

import "unsafe"

func ensureTrueBin(src []byte, offset int, binFn func(src []byte, offset, size int) uint64) bool {
	bin1 := binFn(src, offset, 1)
	bin2 := binFn(src, offset, 2)
	bin3 := binFn(src, offset, 3)
	bin4 := binFn(src, offset, 4)
	_ = binBoolTrue[10]
	return bin1 == binBoolTrue[0] ||
		bin1 == binBoolTrue[1] ||
		bin2 == binBoolTrue[2] ||
		bin2 == binBoolTrue[3] ||
		bin2 == binBoolTrue[4] ||
		bin3 == binBoolTrue[5] ||
		bin3 == binBoolTrue[6] ||
		bin3 == binBoolTrue[7] ||
		bin4 == binBoolTrue[8] ||
		bin4 == binBoolTrue[9] ||
		bin4 == binBoolTrue[10]
}

func ensureFalseBin(src []byte, offset int, binFn func(src []byte, offset, size int) uint64) bool {
	bin1 := binFn(src, offset, 1)
	bin2 := binFn(src, offset, 2)
	bin3 := binFn(src, offset, 3)
	bin5 := binFn(src, offset, 5)
	_ = binBoolFalse[10]
	return bin1 == binBoolFalse[0] ||
		bin1 == binBoolFalse[1] ||
		bin2 == binBoolFalse[2] ||
		bin2 == binBoolFalse[3] ||
		bin2 == binBoolFalse[4] ||
		bin3 == binBoolFalse[5] ||
		bin3 == binBoolFalse[6] ||
		bin3 == binBoolFalse[7] ||
		bin5 == binBoolFalse[8] ||
		bin5 == binBoolFalse[9] ||
		bin5 == binBoolFalse[10]
}

func ensureNullBin(src []byte, offset int, binFn func(src []byte, offset, size int) uint64) bool {
	bin1 := binFn(src, offset, 1)
	bin4 := binFn(src, offset, 4)
	_ = binNull[3]
	return bin1 == binNull[0] ||
		bin4 == binNull[1] ||
		bin4 == binNull[2] ||
		bin4 == binNull[3]
}

func binSafe(src []byte, offset, size int) uint64 {
	n := len(src)
	if offset+size > n {
		return 0
	}
	_ = src[n-1]

	switch size {
	case 1:
		return uint64(src[offset])
	case 2:
		return uint64(src[offset+0]) |
			uint64(src[offset+1])<<8
	case 3:
		return uint64(src[offset+0]) |
			uint64(src[offset+1])<<8 |
			uint64(src[offset+2])<<16
	case 4:
		return uint64(src[offset+0]) |
			uint64(src[offset+1])<<8 |
			uint64(src[offset+2])<<16 |
			uint64(src[offset+3])<<24
	case 5:
		return uint64(src[offset+0]) |
			uint64(src[offset+1])<<8 |
			uint64(src[offset+2])<<16 |
			uint64(src[offset+3])<<24 |
			uint64(src[offset+4])<<32
	case 6:
		return uint64(src[offset+0]) |
			uint64(src[offset+1])<<8 |
			uint64(src[offset+2])<<16 |
			uint64(src[offset+3])<<24 |
			uint64(src[offset+4])<<32 |
			uint64(src[offset+5])<<40
	case 7:
		return uint64(src[offset+0]) |
			uint64(src[offset+1])<<8 |
			uint64(src[offset+2])<<16 |
			uint64(src[offset+3])<<24 |
			uint64(src[offset+4])<<32 |
			uint64(src[offset+5])<<40 |
			uint64(src[offset+6])<<48
	case 8:
		return uint64(src[offset+0]) |
			uint64(src[offset+1])<<8 |
			uint64(src[offset+2])<<16 |
			uint64(src[offset+3])<<24 |
			uint64(src[offset+4])<<32 |
			uint64(src[offset+5])<<40 |
			uint64(src[offset+6])<<48 |
			uint64(src[offset+7])<<56
	default:
		return 0
	}
}

func binUnsafe(src []byte, offset, size int) uint64 {
	n := len(src)
	if offset+size > n {
		return 0
	}
	binSrc := src[offset : offset+size]
	u := *(*uint64)(unsafe.Pointer(&binSrc[0]))
	return u & binSizeMasks[size-1]
}

var (
	bBoolTrue = [11][]byte{
		[]byte("y"),
		[]byte("Y"),
		[]byte("on"),
		[]byte("On"),
		[]byte("ON"),
		[]byte("yes"),
		[]byte("Yes"),
		[]byte("YES"),
		[]byte("true"),
		[]byte("True"),
		[]byte("TRUE"),
	}
	binBoolTrue [11]uint64
	bBoolFalse  = [11][]byte{
		[]byte("n"),
		[]byte("N"),
		[]byte("no"),
		[]byte("No"),
		[]byte("NO"),
		[]byte("off"),
		[]byte("Off"),
		[]byte("OFF"),
		[]byte("false"),
		[]byte("False"),
		[]byte("FALSE"),
	}
	binBoolFalse [11]uint64

	bNull = [4][]byte{
		[]byte("~"),
		[]byte("null"),
		[]byte("None"),
		[]byte("NONE"),
	}
	binNull [4]uint64

	binSizeMasks = [8]uint64{
		0x00000000000000FF,
		0x000000000000FFFF,
		0x0000000000FFFFFF,
		0x00000000FFFFFFFF,
		0x000000FFFFFFFFFF,
		0x0000FFFFFFFFFFFF,
		0x00FFFFFFFFFFFFFF,
		0xFFFFFFFFFFFFFFFF,
	}
)

func init() {
	for i := 0; i < len(bBoolTrue); i++ {
		binBoolTrue[i] = binSafe(bBoolTrue[i], 0, len(bBoolTrue[i]))
	}
	for i := 0; i < len(bBoolFalse); i++ {
		binBoolFalse[i] = binSafe(bBoolFalse[i], 0, len(bBoolFalse[i]))
	}
	for i := 0; i < len(bNull); i++ {
		binNull[i] = binSafe(bNull[i], 0, len(bNull[i]))
	}
}
