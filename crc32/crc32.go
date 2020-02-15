package crc32

import (
	"encoding/binary"
)

const (
	Polynomial int64 = 0xedb88320

	MaxSlice16    = 16
	MaxSliceNoLut = 0
)

func Bitwise(data []byte, prevCrc32 uint32) uint32 {
	var crc = prevCrc32 ^ 0xffffffff
	for _, c := range data {
		crc ^= uint32(c)
		for j := 0; j < 8; j++ {
			crc = (crc >> 1) ^ (-(crc & 1) & uint32(Polynomial))
		}
	}
	return crc ^ 0xffffffff
}

func Halfbyte(data []byte, prevCrc32 uint32) uint32 {
	var crc = prevCrc32 ^ 0xffffffff
	for _, c := range data {
		crc = Lookup16[(crc^uint32(c))&0x0f] ^ (crc >> 4)
		crc = Lookup16[(crc^(uint32(c)>>4))&0x0f] ^ (crc >> 4)
	}
	return crc ^ 0xffffffff
}

func Byte1(data []byte, prevCrc32 uint32) uint32 {
	var crc = prevCrc32 ^ 0xffffffff
	for _, c := range data {
		crc = (crc >> 8) ^ Lookup[MaxSliceNoLut][(crc&0xff)^uint32(c)]
	}
	return crc ^ 0xffffffff
}

func Byte1Tableless(data []byte, prevCrc32 uint32) uint32 {
	var crc = prevCrc32 ^ 0xffffffff
	for _, c := range data {
		s := uint8(crc) ^ uint8(c)
		low := uint32((s ^ (s << 6)) & 0xff)
		a := low * ((1 << 23) + (1 << 14) + (1 << 2))
		crc = (crc >> 8) ^
			(low * ((1 << 24) + (1 << 16) + (1 << 8))) ^
			a ^
			(a >> 1) ^
			(low * ((1 << 20) + (1 << 12))) ^
			(low << 19) ^
			(low << 17) ^
			(low >> 2)
	}
	return crc ^ 0xffffffff
}

func Bytes4(data []byte, prevCrc32 uint32) uint32 {
	var crc = prevCrc32 ^ 0xffffffff
	for len(data) >= 4 {
		one := binary.LittleEndian.Uint32(data[:4]) ^ uint32(crc)
		crc = Lookup[0][(one>>24)&0xff] ^
			Lookup[1][(one>>16)&0xff] ^
			Lookup[2][(one>>8)&0xff] ^
			Lookup[3][one&0xff]
		data = data[4:]
	}
	for _, c := range data {
		crc = (crc >> 8) ^ Lookup[0][(crc&0xff)^uint32(c)]
	}
	return crc ^ 0xffffffff
}

func Bytes8(data []byte, prevCrc32 uint32) uint32 {
	var crc = prevCrc32 ^ 0xffffffff
	for len(data) >= 8 {
		one := binary.LittleEndian.Uint32(data[:4]) ^ crc
		two := binary.LittleEndian.Uint32(data[4:8])
		crc = Lookup[0][(two>>24)&0xff] ^
			Lookup[1][(two>>16)&0xff] ^
			Lookup[2][(two>>8)&0xff] ^
			Lookup[3][two&0xff] ^
			Lookup[4][(one>>24)&0xff] ^
			Lookup[5][(one>>16)&0xff] ^
			Lookup[6][(one>>8)&0xff] ^
			Lookup[7][one&0xff]
		data = data[8:]
	}
	for _, c := range data {
		crc = (crc >> 8) ^ Lookup[0][(crc&0xff)^uint32(c)]
	}
	return crc ^ 0xffffffff
}

func Bytes4x8(data []byte, prevCrc32 uint32) uint32 {
	var crc = prevCrc32 ^ 0xffffffff

	unroll := 4
	bytesAtOnce := 8 * unroll

	for len(data) >= bytesAtOnce {
		for u := 0; u < unroll; u++ {
			one := binary.LittleEndian.Uint32(data[0:4]) ^ uint32(crc)
			two := binary.LittleEndian.Uint32(data[4:8])
			crc = Lookup[0][(two>>24)&0xff] ^
				Lookup[1][(two>>16)&0xff] ^
				Lookup[2][(two>>8)&0xff] ^
				Lookup[3][two&0xff] ^
				Lookup[4][(one>>24)&0xff] ^
				Lookup[5][(one>>16)&0xff] ^
				Lookup[6][(one>>8)&0xff] ^
				Lookup[7][one&0xff]
			data = data[8:]
		}
	}
	for _, c := range data {
		crc = (crc >> 8) ^ Lookup[0][(crc&0xff)^uint32(c)]
	}
	return crc ^ 0xffffffff
}

func Bytes16(data []byte, prevCrc32 uint32) uint32 {
	var crc = prevCrc32 ^ 0xffffffff

	unroll := 4
	bytesAtOnce := 16 * unroll

	for len(data) >= bytesAtOnce {
		for u := 0; u < unroll; u++ {
			one := binary.LittleEndian.Uint32(data[0:4]) ^ uint32(crc)
			two := binary.LittleEndian.Uint32(data[4:8])
			three := binary.LittleEndian.Uint32(data[8:12])
			four := binary.LittleEndian.Uint32(data[12:16])
			crc = Lookup[0][(four>>24)&0xff] ^
				Lookup[1][(four>>16)&0xff] ^
				Lookup[2][(four>>8)&0xff] ^
				Lookup[3][four&0xff] ^
				Lookup[4][(three>>24)&0xff] ^
				Lookup[5][(three>>16)&0xff] ^
				Lookup[6][(three>>8)&0xff] ^
				Lookup[7][three&0xff] ^
				Lookup[8][(two>>24)&0xff] ^
				Lookup[9][(two>>16)&0xff] ^
				Lookup[10][(two>>8)&0xff] ^
				Lookup[11][two&0xff] ^
				Lookup[12][(one>>24)&0xff] ^
				Lookup[13][(one>>16)&0xff] ^
				Lookup[14][(one>>8)&0xff] ^
				Lookup[15][one&0xff]
			data = data[16:]
		}
	}
	for _, c := range data {
		crc = (crc >> 8) ^ Lookup[0][(crc&0xff)^uint32(c)]
	}
	return crc ^ 0xffffffff
}
