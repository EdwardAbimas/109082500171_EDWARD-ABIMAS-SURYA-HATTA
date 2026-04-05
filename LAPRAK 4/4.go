package main

import "fmt"

func barisanMajuMundur(n int) {
	if n == 1 {
		fmt.Print(n, " ")
		return
	}
	fmt.Print(n, " ")
	barisanMajuMundur(n - 1)
	fmt.Print(n, " ")
}

func main() {
	var n int
	fmt.Print("Masukan: ")
	fmt.Scan(&n)
	fmt.Print("Keluaran: ")
	barisanMajuMundur(n)
	fmt.Println()
}