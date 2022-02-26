package main

import "fmt"

func volumesilinder(jarijari float64, tinggi float64) float64 {
	var hasil float64
	hasil = 3.14 * (jarijari * jarijari) * (tinggi)
	return hasil
}

func main() {
	/*
		atau
		var jarijari,tinggi int
		fmt.Scan(&jarijari, &tinggi)

	*/
	var jarijari int
	var tinggi int
	var total float64
	fmt.Scan(&jarijari)
	fmt.Scan(&tinggi)
	total = volumesilinder(float64(jarijari), float64(tinggi))
	fmt.Println(total)
}
