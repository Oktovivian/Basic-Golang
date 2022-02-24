package main

import "fmt"

func diskon(total_belanja float64) float64 {
	var besar_diskon float64
	besar_diskon = 0
	if total_belanja > 250000 {
		besar_diskon = 0.15 * float64(total_belanja)
	}
	return besar_diskon

}

func main() {
	var total_belanja float64
	fmt.Scan(&total_belanja)

	fmt.Println(diskon(total_belanja))
}
