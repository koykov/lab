package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Point3D struct {
	X, Y, Z float64
}

func (p *Point3D) DistanceTo(other *Point3D) float64 {
	dx := p.X - other.X
	dy := p.Y - other.Y
	dz := p.Z - other.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func GeneratePointsInSphere(n int, rng *rand.Rand) []Point3D {
	points := make([]Point3D, 0, n)

	for len(points) < n {

		x := 2*rng.Float64() - 1
		y := 2*rng.Float64() - 1
		z := 2*rng.Float64() - 1

		if x*x+y*y+z*z <= 1.0 {
			points = append(points, Point3D{X: x, Y: y, Z: z})
		}
	}
	return points
}

func FindMinDistances(points []Point3D) []float64 {
	minDistances := make([]float64, len(points))

	for i := range points {
		minDist := math.MaxFloat64

		for j := range points {
			if i == j {
				continue
			}

			dist := points[i].DistanceTo(&points[j])
			if dist < minDist {
				minDist = dist
			}
		}

		minDistances[i] = minDist
	}

	return minDistances
}

func KSTest(distances []float64, lambda float64) (float64, float64) {

	sorted := make([]float64, len(distances))
	copy(sorted, distances)
	sort.Float64s(sorted)

	n := float64(len(sorted))
	maxDiff := 0.0

	theoreticalCDF := func(x float64) float64 {
		return 1 - math.Exp(-lambda*x)
	}

	for i, x := range sorted {
		empirical := float64(i+1) / n
		theoretical := theoreticalCDF(x)
		diff := math.Abs(empirical - theoretical)

		if diff > maxDiff {
			maxDiff = diff
		}
	}

	criticalValue := 1.36 / math.Sqrt(n)

	return maxDiff, criticalValue
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numPoints := 10000
	numTrials := 10

	for trial := 0; trial < numTrials; trial++ {
		points := GeneratePointsInSphere(numPoints, rand.New(rand.NewSource(rand.Int63())))
		minDistances := FindMinDistances(points)
		density := float64(numPoints) / (4.0 * math.Pi / 3.0)
		lambda := math.Pow(4.0*math.Pi*density/3.0, 1.0/3.0)

		ksStat, criticalValue := KSTest(minDistances, lambda)

		fmt.Printf("stage %d:\n", trial+1)
		fmt.Printf("  KS stats: %.6f\n", ksStat)
		fmt.Printf("  criticl value (α=0.05): %.6f\n", criticalValue)

		if ksStat < criticalValue {
			fmt.Printf("  test passed\n")
		} else {
			fmt.Printf("  test failed\n")
		}

		meanDist := 0.0
		for _, d := range minDistances {
			meanDist += d
		}
		meanDist /= float64(len(minDistances))

		theoreticalMean := 1.0 / lambda
		fmt.Printf("  avg distance: %.6f (expected: ~%.6f)\n",
			meanDist, theoreticalMean)
		fmt.Println()
	}
}
