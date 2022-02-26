package main

import "fmt"

func main() {
	var kata string
	var a, b int
	var penjumlahan int

	fmt.Scan(&kata)
	fmt.Println("kata =", kata)

	fmt.Scan(&a, &b)
	penjumlahan = a + b
	fmt.Println("jumlah =", penjumlahan)
}
