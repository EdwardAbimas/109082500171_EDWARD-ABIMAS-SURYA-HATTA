package main

import "fmt"

func main() {
	var berat int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	kg := berat / 1000
	sisa := berat % 1000

	biayaKg := kg * 10000
	var biayaSisa int

	if sisa >= 500 {
		biayaSisa = sisa * 5
	} else {
		biayaSisa = sisa * 15
	}

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisa)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)

	totalBiaya := biayaKg

	if berat <= 10000 {
		totalBiaya += biayaSisa
	}

	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}