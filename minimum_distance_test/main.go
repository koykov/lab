package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Point []float64

func GeneratePoints(n, d int, rng *rand.Rand) []Point {
	points := make([]Point, n)
	for i := range points {
		points[i] = make(Point, d)
		for j := range points[i] {
			points[i][j] = rng.Float64()
		}
	}
	return points
}

func EuclideanDistance(p1, p2 Point) float64 {
	sum := 0.0
	for i := range p1 {
		diff := p1[i] - p2[i]
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

func FindMinDistances(points []Point) []float64 {
	n := len(points)
	minDistances := make([]float64, n)

	for i := 0; i < n; i++ {
		minDist := math.MaxFloat64
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			dist := EuclideanDistance(points[i], points[j])
			if dist < minDist {
				minDist = dist
			}
		}
		minDistances[i] = minDist
	}

	return minDistances
}

func KSTest(minDistances []float64, lambda float64) float64 {
	n := float64(len(minDistances))
	sorted := make([]float64, len(minDistances))
	copy(sorted, minDistances)
	sort.Float64s(sorted)

	maxDiff := 0.0
	for i, x := range sorted {

		empirical := float64(i+1) / n

		theoretical := 1 - math.Exp(-lambda*x)

		diff := math.Abs(empirical - theoretical)
		if diff > maxDiff {
			maxDiff = diff
		}
	}

	return maxDiff
}

func main() {
	rand.Seed(time.Now().UnixNano())
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	n := 10000
	d := 3

	points := GeneratePoints(n, d, rng)

	minDistances := FindMinDistances(points)

	meanDist := 0.0
	for _, dist := range minDistances {
		meanDist += dist
	}
	meanDist /= float64(len(minDistances))
	lambda := 1.0 / meanDist

	ksStat := KSTest(minDistances, lambda)

	criticalValue := 1.36 / math.Sqrt(float64(n))

	fmt.Printf("Minimum distance test (%d points, %dD)\n", n, d)
	fmt.Printf("Avg minimal distance: %.6f\n", meanDist)
	fmt.Printf("Lambda: %.6f\n", lambda)
	fmt.Printf("KS-statistics: %.6f\n", ksStat)
	fmt.Printf("Critical value (α=0.05): %.6f\n", criticalValue)

	if ksStat < criticalValue {
		fmt.Println("test passed")
	} else {
		fmt.Println("test failed")
	}

	fmt.Println("\nHistogram of minimal distance:")
	PrintHistogram(minDistances, 20)
}

func PrintHistogram(data []float64, bins int) {
	minVal, maxVal := data[0], data[0]
	for _, val := range data {
		if val < minVal {
			minVal = val
		}
		if val > maxVal {
			maxVal = val
		}
	}

	binWidth := (maxVal - minVal) / float64(bins)
	histogram := make([]int, bins)

	for _, val := range data {
		binIdx := int((val - minVal) / binWidth)
		if binIdx >= bins {
			binIdx = bins - 1
		}
		histogram[binIdx]++
	}

	maxCount := 0
	for _, count := range histogram {
		if count > maxCount {
			maxCount = count
		}
	}

	scale := 50.0 / float64(maxCount)
	for i, count := range histogram {
		lower := minVal + float64(i)*binWidth
		upper := lower + binWidth
		bars := int(float64(count) * scale)

		fmt.Printf("[%.4f-%.4f] ", lower, upper)
		for j := 0; j < bars; j++ {
			fmt.Print("█")
		}
		fmt.Printf(" %d\n", count)
	}
}
