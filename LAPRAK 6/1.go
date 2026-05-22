package main

import (
	"fmt"
)

func main() {
	var weights [1000]float64
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&weights[i])
	}

	if n > 0 {
		min := weights[0]
		max := weights[0]

		for i := 1; i < n; i++ {
			if weights[i] < min {
				min = weights[i]
			}
			if weights[i] > max {
				max = weights[i]
			}
		}

		fmt.Printf("%.2f %.2f\n", min, max)
	}
}