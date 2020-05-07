package main

import (
	"fmt"
)

func main () {
	
	// Variables can be written with := to denote value of variable
	cars := 100
	
	// Additionally, variables can be written and defined as below
	var (
		spaceInACar int = 4.0
		drivers int = 40
		passengers int = 90
		carsNotDriven int = cars - drivers
		carsDriven int = drivers
		carpoolCapacisty = carsDriven * spaceInACar
		averagePassengersPerCar = float64(passengers) / float64(carsDriven)
	)

	fmt.Println("there are", cars, "cars available.")
	fmt.Println("There are only", drivers, "drivers available.")
	fmt.Println("There will be", carsNotDriven, "empty cars today")
	fmt.Println("We can transport", carpoolCapacisty, "people today")
	fmt.Println("We have", passengers, "to carpool today.")
	fmt.Println("We need to put about", averagePassengersPerCar, "in each car.")

}