package main

import (
	"fmt"
	"math"
	"math/rand"
)

func BinaryMatrixRankTest(bits []int, matrixSize, numMatrices int) float64 {
	totalBitsNeeded := matrixSize * matrixSize * numMatrices
	if len(bits) < totalBitsNeeded {
		panic("Недостаточно бит для анализа")
	}

	matrices := make([][][]int, numMatrices)
	for i := 0; i < numMatrices; i++ {
		matrix := make([][]int, matrixSize)
		for row := 0; row < matrixSize; row++ {
			start := i*matrixSize*matrixSize + row*matrixSize
			end := start + matrixSize
			matrix[row] = bits[start:end]
		}
		matrices[i] = matrix
	}

	rankCounts := make(map[int]int)
	for _, matrix := range matrices {
		rank := computeMatrixRank(matrix)
		rankCounts[rank]++
	}

	expectedProbs := map[int]float64{}

	observed := make([]float64, 3)
	expected := make([]float64, 3)
	for rank, count := range rankCounts {
		if rank == matrixSize {
			observed[0] += float64(count)
		} else if rank == matrixSize-1 {
			observed[1] += float64(count)
		} else {
			observed[2] += float64(count)
		}
	}

	for i, prob := range expectedProbs {
		if i == matrixSize {
			expected[0] = prob * float64(numMatrices)
		} else if i == matrixSize-1 {
			expected[1] = prob * float64(numMatrices)
		} else {
			expected[2] += prob * float64(numMatrices)
		}
	}

	chi2 := 0.0
	for i := 0; i < 3; i++ {
		if expected[i] != 0 {
			chi2 += math.Pow(observed[i]-expected[i], 2) / expected[i]
		}
	}

	pValue := math.Exp(-chi2 / 2)
	return pValue
}

func computeMatrixRank(matrix [][]int) int {
	n := len(matrix)
	rank := 0
	row := make([]int, n)
	copy(row, matrix[0])

	for i := 0; i < n; i++ {
		pivotRow := -1
		for j := i; j < n; j++ {
			if matrix[j][i] == 1 {
				pivotRow = j
				break
			}
		}

		if pivotRow == -1 {
			continue
		}

		matrix[i], matrix[pivotRow] = matrix[pivotRow], matrix[i]
		rank++

		for j := i + 1; j < n; j++ {
			if matrix[j][i] == 1 {
				for k := i; k < n; k++ {
					matrix[j][k] ^= matrix[i][k]
				}
			}
		}
	}

	return rank
}

func main() {
	bits := make([]int, 32*32*100)
	for i := range bits {
		bits[i] = rand.Intn(2)
	}

	pValue := BinaryMatrixRankTest(bits, 32, 100)
	fmt.Printf("p-value: %.4f\n", pValue)

	if pValue >= 0.05 {
		fmt.Println("test passed")
	} else {
		fmt.Println("test failed")
	}
}
