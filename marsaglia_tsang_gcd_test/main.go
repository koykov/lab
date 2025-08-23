package main

import (
	"fmt"
	"math"
	"math/rand"
)

func gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func marsagliaGCDTest(rng func() uint64, numPairs int) (float64, float64, float64) {
	countCoprime := 0

	for i := 0; i < numPairs; i++ {
		a := rng()
		b := rng()
		if gcd(a, b) == 1 {
			countCoprime++
		}
	}
	observedProb := float64(countCoprime) / float64(numPairs)
	expectedProb := 6.0 / (math.Pi * math.Pi)
	stdDev := math.Sqrt(expectedProb * (1 - expectedProb) / float64(numPairs))
	zScore := (observedProb - expectedProb) / stdDev
	return observedProb, expectedProb, zScore
}

func main() {
	rng := func() uint64 {
		return rand.Uint64()
	}

	numPairs := 1000000
	observed, expected, zScore := marsagliaGCDTest(rng, numPairs)

	fmt.Printf("Тест Marsaglia-Tsang GCD:\n")
	fmt.Printf("pairs number: %d\n", numPairs)
	fmt.Printf("GCD=1 possibility (expected): %.6f\n", expected)
	fmt.Printf("GCD=1 possibility (observed): %.6f\n", observed)
	fmt.Printf("Z-score: %.4f\n", zScore)

	// Интерпретация результата
	if math.Abs(zScore) < 3.0 {
		fmt.Println("test passed")
	} else if math.Abs(zScore) < 6.0 {
		fmt.Println("edge case")
	} else {
		fmt.Println("test failed")
	}
}
