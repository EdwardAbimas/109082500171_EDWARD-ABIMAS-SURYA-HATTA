package main

import (
	"fmt"
)

func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutation(n, r int, hasil *int) {
	var factN, factNR int
	factorial(n, &factN)
	factorial(n-r, &factNR)
	*hasil = factN / factNR
}

func combination(n, r int, hasil *int) {
	var factN, factR, factNR int
	factorial(n, &factN)
	factorial(r, &factR)
	factorial(n-r, &factNR)
	*hasil = factN / (factR * factNR)
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	var p1, c1, p2, c2 int

	permutation(a, c, &p1)
	combination(a, c, &c1)

	permutation(b, d, &p2)
	combination(b, d, &c2)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}