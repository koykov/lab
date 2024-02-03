package loop_opt

import (
	"fmt"
	"testing"
)

func BenchmarkBoundsCheckElimination(b *testing.B) {
	stages := [][]int{
		make([]int, 10),
		make([]int, 100),
		make([]int, 1000),
		make([]int, 10000),
		make([]int, 100000),
		make([]int, 1000000),
		make([]int, 10000000),
		make([]int, 100000000),
		make([]int, 1000000000),
	}
	for i := 0; i < len(stages); i++ {
		stage := stages[i]
		b.Run(fmt.Sprintf("on/%d", len(stage)), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				var x int
				_ = stage[len(stage)-1]
				for k := 0; k < len(stage); k++ {
					x += stage[k]
				}
				_ = x
			}
		})
		b.Run(fmt.Sprintf("off/%d", len(stage)), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				var x int
				for k := 0; k < len(stage); k++ {
					x += stage[k]
				}
				_ = x
			}
		})
	}
}
