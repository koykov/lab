package hal

import (
	"testing"

	"github.com/koykov/halvector"
)

var (
	stages = []string{
		"da,en;q=0.9,en-GB;q=0.8,en-US;q=0.7",
		"ru,en;q=0.9",
		"RU,ru;q=0.9,en-US;q=0.8,en;q=0.7",
		"es-ES,es;q=0.9",
	}
)

func TestHAL(t *testing.T) {
	for _, stage := range stages {
		t.Run(stage, func(t *testing.T) {
			vec := halvector.NewVector()
			if err := vec.ParseStr(stage); err != nil {
				t.Error(err)
				return
			}
			t.Log(vec.Root().DotString("0.code"))
		})
	}
}
