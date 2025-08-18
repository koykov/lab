package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	lotSize = 100.0
	carSize = 1.0
	maxCars = 10000
)

type Car struct {
	x, y float64
}

func parkingLotTest(rng *rand.Rand) (filledPercent float64, parkedCars int) {
	cars := make([]Car, 0, maxCars)

	for i := 0; i < maxCars; i++ {

		x := rng.Float64() * (lotSize - carSize)
		y := rng.Float64() * (lotSize - carSize)

		canPark := true
		for _, c := range cars {
			if x < c.x+carSize && x+carSize > c.x &&
				y < c.y+carSize && y+carSize > c.y {
				canPark = false
				break
			}
		}

		if canPark {
			cars = append(cars, Car{x, y})
			parkedCars++
		}
	}

	filledPercent = float64(parkedCars) * (carSize * carSize) / (lotSize * lotSize) * 100
	return
}

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	filled, cars := parkingLotTest(rng)

	fmt.Printf("parking cars: %d\n", cars)
	fmt.Printf("filled area: %.2f%%\n", filled)

	if filled >= 70.0 && filled <= 74.0 {
		fmt.Println("test passed")
	} else {
		fmt.Println("test failed")
	}
}
