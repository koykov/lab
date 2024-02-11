package main

import "testing"

type mapL1 map[int32]struct{}
type mapL2 map[int32]map[int32]struct{}
type mapL3 map[int32]map[int32]map[int32]struct{}

var (
	str = map[string]string{"foo": "bar"}
	l1  = mapL1{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}
	l2  = mapL2{
		1: {1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		2: {1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		3: {1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
	}
	l3 = mapL3{
		1: {
			1: {1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			2: {1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			3: {1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
		2: {
			1: {1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			2: {1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			3: {1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
		3: {
			1: {1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			2: {1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
			3: {1: struct{}{}, 2: struct{}{}, 3: struct{}{}},
		},
	}
)

func BenchmarkAccessStr(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, ok := str["foo"]; !ok {
			b.Error("str: not found")
		}
	}
}

func BenchmarkAccessL1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, ok := l1[2]; !ok {
			b.Error("l1: not found")
		}
	}
}

func BenchmarkAccessL2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, ok := l2[2][2]; !ok {
			b.Error("l2: not found")
		}
	}
}

func BenchmarkAccessL3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, ok := l3[2][2][2]; !ok {
			b.Error("l3: not found")
		}
	}
}
