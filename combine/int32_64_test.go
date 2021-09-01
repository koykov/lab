package combine

import (
	"math"
	"testing"
)

func TestI32to64(t *testing.T) {
	var x, y int32
	x = math.MinInt32
	y = math.MaxInt32
	z := I32To64(x, y)
	if z != -9223372034707292161 {
		t.Error("combined int64 mismatch, need", -9223372034707292161, "got", z)
	}
	x1, y1 := I64To32(z)
	if x1 != x {
		t.Error("decombined int32 mismatch, need", x, "got", x1)
	}
	if y1 != y {
		t.Error("decombined int32 mismatch, need", y, "got", y1)
	}
}

func BenchmarkI32to64(b *testing.B) {
	var x, y int32
	x = math.MinInt32
	y = math.MaxInt32
	for i := 0; i < b.N; i++ {
		z := I32To64(x, y)
		if z != -9223372034707292161 {
			b.Error("combined int64 mismatch, need", -9223372034707292161, "got", z)
		}
		x1, y1 := I64To32(z)
		if x1 != x {
			b.Error("decombined int32 mismatch, need", x, "got", x1)
		}
		if y1 != y {
			b.Error("decombined int32 mismatch, need", y, "got", y1)
		}
	}
}
