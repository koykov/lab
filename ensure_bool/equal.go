package ensure_bool

import "bytes"

func ensureTrueEqual(src []byte, offset int) bool {
	b1 := src[offset : offset+1]
	b2 := src[offset : offset+2]
	b3 := src[offset : offset+3]
	b4 := src[offset : offset+4]
	return bytes.Equal(b1, bBoolTrue[0]) ||
		bytes.Equal(b1, bBoolTrue[1]) ||
		bytes.Equal(b2, bBoolTrue[2]) ||
		bytes.Equal(b2, bBoolTrue[3]) ||
		bytes.Equal(b2, bBoolTrue[4]) ||
		bytes.Equal(b3, bBoolTrue[5]) ||
		bytes.Equal(b3, bBoolTrue[6]) ||
		bytes.Equal(b3, bBoolTrue[7]) ||
		bytes.Equal(b4, bBoolTrue[8]) ||
		bytes.Equal(b4, bBoolTrue[9]) ||
		bytes.Equal(b4, bBoolTrue[10])
}

func ensureFalseEqual(src []byte, offset int) bool {
	b1 := src[offset : offset+1]
	b2 := src[offset : offset+2]
	b3 := src[offset : offset+3]
	b5 := src[offset : offset+5]
	return bytes.Equal(b1, bBoolFalse[0]) ||
		bytes.Equal(b1, bBoolFalse[1]) ||
		bytes.Equal(b2, bBoolFalse[2]) ||
		bytes.Equal(b2, bBoolFalse[3]) ||
		bytes.Equal(b2, bBoolFalse[4]) ||
		bytes.Equal(b3, bBoolFalse[5]) ||
		bytes.Equal(b3, bBoolFalse[6]) ||
		bytes.Equal(b3, bBoolFalse[7]) ||
		bytes.Equal(b5, bBoolFalse[8]) ||
		bytes.Equal(b5, bBoolFalse[9]) ||
		bytes.Equal(b5, bBoolFalse[10])
}

func ensureNullEqual(src []byte, offset int) bool {
	b1 := src[offset : offset+1]
	b4 := src[offset : offset+4]
	return bytes.Equal(b1, bNull[0]) ||
		bytes.Equal(b4, bNull[1]) ||
		bytes.Equal(b4, bNull[2]) ||
		bytes.Equal(b4, bNull[3])
}
