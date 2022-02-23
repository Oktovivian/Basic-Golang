package main

import "fmt"

func main() {
	var x, i int
	// for loop
	for x = 0; x < 10; x++ {
		fmt.Println(x)
	}

	// while loop
	i = 100
	for i <= 100 {
		fmt.Println(i)
		i++
		// atau pake i+=1
	}

}
