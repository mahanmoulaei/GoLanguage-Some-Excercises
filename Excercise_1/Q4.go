package main

import (
	"fmt"
	"math"
)

func main() {

	var inputFahrenheit int
	fmt.Print("Please enter a temperature in degrees fahrenheit : ")
	fmt.Scan(&inputFahrenheit)

	convertTemp(inputFahrenheit)
}

func convertTemp(fahrenheit int) {

	math.Round(float64(fahrenheit))

	var celcius int = (fahrenheit - 32) * 5 / 9

	fmt.Printf("%d degrees fahrenheit is %d degrees Celcius", fahrenheit, celcius)
}
