package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var karakter rune
	fmt.Println("Masukkan karakter (akhiri dengan '.'): ")
	for {
		fmt.Scanf("%c", &karakter)
		if karakter == '.' || *n >= NMAX {
			break
		}
		t[*n] = karakter
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikkanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func cekPalindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var jumlahKarakter int

	fmt.Println("Teks:")
	isiArray(&tab, &jumlahKarakter)

	fmt.Print("Reverse teks: ")
	cetakArray(tab, jumlahKarakter)

	balikkanArray(&tab, jumlahKarakter)

	fmt.Print("Teks setelah dibalik: ")
	cetakArray(tab, jumlahKarakter)

	isPalindrom := cekPalindrom(tab, jumlahKarakter)
	fmt.Print("Palindrom? ")
	if isPalindrom {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}