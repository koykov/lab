package main

// See unicode/utf8/utf8.go

import (
	"unicode/utf8"

	"github.com/koykov/fastconv"
)

const (
	locb = 0b10000000
	hicb = 0b10111111

	xx = 0xF1
	as = 0xF0
	s1 = 0x02
	s2 = 0x13
	s3 = 0x03
	s4 = 0x23
	s5 = 0x34
	s6 = 0x04
	s7 = 0x44
)

type acceptRange struct {
	lo uint8
	hi uint8
}

var (
	first = [256]uint8{
		as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
		as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
		as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
		as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
		as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
		as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
		as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
		as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as,
		xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
		xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
		xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
		xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
		xx, xx, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1,
		s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1,
		s2, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s4, s3, s3,
		s5, s6, s6, s6, s7, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx,
	}
	acceptRanges = [16]acceptRange{
		0: {locb, hicb},
		1: {0xA0, hicb},
		2: {locb, 0x9F},
		3: {0x90, hicb},
		4: {locb, 0x8F},
	}
)

// Strsan sanitizes string by replacing surrogate pair parts with repl byte.
//
// Similar to utf8.Valid(), but uses offset instead of shift among subslices.
func Strsan(s string, repl byte) string {
	if len(s) == 0 {
		return s
	}
	buf := fastconv.S2B(s)
	off := 0
	for len(buf)-off >= 8 {
		first32 := uint32(buf[off+0]) | uint32(buf[off+1])<<8 | uint32(buf[off+2])<<16 | uint32(buf[off+3])<<24
		second32 := uint32(buf[off+4]) | uint32(buf[off+5])<<8 | uint32(buf[off+6])<<16 | uint32(buf[off+7])<<24
		if (first32|second32)&0x80808080 != 0 {
			break
		}
		off += 8
	}
	n := len(buf) - off
	for i := 0; i < n; {
		j := off + i
		si := buf[j]
		if si < utf8.RuneSelf {
			i++
			continue
		}
		x := first[si]
		if x == xx {
			buf[j] = repl
			i++
			continue
		}
		size := int(x & 7)
		if i+size > n {
			buf[j] = repl
			i += size
			continue
		}
		accept := acceptRanges[x>>4]
		if c := buf[j+1]; c < accept.lo || accept.hi < c {
			buf[j] = repl
			i += size
			continue
		} else if size == 2 {
		} else if c := buf[j+2]; c < locb || hicb < c {
			buf[j] = repl
			i += size
			continue
		} else if size == 3 {
		} else if c := buf[j+3]; c < locb || hicb < c {
			buf[j] = repl
			i += size
			continue
		}
		i += size
	}
	return fastconv.B2S(buf)
}
