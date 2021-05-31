package timezone

import "testing"

type offsetRecord struct {
	offset, offsetDST int32
}

var (
	offsets = map[string]offsetRecord{
		"Africa/Cairo":      {2, 2},
		"Africa/Casablanca": {1, 0},
		"Africa/Ceuta":      {1, 2},
		"America/Curacao":   {-4, -4},
	}
)

func BenchmarkOffsetTZ(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		x, ok := offsets["Africa/Casablanca"]
		_, _ = x, ok
	}
}
