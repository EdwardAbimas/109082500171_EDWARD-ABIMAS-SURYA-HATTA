package main

import (
	"fmt"
)

const NMAX int = 127
type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input string
	*n = 0
	for {
		fmt.Scan(&input)
		if input == "." || *n >= NMAX {
			break
		}
		(*t)[*n] = rune(input[0])
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		temp := (*t)[i]
		(*t)[i] = (*t)[n-1-i]
		(*t)[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	var tempTabel tabel = t
	balikanArray(&tempTabel, n)
	
	for i := 0; i < n; i++ {
		if t[i] != tempTabel[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks\n: ")
	isiArray(&tab, &m)

	fmt.Print("Reverse teks: ")
	var reversedTab = tab
	balikanArray(&reversedTab, m)
	cetakArray(reversedTab, m)

	isPalindrom := palindrom(tab, m)
	fmt.Printf("Palindrom ? %t\n", isPalindrom)
}