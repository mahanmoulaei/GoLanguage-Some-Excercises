package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var input string
	var deposit float64
	var years int
	var interestRate float64

	for {
		fmt.Println("Enter your deposit amount : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if float, error := strconv.ParseFloat(input, 64); error == nil {
				deposit = float
				break
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	for {
		fmt.Println("Enter the number of years : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				years = integer
				break
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	for {
		fmt.Println("Enter the interest rate : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if float, error := strconv.ParseFloat(input, 64); error == nil {
				interestRate = float
				break
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	fmt.Println("a) Your savings now:", deposit)
	fmt.Println("b) Your savings in", years, "years after a %", interestRate, "interest rate would be $", fmt.Sprintf("%.2f", savings(deposit, years, interestRate)))

}

func savings(deposit float64, years int, interestRate float64) float64 {
	var investment float64
	interestRate = interestRate / 100
	investment = deposit * math.Pow((1+(interestRate)), float64(years))
	return investment
}
