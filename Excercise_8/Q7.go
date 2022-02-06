package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var salary float64

	for {
		fmt.Println("Enter the representative's sales in dollars (-1 to terminate):")
		fmt.Scan(&input)
		if input == "-1" {
			salary = -1
			break
		} else {
			if float, error := strconv.ParseFloat(input, 64); error == nil {
				salary = float
				weeklySalary := fmt.Sprintf("%.2f", calculate(salary))
				fmt.Println("The representative's salary is $", weeklySalary)
				continue
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}
}

func calculate(salary float64) float64 {
	return (salary * 0.09) + 200
}
