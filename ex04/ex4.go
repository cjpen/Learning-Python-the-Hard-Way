package main

import (
	"fmt"
)

func main () {
	var (
		cars = 100
		space_in_a_car = 4.0
		drivers = 30
		passengers = 90
		cars_not_driven = cars - drivers
		cars_driven = drivers
		carpool_capacisty = cars_driven * space_in_a_car
		average_passengers_per_car = passengers / cars_driven
	)

	fmt.Println("there are", cars, "cars available.")
	fmt.Println("There are only", drivers, "drivers available.")
	fmt.Println("There will be", cars_not_driven, "empty cars today")
	fmt.Println("We can transport", carpool_capacisty, "people today")
	fmt.Println("We have", passengers, "to carpool today.")
	fmt.Println("We need to put about", average_passengers_per_car, "in each car.")

}