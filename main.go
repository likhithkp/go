package main

import "fmt"

type Car struct {
	Name  string
	Fuel  string
	Model int
}

func createCar(car Car) Car {
	car.Name = car.Name
	car.Fuel = car.Fuel
	car.Model = car.Model

	return car
}

func main() {
	newCar := createCar(Car{
		Name:  "BMW",
		Fuel:  "Diesel",
		Model: 2022,
	})

	fmt.Println(newCar)
}
