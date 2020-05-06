package main

import (
	"fmt"
)

func main()  {
	// Displays text string
	fmt.Println("I will now count my chickens:")

	// Displays text string, then performs calculations. The order of operations is observed.
	// Divides 30 by 6 to equal 5 then adds 25
	fmt.Println("Hens", 25 + 30 / 6)
	// Multiplies 25 by 3, then the remainer of 75 divided by 4 which equals 2 (the quotient is dismissed and the value shown is the remainder). This number is then subtracted from 100 to equal 98.
	fmt.Println("Roosters", 100 - 25 * 3 % 4)

	// Displays text string
	fmt.Println("Now I will count the eggs:")

	// Displays computational output. See ex3.py for description
	fmt.Println(3 + 2 + 1 - 5 + 4 % 2 - 1 / 4 + 6)

	// Displays text string
	fmt.Println("Is it true that 3 + 2 < 5 - 7?")

	// Performs the computational output
	fmt.Println(3 + 2 < 5 - 7)

	// Displays text string, then the computational ouptut
	fmt.Println("What is 3 + 2?", 3 + 2)
	fmt.Println("What is 5 - 7?", 5 - 7)

	// Displays text string
	fmt.Println("Oh, that's why it's False")

	// Displays text string
	fmt.Println("How about some more.")

	// Displays text string, then evaluates the expression as True or False
	fmt.Println("Is it greater?", 5 > -2)
	fmt.Println("Is it greater or equal?", 5 >= -2)
	fmt.Println("Is it less or equal?", 5 <= -2)
}