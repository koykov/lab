package bincheck

import (
	"math"

	"github.com/koykov/byteconv"
	"github.com/koykov/vector"
)

func tokenHash(src []byte, offset int) (hsum uint64) {
	lim := 8
	if rest := len(src) - offset; rest < 8 {
		lim = rest
	}
	_ = tokend[math.MaxUint8-1]
	for i := 0; i < lim && !tokend[i]; i++ {
		hsum = hsum | uint64(src[offset+i])<<(i*8)
	}
	return
}

func ensureNullOrBool(src []byte, offset int, typ *vector.Type, b *bool) bool {
	hsum := tokenHash(src, offset)
	idx := hsum % 10000
	_ = bca[9999]
	if tuple := bca[idx]; tuple != nil {
		*typ = tuple.t
		*b = tuple.b
		return true
	}
	return false
}

type bct struct {
	t vector.Type
	b bool
}

var (
	bcRegistry = map[string]bct{
		"~":     {vector.TypeNull, false},
		"null":  {vector.TypeNull, false},
		"None":  {vector.TypeNull, false},
		"NONE":  {vector.TypeNull, false},
		"y":     {vector.TypeBool, true},
		"Y":     {vector.TypeBool, true},
		"on":    {vector.TypeBool, true},
		"On":    {vector.TypeBool, true},
		"ON":    {vector.TypeBool, true},
		"yes":   {vector.TypeBool, true},
		"Yes":   {vector.TypeBool, true},
		"YES":   {vector.TypeBool, true},
		"true":  {vector.TypeBool, true},
		"True":  {vector.TypeBool, true},
		"TRUE":  {vector.TypeBool, true},
		"n":     {vector.TypeBool, false},
		"N":     {vector.TypeBool, false},
		"no":    {vector.TypeBool, false},
		"No":    {vector.TypeBool, false},
		"NO":    {vector.TypeBool, false},
		"off":   {vector.TypeBool, false},
		"Off":   {vector.TypeBool, false},
		"OFF":   {vector.TypeBool, false},
		"false": {vector.TypeBool, false},
		"False": {vector.TypeBool, false},
		"FALSE": {vector.TypeBool, false},
	}
	bca    [10000]*bct
	tokend [math.MaxUint8]bool
)

func init() {
	for k, v := range bcRegistry {
		x := binSafe(byteconv.S2B(k), 0, len(k))
		c := x % 10000
		bca[c] = &v
	}

	tokend[' '] = true
	tokend[','] = true
	tokend['\n'] = true
	tokend['\r'] = true
	tokend['\t'] = true
	tokend[']'] = true
	tokend['['] = true
	tokend['}'] = true
	tokend['{'] = true
}
