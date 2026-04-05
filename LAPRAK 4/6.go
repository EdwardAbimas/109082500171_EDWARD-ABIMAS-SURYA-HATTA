package main

import "fmt"

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Print("Masukan (x y): ")
	fmt.Scan(&x, &y)
	fmt.Printf("Keluaran: %d\n", pangkat(x, y))
}