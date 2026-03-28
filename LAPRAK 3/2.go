package main

import (
	"fmt"
)

func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0
	var waktu int
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu < 301 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama, pemenang string
	var soal, skor int

	maxSoal := -1
	minSkor := 9999999 

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		hitungSkor(&soal, &skor)

		if soal > maxSoal {
			maxSoal = soal
			minSkor = skor
			pemenang = nama
		} else if soal == maxSoal {
			if skor < minSkor {
				minSkor = skor
				pemenang = nama
			}
		}
	}

	fmt.Println(pemenang, maxSoal, minSkor)
}