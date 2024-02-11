package index_byte

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/koykov/bytealg"
)

type stage struct {
	buf []byte
	i   int
}

var stages []*stage

func init() {
	for i := 0; i < 10000; i++ {
		stg := &stage{
			buf: []byte(strings.Repeat("0", i) + "1"),
			i:   i,
		}
		stages = append(stages, stg)
	}
}

func TestIndexByte(t *testing.T) {
	t.Run("native", func(t *testing.T) {
		for i := 0; i < len(stages); i++ {
			stg := stages[i]
			j := bytes.IndexByte(stg.buf, '1')
			if j != stg.i {
				t.Errorf("%d vs %d", j, stg.i)
			}
		}
	})
	t.Run("bytealg lur", func(t *testing.T) {
		for i := 0; i < len(stages); i++ {
			stg := stages[i]
			j := bytealg.IndexByteAtLUR(stg.buf, '1', 0)
			if j != stg.i {
				t.Errorf("%d vs %d", j, stg.i)
			}
		}
	})
}

func BenchmarkIndexByte(b *testing.B) {
	d := 1
	for i := 0; i < len(stages); i += d {
		off := i
		b.Run(fmt.Sprintf("native %d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				idx := off + j%50
				stg := stages[idx]
				_ = bytes.IndexByte(stg.buf, '1')
			}
		})
		b.Run(fmt.Sprintf("bytealg lur %d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				idx := off + j%50
				stg := stages[idx]
				_ = bytealg.IndexByteAtLUR(stg.buf, '1', 0)
			}
		})
		switch {
		case i >= 10 && i < 50:
			d = 10
		case i >= 100 && i < 500:
			d = 50
		case i >= 500 && i < 1000:
			d = 100
		case i >= 1000 && i < 5000:
			d = 200
		case i >= 5000:
			d = 500
		}
		println(" ")
	}
}
