package mapkeys_complex

import (
	"testing"
	"unsafe"
)

func TestMapkeys(t *testing.T) {
	t.Run("80", func(t *testing.T) {
		k := k80{k0: 45435, k1: 345}
		k1 := k80{k0: 111, k1: 222}
		t.Log(unsafe.Sizeof(k))
		t.Log(unsafe.Sizeof(uint64(1)))
		t.Log(unsafe.Sizeof(uint16(1)))
		m := map80{}
		m[k] = struct{}{}
		if _, ok := m[k]; !ok {
			t.FailNow()
		}
		if _, ok := m[k1]; ok {
			t.FailNow()
		}
	})
}
