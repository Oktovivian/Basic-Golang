package main

import "fmt"

func main() {
	var x int
	var a, b, c int
	var d, e, f int
	fmt.Scan(&a, &b, &c)
	fmt.Scan(&d, &e, &f)
	fmt.Scan(&x)

	if x == 0 || x == 360 {
		fmt.Println(a, b, c)
		fmt.Println(d, e, f)
	} else if x == 90 {
		fmt.Println(d, a)
		fmt.Println(e, b)
		fmt.Println(f, c)
	} else if x == 90 {
		fmt.Println(f, e, d)
		fmt.Println(c, b, a)
	} else if x == 90 {
		fmt.Println(f, e)
		fmt.Println(b, e)
		fmt.Println(a, d)

	}
}
