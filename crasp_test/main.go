package main

import (
	"fmt"
	"math"
	"math/rand"
)

func CrapsTest(rng func() float64, games int) float64 {
	wins := 0

	roll := func() int {

		dice1 := int(rng()*6) + 1
		dice2 := int(rng()*6) + 1
		return dice1 + dice2
	}

	for i := 0; i < games; i++ {

		firstRoll := roll()

		switch firstRoll {
		case 7, 11:
			wins++
		case 2, 3, 12:

		default:

			point := firstRoll

			for {
				newRoll := roll()
				if newRoll == point {
					wins++
					break
				} else if newRoll == 7 {

					break
				}
			}
		}
	}

	expectedWinProb := 244.0 / 495.0

	expectedWins := expectedWinProb * float64(games)
	chi2 := math.Pow(float64(wins)-expectedWins, 2) / expectedWins

	pValue := math.Exp(-chi2 / 2)

	return pValue
}

func main() {
	games := 1_000_000
	pValue := CrapsTest(rand.Float64, games)

	fmt.Printf("games: %d\n", games)
	fmt.Printf("p-value: %.6f\n", pValue)

	if pValue < 0.01 {
		fmt.Println("test failed")
	} else if pValue < 0.05 {
		fmt.Println("edge result")
	} else {
		fmt.Println("test passed")
	}
}
