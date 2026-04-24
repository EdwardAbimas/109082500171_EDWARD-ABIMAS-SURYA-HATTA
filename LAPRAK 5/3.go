package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang []string

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	pertandingan := 1
	for {
		fmt.Printf("Pertandingan %d: ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorB > skorA {
			pemenang = append(pemenang, klubB)
		} else {
			pemenang = append(pemenang, "Draw")
		}
		pertandingan++
	}

	fmt.Println()
	for i := 0; i < len(pemenang); i++ {
		fmt.Printf("Hasil %d: %s\n", i+1, pemenang[i])
	}
	fmt.Println("Pertandingan selesai")
}