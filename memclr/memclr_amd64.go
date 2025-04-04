package memclr

import "golang.org/x/sys/cpu"

var funcAMD64 func([]uint64)

func init() {
	// if cpu.X86.HasBMI2 && cpu.X86.HasAVX512BW {
	// 	funcAMD64 = countAVX512
	// 	return
	// }
	// if cpu.X86.HasBMI2 && cpu.X86.HasAVX2 {
	// 	funcAMD64 = memclr64AVX2
	// 	return
	// }
	if cpu.X86.HasSSE2 {
		funcAMD64 = memclr64SSE2
		return
	}
	funcAMD64 = memclr64generic
}

func memclr64(data []uint64) {
	funcAMD64(data)
}

//go:noescape
func memclr64SSE2(p []uint64)

//go:noescape
func memclr64AVX2(p []byte)
