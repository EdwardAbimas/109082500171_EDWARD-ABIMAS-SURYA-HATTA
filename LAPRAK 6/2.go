package main

import (
	"fmt"
)

func main() {
	var x, y int
	var weights [1000]float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&weights[i])
	}

	var containerWeights []float64
	currentWeight := 0.0
	count := 0

	for i := 0; i < x; i++ {
		currentWeight += weights[i]
		count++

		if count == y || i == x-1 {
			containerWeights = append(containerWeights, currentWeight)
			currentWeight = 0.0
			count = 0
		}
	}

	for i, w := range containerWeights {
		fmt.Printf("%.2f", w)
		if i < len(containerWeights)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	totalAllWeight := 0.0
	for _, w := range containerWeights {
		totalAllWeight += w
	}
	
	if len(containerWeights) > 0 {
		avg := totalAllWeight / float64(len(containerWeights))
		fmt.Printf("%.2f\n", avg)
	}
}