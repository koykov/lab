package main

import (
	"testing"
	"unsafe"

	"github.com/koykov/hash/fnv"
)

var (
	mapKS = map[string]int{}
	mapKH = map[uint64]int{}

	keyPool = []string{
		"27bf0ea8-a3c9-49bb-b55d-3305642955fe",
		"488bd1ca-69be-4b19-aae8-c55384206fec",
		"67093d89-864d-4723-b3a9-1b1b162321bd",
		"d04e60af-e367-416d-ba01-6bd2be302794",
		"10e04f8f-35b4-4259-b45a-cc89cd22585e",
		"3b5e14a5-36f0-4d7f-8f8b-25fa30244a0d",
		"d6a6912c-7ee8-4f13-ae4b-933cb6286d73",
		"e9fd7e90-9d9f-4aea-b77d-a3127a2c0cf2",
		"5eda17fb-a780-4ef3-a92b-cfae11267d4b",
		"9041dbcf-5a48-4243-8f92-c25532246c25",
		"a2c427bd-4f2a-437c-bf31-3fce7836897f",
		"10abe074-3165-4c91-ace7-e352673f3cd4",
		"5338ffff-6aa4-46fc-8cc2-bfdf7430e453",
		"f52f6528-8b30-4d32-b0ed-c6e3972d73ac",
		"60e1fa94-4fdc-414f-8be4-84039e2294a1",
		"14935f02-eed2-4891-bb4b-167e21209579",
		"6bf30828-57cb-4c9d-a8b8-07ab362114ed",
		"764bdcb6-27aa-4f1a-9149-8223cc2f4127",
	}
)

func init() {
	for i := 0; i < len(keyPool); i++ {
		key := keyPool[i]
		mapKS[key] = i
		mapKH[fnv.Hash64String(key)] = i
	}
}

func BenchmarkKeyStr(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if j := mapKS["67093d89-864d-4723-b3a9-1b1b162321bd"]; j != 2 {
			b.Error("key index mismatch")
		}
	}
}

func BenchmarkKeyHashStr(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		key := "67093d89-864d-4723-b3a9-1b1b162321bd"
		hptr := hashstr(uintptr(unsafe.Pointer(&key)))
		h := *(*uint64)(unsafe.Pointer(hptr))
		if j := mapKH[h]; j != 2 {
			b.Error("key index mismatch")
		}
	}
}

func BenchmarkKeyHash(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if j := mapKH[fnv.Hash64String("67093d89-864d-4723-b3a9-1b1b162321bd")]; j != 2 {
			b.Error("key index mismatch")
		}
	}
}
