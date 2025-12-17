package xtou

import (
	"strconv"
	"testing"
)

func TestXTOU(t *testing.T) {
	tests := []struct {
		input    string
		expected uint64
	}{
		{"DCAE", 56494},
		{"D835", 55349},
		{"0643", 1603},
		{"0000", 0},
		{"FFFF", 65535},
		{"1234", 4660},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			t.Run("native", func(t *testing.T) {
				result, _ := strconv.ParseUint(test.input, 16, 16)
				if result != test.expected {
					t.Errorf("generic(%s) = %d, expected %d",
						test.input, result, test.expected)
				}
			})
			t.Run("generic", func(t *testing.T) {
				result := xtouGeneric([]byte(test.input))
				if result != test.expected {
					t.Errorf("generic(%s) = %d, expected %d",
						test.input, result, test.expected)
				}
			})
			t.Run("table", func(t *testing.T) {
				result := xtouTable([]byte(test.input))
				if result != test.expected {
					t.Errorf("table(%s) = %d, expected %d",
						test.input, result, test.expected)
				}
			})
		})
	}
}

func BenchmarkXTOU(b *testing.B) {
	tests := []struct {
		input    string
		expected uint64
	}{
		{"DCAE", 56494},
		{"D835", 55349},
		{"0643", 1603},
		{"0000", 0},
		{"FFFF", 65535},
		{"1234", 4660},
	}
	for _, test := range tests {
		b.Run(test.input, func(b *testing.B) {
			b.Run("native", func(b *testing.B) {
				for range b.N {
					_, _ = strconv.ParseUint(test.input, 16, 16)
				}
			})
			b.Run("generic", func(b *testing.B) {
				x := []byte(test.input)
				for range b.N {
					xtouGeneric(x)
				}
			})
			b.Run("table", func(b *testing.B) {
				x := []byte(test.input)
				for range b.N {
					xtouTable(x)
				}
			})
		})
	}
}
