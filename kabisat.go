package main

import "fmt"

func main() {
	var tahun int
	var kabisat bool

	fmt.Scan(&tahun)
	kabisat = tahun%400 == 0 || tahun%4 == 0 && tahun%100 != 0
	fmt.Print(kabisat)
}
