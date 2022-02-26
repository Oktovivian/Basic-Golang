package main

import "fmt"

/* MODUL 01 SOAL C */

func main() {
	var x int
	fmt.Scan(&x)
	if x%3 == 0 {
		fmt.Println("Fizz")
	}
	if x%5 == 0 {
		fmt.Println("Buzz")
	}
}
