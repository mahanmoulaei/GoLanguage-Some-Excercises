package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var litters, km, totalLitters, totalKm float64

	for {

		for {
			fmt.Println("Enter the number of litres used: (-1 to terminate)")
			fmt.Scan(&input)
			if input == "-1" {
				litters = -1
				break
			} else {
				if float, error := strconv.ParseFloat(input, 64); error == nil {
					if float > 0 {
						litters = float
						break
					}
				}
				fmt.Println("Error In The Entered Number! Try Again...")
			}
		}

		if litters == -1 {
			break
		}

		totalLitters += litters

		for {
			fmt.Println("Enter the number of kilometers traveled:")
			fmt.Scan(&input)
			if float, error := strconv.ParseFloat(input, 64); error == nil {
				if float > 0 {
					km = float
					break
				}
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}

		totalKm += km

		rate := fmt.Sprintf("%.6f", rateOfConsumption(litters, km))
		fmt.Println("The rate of gas consumption in litres per 100 kilometers for this gas refill is", rate)
	}

	totalRate := fmt.Sprintf("%.6f", rateOfConsumption(totalLitters, totalKm))
	fmt.Println("The total rate of gas consumption in litres per 100 kilometers is", totalRate)

}

func rateOfConsumption(litters float64, km float64) float64 {
	return (litters / km) * 100
}
