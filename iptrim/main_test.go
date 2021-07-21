package main

import (
	"net"
	"testing"

	"github.com/koykov/fastconv"
)

const (
	ipOrigin = "255.255.255.255"
	ipExpect = "255.255.255.0"
)

func trimNative(raw string) string {
	ip := net.ParseIP(raw)
	ip[15] = 0
	return ip.String()
}

func trimCustom(raw string) string {
	i, l := -1, len(raw)-1
	switch {
	case raw[l] == '.':
		i = l
	case raw[l-1] == '.':
		i = l - 1
	case raw[l-2] == '.':
		i = l - 2
	case raw[l-3] == '.':
		i = l - 3
	}
	if i < 0 {
		return raw
	}
	b := fastconv.S2B(raw)
	b = append(b[:i+1], '0')
	return fastconv.B2S(b)
}

func TestIPTrimNative(t *testing.T) {
	if x := trimNative(ipOrigin); x != ipExpect {
		t.Error("trim native fail, need", ipExpect, "got", x)
	}
}

func TestIPTrimCustom(t *testing.T) {
	var buf []byte
	buf = append(buf[:0], ipOrigin...)
	if x := trimCustom(fastconv.B2S(buf)); x != ipExpect {
		t.Error("trim custom fail, need", ipExpect, "got", x)
	}
}

func BenchmarkIPTrimNative(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if x := trimNative(ipOrigin); x != ipExpect {
			b.Error("trim native fail, need", ipExpect, "got", x)
		}
	}
}

func BenchmarkIPTrimCustom(b *testing.B) {
	var buf []byte
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf = append(buf[:0], ipOrigin...)
		if x := trimCustom(fastconv.B2S(buf)); x != ipExpect {
			b.Error("trim custom fail, need", ipExpect, "got", x)
		}
	}
}
