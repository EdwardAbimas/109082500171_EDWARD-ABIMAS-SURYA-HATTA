package main

import (
	"fmt"
	"math"
)

const NMAX = 100
type arrInt [NMAX]int

func main() {
	var data arrInt
	var n, x, hapusIdx int

	fmt.Print("Masukkan jumlah elemen (N): ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	fmt.Print("a. Isi array: ")
	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	fmt.Print("b. Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	fmt.Print("c. Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("d. Indeks kelipatan ", x, ": ")
	if x > 0 {
		for i := x; i < n; i += x {
			fmt.Print(data[i], " ")
		}
	}
	fmt.Println()

	fmt.Print("Masukkan indeks yang dihapus: ")
	fmt.Scan(&hapusIdx)
	if hapusIdx >= 0 && hapusIdx < n {
		for i := hapusIdx; i < n-1; i++ {
			data[i] = data[i+1]
		}
		n--
	}
	
	fmt.Print("e. Array setelah dihapus: ")
	for i := 0; i < n; i++ {
		fmt.Print(data[i], " ")
	}
	fmt.Println()

	sum := 0
	for i := 0; i < n; i++ {
		sum += data[i]
	}
	rata := float64(sum) / float64(n)
	fmt.Printf("f. Rata-rata: %.2f\n", rata)

	var sumVar float64
	for i := 0; i < n; i++ {
		sumVar += math.Pow(float64(data[i])-rata, 2)
	}
	stdDev := math.Sqrt(sumVar / float64(n))
	fmt.Printf("g. Standar Deviasi: %.2f\n", stdDev)

	var cari, frekuensi int
	fmt.Print("Masukkan bilangan untuk dicari frekuensinya: ")
	fmt.Scan(&cari)
	for i := 0; i < n; i++ {
		if data[i] == cari {
			frekuensi++
		}
	}
	fmt.Printf("h. Frekuensi bilangan %d adalah: %d\n", cari, frekuensi)
}