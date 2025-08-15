package main

import "fmt"

type BinaryMatrix struct {
	A [][]int
	M int
	Q int
	m int
}

func NewBinaryMatrix(matrix [][]int, rows, cols int) *BinaryMatrix {
	m := rows
	if cols < rows {
		m = cols
	}
	return &BinaryMatrix{
		A: matrix,
		M: rows,
		Q: cols,
		m: m,
	}
}

func (bm *BinaryMatrix) ComputeRank(verbose bool) int {
	if verbose {
		fmt.Println("Original Matrix:")
		bm.printMatrix()
	}

	i := 0
	for i < bm.m-1 {
		if bm.A[i][i] == 1 {
			bm.performRowOperations(i, true)
		} else {
			found := bm.findUnitElementSwap(i, true)
			if found == 1 {
				bm.performRowOperations(i, true)
			}
		}
		i++
	}

	if verbose {
		fmt.Println("Intermediate Matrix:")
		bm.printMatrix()
	}

	i = bm.m - 1
	for i > 0 {
		if bm.A[i][i] == 1 {
			bm.performRowOperations(i, false)
		} else {
			if bm.findUnitElementSwap(i, false) == 1 {
				bm.performRowOperations(i, false)
			}
		}
		i--
	}

	if verbose {
		fmt.Println("Final Matrix:")
		bm.printMatrix()
	}

	return bm.determineRank()
}

func (bm *BinaryMatrix) performRowOperations(i int, forwardElimination bool) {
	if forwardElimination {
		for j := i + 1; j < bm.M; j++ {
			if bm.A[j][i] == 1 {
				for k := 0; k < bm.Q; k++ {
					bm.A[j][k] = (bm.A[j][k] + bm.A[i][k]) % 2
				}
			}
		}
	} else {
		for j := i - 1; j >= 0; j-- {
			if bm.A[j][i] == 1 {
				for k := 0; k < bm.Q; k++ {
					bm.A[j][k] = (bm.A[j][k] + bm.A[i][k]) % 2
				}
			}
		}
	}
}

func (bm *BinaryMatrix) findUnitElementSwap(i int, forwardElimination bool) int {
	rowOp := 0
	if forwardElimination {
		index := i + 1
		for index < bm.M && bm.A[index][i] == 0 {
			index++
		}
		if index < bm.M {
			rowOp = bm.swapRows(i, index)
		}
	} else {
		index := i - 1
		for index >= 0 && bm.A[index][i] == 0 {
			index--
		}
		if index >= 0 {
			rowOp = bm.swapRows(i, index)
		}
	}
	return rowOp
}

func (bm *BinaryMatrix) swapRows(i, ix int) int {
	bm.A[i], bm.A[ix] = bm.A[ix], bm.A[i]
	return 1
}

func (bm *BinaryMatrix) determineRank() int {
	rank := bm.m
	for i := 0; i < bm.M; i++ {
		allZeros := true
		for j := 0; j < bm.Q; j++ {
			if bm.A[i][j] == 1 {
				allZeros = false
				break
			}
		}
		if allZeros {
			rank--
		}
	}
	return rank
}

func (bm *BinaryMatrix) printMatrix() {
	for _, row := range bm.A {
		fmt.Println(row)
	}
}

func main() {
	matrix := [][]int{
		{1, 0, 1},
		{0, 1, 0},
		{1, 1, 1},
	}
	bm := NewBinaryMatrix(matrix, 3, 3)
	rank := bm.ComputeRank(true)
	fmt.Println("Rank of the matrix:", rank)
}
