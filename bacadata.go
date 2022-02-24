package main

import "fmt"

func main() {
	var teks string
	fmt.Scan(&teks)
	for teks != "SELESAI" {
		fmt.Print(teks, " ")
		fmt.Scan(&teks)
	}
}
