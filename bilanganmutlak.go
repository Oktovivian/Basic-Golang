package main

import "fmt"

func main() {
	var angka1, angka2 float64
	var nilai float64
	fmt.Scan(&angka1, &angka2)
	nilai = angka1 - angka2
	if nilai <= 0 {
		nilai = nilai * float64(-1)

		fmt.Println("nilainya adalah", nilai)
	}
}
