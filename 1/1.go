package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y float64
}

type Lingkaran struct {
	xPusat, yPusat, jariJari float64
}

func jarak(titik1, titik2 Titik) float64 {
	return math.Sqrt(math.Pow(titik1.x-titik2.x, 2) + math.Pow(titik1.y-titik2.y, 2))
}

func diDalam(lingkaran Lingkaran, titik Titik) bool {
	return jarak(Titik{lingkaran.xPusat, lingkaran.yPusat}, titik) <= lingkaran.jariJari
}

func main() {
	var lingkaran1, lingkaran2 Lingkaran
	var titik Titik

	fmt.Println("Masukkan koordinat pusat (xPusat, yPusat) dan jari-jari untuk Lingkaran 1:")
	fmt.Scan(&lingkaran1.xPusat, &lingkaran1.yPusat, &lingkaran1.jariJari)

	fmt.Println("Masukkan koordinat pusat (xPusat, yPusat) dan jari-jari untuk Lingkaran 2:")
	fmt.Scan(&lingkaran2.xPusat, &lingkaran2.yPusat, &lingkaran2.jariJari)

	fmt.Println("Masukkan koordinat (x, y) untuk Titik:")
	fmt.Scan(&titik.x, &titik.y)

	diLingkaran1 := diDalam(lingkaran1, titik)
	diLingkaran2 := diDalam(lingkaran2, titik)

	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik di dalam Lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik di dalam Lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("Titik di dalam Lingkaran 2")
	} else {
		fmt.Println("Titik di luar Lingkaran 1 dan 2")
	}
}