package combine

import "testing"

func TestU32to64(t *testing.T) {
	var x, y uint32
	x = 15
	y = 342342
	z := U32To64(x, y)
	if z != 64424851782 {
		t.Error("combined uint64 mismatch, need", 64424851782, "got", z)
	}
	x1, y1 := U64To32(z)
	if x1 != x {
		t.Error("decombined uint32 mismatch, need", x, "got", x1)
	}
	if y1 != y {
		t.Error("decombined uint32 mismatch, need", y, "got", y1)
	}
}

func BenchmarkU32to64(b *testing.B) {
	var x, y uint32
	x = 15
	y = 342342
	for i := 0; i < b.N; i++ {
		z := U32To64(x, y)
		if z != 64424851782 {
			b.Error("combined uint64 mismatch, need", 64424851782, "got", z)
		}
		x1, y1 := U64To32(z)
		if x1 != x {
			b.Error("decombined uint32 mismatch, need", x, "got", x1)
		}
		if y1 != y {
			b.Error("decombined uint32 mismatch, need", y, "got", y1)
		}
	}
}
