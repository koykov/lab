package bincheck

import (
	"github.com/koykov/byteconv"
	"github.com/koykov/vector"
)

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
	bca [10000]bct
)

func init() {
	for k, v := range bcRegistry {
		x := binSafe(byteconv.S2B(k), 0, len(k))
		c := x % 10000
		bca[c] = v
	}
}
