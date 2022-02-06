package main

import "fmt"

func main() {
	forLoop()
	whileLoop()
}

func forLoop() {
	min := -40
	max := 40
	fmt.Println("\n*** Using For Loop ***")
	for i := min; i <= max; i += 5 {
		fmt.Printf("%v F = %v C \n", fahrenheit(i), i)
	}
}

func whileLoop() {
	i := -40
	var max = 40
	fmt.Println("\n*** Using While Loop ***")
	for i <= max {
		fmt.Printf("%v F = %v C \n", fahrenheit(i), i)
		i += 5
	}
}

func fahrenheit(tempC int) float64 {
	var tempF float64
	tempF = (9 * float64(tempC) / 5) + 32
	return tempF
}
