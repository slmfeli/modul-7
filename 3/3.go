package main

import "fmt"

func main() {
	var namaKlubA, namaKlubB string

	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scanln(&namaKlubA)
	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scanln(&namaKlubB)

	var pemenang [100]string
	var jumlahPertandingan int

	pertandingan := 1
	for {
		var skorA, skorB int

		fmt.Printf("Pertandingan %d (Masukkan SkorA SkorB): ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang[jumlahPertandingan] = namaKlubA
		} else if skorA < skorB {
			pemenang[jumlahPertandingan] = namaKlubB
		} else {
			pemenang[jumlahPertandingan] = "Draw"
		}

		jumlahPertandingan++
		pertandingan++
	}

	fmt.Println("\nHasil pertandingan:")
	for i := 0; i < jumlahPertandingan; i++ {
		fmt.Printf("Pertandingan %d: %s\n", i+1, pemenang[i])
	}
}