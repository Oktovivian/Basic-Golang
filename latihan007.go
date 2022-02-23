package main

import "fmt"

func hitungluaslingkaran(jari2 float64) float64 {
	var hasil float64
	hasil = 3.14 * (jari2 * jari2)
	return hasil
}

func main() {
	var jari2 float64
	var total float64
	fmt.Scan(&jari2)
	total = hitungluaslingkaran(jari2)
	fmt.Println(total)
}
