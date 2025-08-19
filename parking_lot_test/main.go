package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	lotSize     = 100.0  // Размер стоянки
	carSize     = 1.0    // Размер машины
	maxAttempts = 500000 // Количество попыток парковки
)

type Point struct {
	x, y float64
}

func parkingLotTest(rng *rand.Rand) (filledPercent float64, parkedCars int) {
	cars := make([]Point, 0, 10000)

	for attempt := 0; attempt < maxAttempts; attempt++ {
		x := rng.Float64() * (lotSize - carSize)
		y := rng.Float64() * (lotSize - carSize)

		// Проверяем, можно ли припарковаться
		canPark := true
		for _, car := range cars {
			// Расстояние между центрами должно быть >= carSize
			dx := math.Abs((x + carSize/2) - (car.x + carSize/2))
			dy := math.Abs((y + carSize/2) - (car.y + carSize/2))

			if dx < carSize && dy < carSize {
				canPark = false
				break
			}
		}

		if canPark {
			cars = append(cars, Point{x, y})
			parkedCars++
		}
	}

	filledPercent = float64(parkedCars) * (carSize * carSize) / (lotSize * lotSize) * 100
	return
}

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Запускаем несколько раз для усреднения
	totalFilled := 0.0
	totalCars := 0
	runs := 10

	for i := 0; i < runs; i++ {
		filled, cars := parkingLotTest(rng)
		totalFilled += filled
		totalCars += cars
		fmt.Printf("run %d: %.2f%% (%d cars)\n", i+1, filled, cars)
	}

	avgFilled := totalFilled / float64(runs)
	avgCars := totalCars / runs

	fmt.Printf("\navg: %.2f%% (%d cars)\n", avgFilled, avgCars)

	if avgFilled >= 68.0 && avgFilled <= 76.0 {
		fmt.Println("test passed")
	} else {
		fmt.Println("test failed")
	}
}
