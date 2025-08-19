package skipfmt4

import (
	"fmt"
	"testing"
)

type stage struct {
	s        []byte
	off, exp int
	eof      bool
}

func TestSkipFmt4(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		t.Run(fmt.Sprintf("generic/%d", i), func(t *testing.T) {
			r, eof := skipFmt4(stg.s, len(stg.s), stg.off)
			if r != stg.exp || eof != stg.eof {
				t.FailNow()
			}
		})
		t.Run(fmt.Sprintf("table/%d", i), func(t *testing.T) {
			r, eof := skipFmt4Table(stg.s, len(stg.s), stg.off)
			if r != stg.exp || eof != stg.eof {
				t.FailNow()
			}
		})
		t.Run(fmt.Sprintf("bits/%d", i), func(t *testing.T) {
			r, eof := skipFmt4Bits(stg.s, len(stg.s), stg.off)
			if r != stg.exp || eof != stg.eof {
				t.FailNow()
			}
		})
		t.Run(fmt.Sprintf("SSE2/%d", i), func(t *testing.T) {
			r, eof := skipFmt4SSE2(stg.s, len(stg.s), stg.off)
			if r != stg.exp || eof != stg.eof {
				skipFmt4SSE2(stg.s, len(stg.s), stg.off)
				t.Errorf("need %d, got %d", stg.exp, r)
			}
		})
	}
}

func BenchmarkSkipFmt4(b *testing.B) {
	for i := 0; i < len(stages); i++ {
		stg := &stages[i]
		b.Run(fmt.Sprintf("generic/%d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				skipFmt4(stg.s, len(stg.s), stg.off)
			}
		})
		b.Run(fmt.Sprintf("table/%d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				skipFmt4Table(stg.s, len(stg.s), stg.off)
			}
		})
		b.Run(fmt.Sprintf("bits/%d", i), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				skipFmt4Bits(stg.s, len(stg.s), stg.off)
			}
		})
		print("\n")
	}
}

var stages = []stage{
	{
		s:   []byte("{\n\"key\":1}"),
		off: 1,
		exp: 2,
	},
	{
		s:   []byte("{\r \"key\":1}"),
		off: 1,
		exp: 3,
	},
	{
		s:   []byte("{\r\t\t\"key\":1}"),
		off: 1,
		exp: 4,
	},
	{
		s:   []byte("{\n   \"key\":1}"),
		off: 1,
		exp: 5,
	},
	{
		s:   []byte("{\n    \"key\":1}"),
		off: 1,
		exp: 6,
	},
	{
		s:   []byte("{\n     \"key\":1}"),
		off: 1,
		exp: 7,
	},
	{
		s:   []byte("{\n      \"key\":1}"),
		off: 1,
		exp: 8,
	},
	{
		s:   []byte("{\n       \"key\":1}"),
		off: 1,
		exp: 9,
	},
	{
		s:   []byte("{\n        \"key\":1}"),
		off: 1,
		exp: 10,
	},
	{
		s:   []byte("{\n         \"key\":1}"),
		off: 1,
		exp: 11,
	},
	{
		s:   []byte("{\n          \"key\":1}"),
		off: 1,
		exp: 12,
	},
	{
		s:   []byte("{\n           \"key\":1}"),
		off: 1,
		exp: 13,
	},
	{
		s:   []byte("{\n            \"key\":1}"),
		off: 1,
		exp: 14,
	},
	{
		s:   []byte("{\n             \"key\":1}"),
		off: 1,
		exp: 15,
	},
	{
		s:   []byte("{\n              \"key\":1}"),
		off: 1,
		exp: 16,
	},
	{
		s:   []byte("{\n               \"key\":1}"),
		off: 1,
		exp: 17,
	},
	{
		s:   []byte("{\n                \"key\":1}"),
		off: 1,
		exp: 18,
	},
	{
		s:   []byte("{\n                 \"key\":1}"),
		off: 1,
		exp: 19,
	},
	{
		s:   []byte("{\n                  \"key\":1}"),
		off: 1,
		exp: 20,
	},
	{
		s:   []byte("{\n                      \"key\":1}"),
		off: 1,
		exp: 24,
	},
	{
		s:   []byte("{\n                          \"key\":1}"),
		off: 1,
		exp: 28,
	},
	{
		s:   []byte("{\n                              \"key\":1}"),
		off: 1,
		exp: 32,
	},
	{
		s:   []byte("{\n                                                              \"key\":1}"),
		off: 1,
		exp: 64,
	},
	{
		s:   []byte("{\n                                                                                              \"key\":1}"),
		off: 1,
		exp: 96,
	},
	{
		s:   []byte("{\n                                                                                                                              \"key\":1}"),
		off: 1,
		exp: 128,
	},
}
