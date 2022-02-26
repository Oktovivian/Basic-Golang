package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	if d > a && d > b && d > c {
		fmt.Println("Ada rekor baru")
	} else {
		fmt.Println("Tidak ada rekor baru")
	}
}
