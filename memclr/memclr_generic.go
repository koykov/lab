//go:build (!amd64 && !arm64 && !ppc64le && !riscv64) || appengine || !gc || purego

package memclr

func memcrl64(p []uint64) {
	memclr64generic(p)
}
