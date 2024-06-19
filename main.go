package main

import "fmt"

type car struct {
	brand string
	fuel  string
	year  int
	model string
}

func main() {
	newCar := car{
		brand: "BMW",
		fuel:  "Diesel",
		year:  2023,
		model: "X5",
	}

	fmt.Println(newCar.brand, newCar.fuel)
}
