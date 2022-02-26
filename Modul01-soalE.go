package main

/* MODUL 01 SOAL E */
import "fmt"

func main() {
	var teks string
	fmt.Scan(&teks)
	for teks != "SELESAI" {
		fmt.Println(teks)
		fmt.Scan(&teks)
	}
}
