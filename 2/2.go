package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, indeksHapus, bilanganCari int
	var array [100]int 
	var total int       

	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	fmt.Println("\nIsi array:")
	for i := 0; i < n; i++ {
		fmt.Printf("array[%d] = %d\n", i, array[i])
	}

	fmt.Println("\nElemen dengan indeks ganjil:")
	for i := 1; i < n; i += 2 {
		fmt.Printf("array[%d] = %d\n", i, array[i])
	}

	fmt.Println("\nElemen dengan indeks genap:")
	for i := 0; i < n; i += 2 {
		fmt.Printf("array[%d] = %d\n", i, array[i])
	}

	fmt.Print("\nMasukkan nilai x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Printf("\nElemen dengan indeks kelipatan %d:\n", x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Printf("array[%d] = %d\n", i, array[i])
		}
	}

	for i := 0; i < n; i++ {
		total += array[i]
	}
	rataRata := float64(total) / float64(n)

	var jumlahKuadrat float64
	for i := 0; i < n; i++ {
		selisih := float64(array[i]) - rataRata
		jumlahKuadrat += selisih * selisih
	}
	stdDeviasi := math.Sqrt(jumlahKuadrat / float64(n))

	fmt.Printf("\nRata-rata bilangan dalam array sebelum elemen dihapus: %.2f\n", rataRata)
	fmt.Printf("Standar deviasi bilangan dalam array sebelum elemen dihapus: %.2f\n", stdDeviasi)

	fmt.Print("\nMasukkan indeks yang akan dihapus: ")
	fmt.Scan(&indeksHapus)
	if indeksHapus >= 0 && indeksHapus < n {
		fmt.Println("Array setelah elemen dihapus:")
		for i := 0; i < n; i++ {
			if i != indeksHapus {
				fmt.Printf("array[%d] = %d\n", i, array[i])
			}
		}
		n-- 
	} else {
		fmt.Println("Indeks tidak valid!")
	}

	fmt.Print("\nMasukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&bilanganCari)
	frekuensi := 0
	for i := 0; i < n; i++ {
		if array[i] == bilanganCari {
			frekuensi++
		}
	}
	fmt.Printf("Frekuensi bilangan %d dalam array: %d\n", bilanganCari, frekuensi)
}