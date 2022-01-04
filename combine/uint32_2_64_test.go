package combine

import "testing"

func TestU322to64(t *testing.T) {
	var (
		x, y  uint32
		f, t0 uint8
	)
	x = 15
	y = 342342
	f = 1
	t0 = 5
	z := U322To64(x, y, f, t0)
	e := uint64(9583660011071289670)
	if z != e {
		t.Error("combined uint64 mismatch, need", e, "got", z)
	}
	x1, y1, f1, t1 := U64To322(z)
	if x1 != x {
		t.Error("decombined uint32 mismatch, need", x, "got", x1)
	}
	if y1 != y {
		t.Error("decombined uint32 mismatch, need", y, "got", y1)
	}
	if f1 != f {
		t.Error("decombined uint32 mismatch, need", f, "got", f1)
	}
	if t1 != t0 {
		t.Error("decombined uint32 mismatch, need", t, "got", t1)
	}
}

func BenchmarkU322to64(b *testing.B) {
	var (
		x, y uint32
		f, t uint8
	)
	x = 15
	y = 342342
	f = 1
	t = 5
	e := uint64(9583660011071289670)
	for i := 0; i < b.N; i++ {
		z := U322To64(x, y, f, t)
		if z != e {
			b.Error("combined uint64 mismatch, need", e, "got", z)
		}
		x1, y1, f1, t1 := U64To322(z)
		if x1 != x {
			b.Error("decombined uint32 mismatch, need", x, "got", x1)
		}
		if y1 != y {
			b.Error("decombined uint32 mismatch, need", y, "got", y1)
		}
		if f1 != f {
			b.Error("decombined uint32 mismatch, need", f, "got", f1)
		}
		if t1 != t {
			b.Error("decombined uint32 mismatch, need", t, "got", t1)
		}
	}
}
