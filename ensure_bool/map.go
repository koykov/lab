package ensure_bool

import "github.com/koykov/byteconv"

func ensureTrueMap(src []byte, offset int) bool {
	b1 := src[offset : offset+1]
	b2 := src[offset : offset+2]
	b3 := src[offset : offset+3]
	b4 := src[offset : offset+4]
	return mapBoolTrue[byteconv.B2S(b1)] ||
		mapBoolTrue[byteconv.B2S(b2)] ||
		mapBoolTrue[byteconv.B2S(b3)] ||
		mapBoolTrue[byteconv.B2S(b4)]
}

func ensureFalseMap(src []byte, offset int) bool {
	b1 := src[offset : offset+1]
	b2 := src[offset : offset+2]
	b3 := src[offset : offset+3]
	b5 := src[offset : offset+4]
	return mapBoolFalse[byteconv.B2S(b1)] ||
		mapBoolFalse[byteconv.B2S(b2)] ||
		mapBoolFalse[byteconv.B2S(b3)] ||
		mapBoolFalse[byteconv.B2S(b5)]
}

var (
	mapBoolTrue = map[string]bool{
		"y":    true,
		"Y":    true,
		"on":   true,
		"On":   true,
		"ON":   true,
		"yes":  true,
		"Yes":  true,
		"YES":  true,
		"true": true,
		"True": true,
		"TRUE": true,
	}
	mapBoolFalse = map[string]bool{
		"n":     true,
		"N":     true,
		"no":    true,
		"No":    true,
		"NO":    true,
		"off":   true,
		"Off":   true,
		"OFF":   true,
		"false": true,
		"False": true,
		"FALSE": true,
	}
)
