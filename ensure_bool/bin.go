package ensure_bool

import "unsafe"

func ensureTrueBin(src []byte, offset int) bool {
	bin1 := bin(src, offset, 1)
	bin2 := bin(src, offset, 1)
	bin3 := bin(src, offset, 1)
	bin4 := bin(src, offset, 1)
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

func ensureFalseBin(src []byte, offset int) bool {
	bin1 := bin(src, offset, 1)
	bin2 := bin(src, offset, 1)
	bin3 := bin(src, offset, 1)
	bin5 := bin(src, offset, 1)
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

func bin(src []byte, offset, size int) uint64 {
	n := len(src)
	if offset+size >= n {
		return 0
	}
	_ = src[n-1]
	binSrc := src[offset : offset+size]
	bin := *(*uint64)(unsafe.Pointer(&binSrc[0]))
	return bin
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
	binBoolTrue = [11]uint64{}
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
	binBoolFalse = [11]uint64{}
)

func init() {
	for i := 0; i < len(bBoolTrue); i++ {
		binBoolTrue[i] = *(*uint64)(unsafe.Pointer(&bBoolTrue[i][0]))
	}
	for i := 0; i < len(bBoolFalse); i++ {
		binBoolFalse[i] = *(*uint64)(unsafe.Pointer(&bBoolFalse[i][0]))
	}
}
