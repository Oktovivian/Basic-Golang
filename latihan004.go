package main

import "fmt"

// jika kita input tipe data beda dgn yg dideklar, maka outputnya 0
func main() {
	// input 1 nilai
	/*
		var a int
		fmt.Scan(&a)
		fmt.Println(a)
	*/

	// input banyak nilai
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Println(x, y, z)
}
