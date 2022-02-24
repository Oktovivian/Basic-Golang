package main

import "fmt"

func main() {
	var baterai int
	fmt.Scan(&baterai)
	if baterai == 100 {
		fmt.Print("Cabut charger")

	} else if baterai >= 75 && baterai <= 99 {
		fmt.Print("Charge dalam 30 menit")

	} else if baterai >= 20 && baterai <= 74 {
		fmt.Print("Charge dalam 10 menit")

	} else {
		fmt.Print("Segera charge")
	}
}
