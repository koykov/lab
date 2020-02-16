package crc32

import (
	"bytes"
	"hash/crc32"
	"testing"

	"github.com/koykov/fastconv"
)

var (
	dataShort     = fastconv.S2B(`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean non lectus dui. Nullam gravida purus libero, sit amet interdum massa pretium viverra. Praesent cursus eu mauris nec rhoncus. Fusce dignissim justo et lorem fermentum eleifend. Sed nisi orci, hendrerit quis mauris vitae, scelerisque blandit risus. Ut imperdiet fermentum diam, vel dapibus velit ornare a. Vivamus non mattis ante. Morbi semper, tortor a convallis rutrum, augue diam ullamcorper ligula, ut luctus nisi orci vitae nunc. Suspendisse dictum porttitor est, id lacinia neque bibendum vel. Suspendisse tristique scelerisque nisi quis consequat. Sed sit amet pulvinar nulla, nec lacinia elit.`)
	expectedShort = uint32(0x607650b0)
	dataLong      = bytes.Repeat(dataShort, 1000)
	expectedLong  = uint32(0x1e4674d1)
)

func BenchmarkCrc32NativeSimple(b *testing.B) {
	h := uint32(0xffffffff)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		tab := crc32.MakeTable(0)
		r := crc32.Checksum(dataShort, tab)
		if h != r {
			b.Error(h, "not equal to", r)
		}
	}
}

func BenchmarkCrc32NativeSimpleLong(b *testing.B) {
	h := uint32(0xffffffff)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		tab := crc32.MakeTable(0)
		r := crc32.Checksum(dataLong, tab)
		if h != r {
			b.Error(h, "not equal to", r)
		}
	}
}

func BenchmarkCrc32NativeIEEE(b *testing.B) {
	h := uint32(0x607650b0)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		tab := crc32.MakeTable(crc32.IEEE)
		r := crc32.Checksum(dataShort, tab)
		if h != r {
			b.Error(h, "not equal to", r)
		}
	}
}

func BenchmarkCrc32NativeIEEELong(b *testing.B) {
	h := uint32(0x1e4674d1)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		tab := crc32.MakeTable(crc32.IEEE)
		r := crc32.Checksum(dataLong, tab)
		if h != r {
			b.Error(h, "not equal to", r)
		}
	}
}

func BenchmarkCrc32NativeCastagnoli(b *testing.B) {
	h := uint32(0x9c498ba9)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		tab := crc32.MakeTable(crc32.Castagnoli)
		r := crc32.Checksum(dataShort, tab)
		if h != r {
			b.Error(h, "not equal to", r)
		}
	}
}

func BenchmarkCrc32NativeCastagnoliLong(b *testing.B) {
	h := uint32(0xa74275c5)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		tab := crc32.MakeTable(crc32.Castagnoli)
		r := crc32.Checksum(dataLong, tab)
		if h != r {
			b.Error(h, "not equal to", r)
		}
	}
}

func BenchmarkCrc32Bitwise(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Bitwise(dataShort, 0)
		if r != expectedShort {
			b.Error(expectedShort, "not equal to", r)
		}
	}
}

func BenchmarkCrc32BitwiseLong(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Bitwise(dataLong, 0)
		if r != expectedLong {
			b.Error(expectedLong, "not equal to", r)
		}
	}
}

func BenchmarkCrc32Halfbyte(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Halfbyte(dataShort, 0)
		if r != expectedShort {
			b.Error(expectedShort, "not equal to", r)
		}
	}
}

func BenchmarkCrc32HalfbyteLong(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Halfbyte(dataLong, 0)
		if r != expectedLong {
			b.Error(expectedLong, "not equal to", r)
		}
	}
}

func BenchmarkCrc32Byte1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Byte1(dataShort, 0)
		if r != expectedShort {
			b.Error(expectedShort, "not equal to", r)
		}
	}
}

func BenchmarkCrc32Byte1Long(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Byte1(dataLong, 0)
		if r != expectedLong {
			b.Error(expectedLong, "not equal to", r)
		}
	}
}

func BenchmarkCrc32Byte1Tableless(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Byte1Tableless(dataShort, 0)
		if r != expectedShort {
			b.Error(expectedShort, "not equal to", r)
		}
	}
}

func BenchmarkCrc32Byte1TablelessLong(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Byte1Tableless(dataLong, 0)
		if r != expectedLong {
			b.Error(expectedLong, "not equal to", r)
		}
	}
}

func BenchmarkCrc32Bytes4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Bytes4(dataShort, 0)
		if r != expectedShort {
			b.Error(expectedShort, "not equal to", r)
		}
	}
}

func BenchmarkCrc32Bytes4Long(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Bytes4(dataLong, 0)
		if r != expectedLong {
			b.Error(expectedLong, "not equal to", r)
		}
	}
}

func BenchmarkCrc32Bytes8(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Bytes8(dataShort, 0)
		if r != expectedShort {
			b.Error(expectedShort, "not equal to", r)
		}
	}
}

func BenchmarkCrc32Bytes8Long(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Bytes8(dataLong, 0)
		if r != expectedLong {
			b.Error(expectedLong, "not equal to", r)
		}
	}
}

func BenchmarkCrc32Bytes4x8(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Bytes4x8(dataShort, 0)
		if r != expectedShort {
			b.Error(expectedShort, "not equal to", r)
		}
	}
}

func BenchmarkCrc32Bytes4x8Long(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Bytes4x8(dataLong, 0)
		if r != expectedLong {
			b.Error(expectedLong, "not equal to", r)
		}
	}
}

func BenchmarkCrc32Bytes16(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Bytes16(dataShort, 0)
		if r != expectedShort {
			b.Error(expectedShort, "not equal to", r)
		}
	}
}

func BenchmarkCrc32Bytes16Long(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		r := Bytes16(dataLong, 0)
		if r != expectedLong {
			b.Error(expectedLong, "not equal to", r)
		}
	}
}
