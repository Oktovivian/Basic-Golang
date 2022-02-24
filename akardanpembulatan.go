package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	var sisi_miring float64
	var sisi_miring_final float64
	fmt.Scan(&a)
	fmt.Scan(&b)

	sisi_miring = math.Sqrt(a*a + b*b)
	sisi_miring_final = math.Round(sisi_miring)
	fmt.Printf("sisi miring: %.1f", sisi_miring_final)
}
