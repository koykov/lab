package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type rng struct {
	state uint64
}

func (r *rng) Next() float64 {
	return rand.Float64()
}

func SqueezeTest(rng func() float64, samples int, multiplier float64) (float64, []float64) {
	squeezed := make([]float64, samples)

	for i := 0; i < samples; i++ {
		original := rng()
		squeezed[i] = math.Mod(original*multiplier, 1.0)
	}

	bins := 10
	expected := float64(samples) / float64(bins)
	observed := make([]float64, bins)

	for _, val := range squeezed {
		bin := int(math.Floor(val * float64(bins)))
		if bin >= bins {
			bin = bins - 1
		}
		observed[bin]++
	}

	chi2 := 0.0
	for _, obs := range observed {
		chi2 += math.Pow(obs-expected, 2) / expected
	}

	return chi2, squeezed
}

func printHistogram(data []float64, bins int) {
	counts := make([]int, bins)
	for _, val := range data {
		idx := int(math.Floor(val * float64(bins)))
		if idx >= bins {
			idx = bins - 1
		}
		counts[idx]++
	}

	fmt.Println("Histogram of squeezed values:")
	for i, count := range counts {
		bar := ""
		for j := 0; j < count/1000; j++ {
			bar += "█"
		}
		fmt.Printf("bin %d: %s (%d)\n", i, bar, count)
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())
	yourRNG := &rng{state: uint64(time.Now().UnixNano())}

	samples := 1000000
	multiplier := math.Pow(2, 32)

	chi2, squeezedValues := SqueezeTest(yourRNG.Next, samples, multiplier)

	fmt.Printf("Squeeze Test:\n")
	fmt.Printf("samples: %d\n", samples)
	fmt.Printf("multiplier: %.0f\n", multiplier)
	fmt.Printf("Chi-square statistics: %.4f\n", chi2)

	criticalValue := 16.919
	fmt.Printf("critical value (α=0.05, df=9): %.3f\n", criticalValue)

	if chi2 < criticalValue {
		fmt.Println("test passed")
	} else {
		fmt.Println("test failed")
	}
	printHistogram(squeezedValues, 10)

}
