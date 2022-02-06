package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const pi float64 = 3.14159

func main() {
	var input string
	var radius float64

	for {
		fmt.Println("Enter the radius of a circle : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if float, error := strconv.ParseFloat(input, 64); error == nil {
				if float > 0 {
					radius = float
					break
				}
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	fmt.Println("The diameter of a circle with a radius of", radius, "is:", getDiameter(radius))
	fmt.Println("The circomference of a circle with a radius of", radius, "is:", getCircumference(radius))
	fmt.Println("The Area of a circle with a radius of", radius, "is:", getArea(radius))

}

func getDiameter(radius float64) float64 {
	return radius * 2
}

func getCircumference(radius float64) float64 {
	return getDiameter(radius) * pi
}

func getArea(radius float64) float64 {
	return math.Pow(radius, 2) * pi
}
