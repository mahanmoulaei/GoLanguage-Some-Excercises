package main

import (
	"fmt"
)

func main() {

	var num1 int
	var num2 int
	var num3 int

	fmt.Print("Enter 3 numbers to check if one of the numbers is equal to the sum of the other two\n")
	fmt.Print("Enter number 1 : ")
	fmt.Scan(&num1)
	fmt.Print("Enter number 2 : ")
	fmt.Scan(&num2)
	fmt.Print("Enter number 3 : ")
	fmt.Scan(&num3)

	if num1 == num2+num3 {

		fmt.Print(num1, " is equal to ", num2, " + ", num3)
	} else if num2 == num1+num3 {
		fmt.Print(num2, " is equal to ", num1, " + ", num3)
	} else if num3 == num1+num2 {
		fmt.Print(num3, " is equal to ", num1, " + ", num2)
	} else {
		fmt.Print("No solution")
	}
}
