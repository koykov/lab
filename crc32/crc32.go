package crc32

import (
	"encoding/binary"
)

const (
	Polynomial int64 = 0xEDB88320

	MaxSlice16    = 16
	MaxSlice8     = 8
	MaxSlice4     = 4
	MaxSlice1     = 1
	MaxSliceNoLut = 0
)

func Crc32Bitwise(data []byte, prevCrc32 uint32) uint32 {
	var crc = int64(prevCrc32) ^ 0xFFFFFFFF
	for _, c := range data {
		crc ^= int64(c)
		for j := 0; j < 8; j++ {
			crc = (crc >> 1) ^ (-int64(crc&1) & Polynomial)
		}
	}
	return uint32(crc ^ 0xFFFFFFFF)
}

func Crc32Halfbyte(data []byte, prevCrc32 uint32) uint32 {
	var crc = int64(prevCrc32) ^ 0xFFFFFFFF
	for _, c := range data {
		crc = int64(Lookup16[(uint32(crc)^uint32(c))&0x0f] ^ (uint32(crc) >> 4))
		crc = int64(Lookup16[(uint32(crc)^(uint32(c)>>4))&0x0f] ^ (uint32(crc) >> 4))
	}
	return uint32(crc ^ 0xFFFFFFFF)
}

func Crc32Byte1(data []byte, prevCrc32 uint32) uint32 {
	var crc = int64(prevCrc32) ^ 0xFFFFFFFF
	for _, c := range data {
		crc = (crc >> 8) ^ int64(Lookup[MaxSliceNoLut][(crc&0xFF)^int64(c)])
	}
	return uint32(crc ^ 0xFFFFFFFF)
}

func Crc32Byte1Tableless(data []byte, prevCrc32 uint32) uint32 {
	var crc = int64(prevCrc32) ^ 0xFFFFFFFF
	for _, c := range data {
		s := uint8(crc) ^ uint8(c)
		low := int64((s ^ (s << 6)) & 0xFF)
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
	return uint32(crc ^ 0xFFFFFFFF)
}

func Crc32Bytes4(data []byte, prevCrc32 uint32) uint32 {
	var crc = int64(prevCrc32) ^ 0xFFFFFFFF
	for len(data) >= 4 {
		one := binary.LittleEndian.Uint32(data[:4]) ^ uint32(crc)
		crc = int64(Lookup[0][(one>>24)&0xFF] ^
			Lookup[1][(one>>16)&0xFF] ^
			Lookup[2][(one>>8)&0xFF] ^
			Lookup[3][one&0xFF])
		data = data[4:]
	}
	for _, c := range data {
		crc = (crc >> 8) ^ int64(Lookup[0][(crc&0xFF)^int64(c)])
	}
	return uint32(crc ^ 0xFFFFFFFF)
}

func Crc32Bytes8(data []byte, prevCrc32 uint32) uint32 {
	var crc = uint32(prevCrc32 ^ 0xFFFFFFFF)
	for len(data) >= 8 {
		one := binary.LittleEndian.Uint32(data[:4]) ^ crc
		two := binary.LittleEndian.Uint32(data[4:8])
		crc = Lookup[0][(two>>24)&0xFF] ^
			Lookup[1][(two>>16)&0xFF] ^
			Lookup[2][(two>>8)&0xFF] ^
			Lookup[3][two&0xFF] ^
			Lookup[4][(one>>24)&0xFF] ^
			Lookup[5][(one>>16)&0xFF] ^
			Lookup[6][(one>>8)&0xFF] ^
			Lookup[7][one&0xFF]
		data = data[8:]
	}
	for _, c := range data {
		crc = (crc >> 8) ^ Lookup[0][(crc&0xFF)^uint32(c)]
	}
	return crc ^ 0xFFFFFFFF
}

func Crc32Bytes4x8(data []byte, prevCrc32 uint32) uint32 {
	var crc = int64(prevCrc32) ^ 0xFFFFFFFF

	unroll := 4
	bytesAtOnce := 8 * unroll

	for len(data) >= bytesAtOnce {
		for u := 0; u < unroll; u++ {
			one := binary.LittleEndian.Uint32(data[0:4]) ^ uint32(crc)
			two := binary.LittleEndian.Uint32(data[4:8])
			crc = int64(Lookup[0][(two>>24)&0xFF] ^
				Lookup[1][(two>>16)&0xFF] ^
				Lookup[2][(two>>8)&0xFF] ^
				Lookup[3][two&0xFF] ^
				Lookup[4][(one>>24)&0xFF] ^
				Lookup[5][(one>>16)&0xFF] ^
				Lookup[6][(one>>8)&0xFF] ^
				Lookup[7][one&0xFF])
			data = data[8:]
		}
	}
	for _, c := range data {
		crc = (crc >> 8) ^ int64(Lookup[0][(crc&0xFF)^int64(c)])
	}
	return uint32(crc ^ 0xFFFFFFFF)
}

func Crc32Bytes16(data []byte, prevCrc32 uint32) uint32 {
	var crc = int64(prevCrc32) ^ 0xFFFFFFFF

	unroll := 4
	bytesAtOnce := 16 * unroll

	for len(data) >= bytesAtOnce {
		for u := 0; u < unroll; u++ {
			one := binary.LittleEndian.Uint32(data[0:4]) ^ uint32(crc)
			two := binary.LittleEndian.Uint32(data[4:8])
			three := binary.LittleEndian.Uint32(data[8:12])
			four := binary.LittleEndian.Uint32(data[12:16])
			crc = int64(Lookup[0][(four>>24)&0xFF] ^
				Lookup[1][(four>>16)&0xFF] ^
				Lookup[2][(four>>8)&0xFF] ^
				Lookup[3][four&0xFF] ^
				Lookup[4][(three>>24)&0xFF] ^
				Lookup[5][(three>>16)&0xFF] ^
				Lookup[6][(three>>8)&0xFF] ^
				Lookup[7][three&0xFF] ^
				Lookup[8][(two>>24)&0xFF] ^
				Lookup[9][(two>>16)&0xFF] ^
				Lookup[10][(two>>8)&0xFF] ^
				Lookup[11][two&0xFF] ^
				Lookup[12][(one>>24)&0xFF] ^
				Lookup[13][(one>>16)&0xFF] ^
				Lookup[14][(one>>8)&0xFF] ^
				Lookup[15][one&0xFF])
			data = data[16:]
		}
	}
	for _, c := range data {
		crc = (crc >> 8) ^ int64(Lookup[0][(crc&0xFF)^int64(c)])
	}
	return uint32(crc ^ 0xFFFFFFFF)
}
