package main

import "fmt"

func main() {
	var n, t1, t2, t3 int
	var jumlah = 0
	var penjumlahan int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t1, &t2, &t3)
		penjumlahan = t1 + t2 + t3
		if penjumlahan >= 2 {
			jumlah += 1
		}
	}
	fmt.Println(jumlah)

}
