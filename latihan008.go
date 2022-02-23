package main

import "fmt"

func int_division(a int, b int) int {
	return a / b
}

func division(a int, b int) float64 {
	return float64(a) / float64(b)
}

func main() {
	var n1, n2 int
	n1 = 10
	n2 = 13
	fmt.Println(int_division(n2, n1))
	fmt.Println(division(n1, n2))
}
