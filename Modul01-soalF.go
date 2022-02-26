package main

import "fmt"

/* MODUL 01 SOAL F */

func main() {
	var nilai int
	var jumlah = 0
	var n = 0
	var rata2 float64
	fmt.Scan(&nilai)
	for nilai != -1 {
		n += 1
		jumlah += nilai
		fmt.Scan(&nilai)

	}
	if n == 0 {
		rata2 = 0
	} else {
		rata2 = float64(jumlah) / float64(n)

	}
	fmt.Println(rata2)
}
