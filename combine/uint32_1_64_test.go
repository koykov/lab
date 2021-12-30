package combine

import "testing"

func TestU321to64(t *testing.T) {
	var (
		x, y uint32
		f    uint8
	)
	x = 15
	y = 342342
	f = 1
	z := U321To64(x, y, 1)
	if z != 64425194125 {
		t.Error("combined uint64 mismatch, need", 64425194125, "got", z)
	}
	x1, y1, f1 := U64To321(z)
	if x1 != x {
		t.Error("decombined uint32 mismatch, need", x, "got", x1)
	}
	if y1 != y {
		t.Error("decombined uint32 mismatch, need", y, "got", y1)
	}
	if f1 != f {
		t.Error("decombined uint32 mismatch, need", f, "got", f1)
	}
}

func BenchmarkU321to64(b *testing.B) {
	var (
		x, y uint32
		f    uint8
	)
	x = 15
	y = 342342
	f = 1
	for i := 0; i < b.N; i++ {
		z := U321To64(x, y, f)
		if z != 64425194125 {
			b.Error("combined uint64 mismatch, need", 64425194125, "got", z)
		}
		x1, y1, f1 := U64To321(z)
		if x1 != x {
			b.Error("decombined uint32 mismatch, need", x, "got", x1)
		}
		if y1 != y {
			b.Error("decombined uint32 mismatch, need", y, "got", y1)
		}
		if f1 != f {
			b.Error("decombined uint32 mismatch, need", f, "got", f1)
		}
	}
}
