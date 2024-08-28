package ensure_bool

import "github.com/koykov/byteconv"

func ensureTrueMap(src []byte, offset int) bool {
	b1 := byteconv.B2S(src[offset : offset+1])
	b2 := byteconv.B2S(src[offset : offset+2])
	b3 := byteconv.B2S(src[offset : offset+3])
	b4 := byteconv.B2S(src[offset : offset+4])
	return mapBoolTrue[b1] ||
		mapBoolTrue[b2] ||
		mapBoolTrue[b3] ||
		mapBoolTrue[b4]
}

func ensureFalseMap(src []byte, offset int) bool {
	b1 := byteconv.B2S(src[offset : offset+1])
	b2 := byteconv.B2S(src[offset : offset+2])
	b3 := byteconv.B2S(src[offset : offset+3])
	b5 := byteconv.B2S(src[offset : offset+5])
	return mapBoolFalse[b1] ||
		mapBoolFalse[b2] ||
		mapBoolFalse[b3] ||
		mapBoolFalse[b5]
}

func ensureNullMap(src []byte, offset int) bool {
	b1 := byteconv.B2S(src[offset : offset+1])
	b4 := byteconv.B2S(src[offset : offset+4])
	return mapNull[b1] || mapNull[b4]
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
	mapNull = map[string]bool{
		"~":    true,
		"null": true,
		"None": true,
		"NONE": true,
	}
)
