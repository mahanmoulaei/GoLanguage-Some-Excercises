package main

import (
	"fmt"
)

func main() {

	var num1 float64
	var num2 float64
	var num3 float64

	fmt.Print("Enter 3 numbers to determine the greatest\n")
	fmt.Print("Enter number 1 : ")
	fmt.Scan(&num1)
	fmt.Print("Enter number 2 : ")
	fmt.Scan(&num2)
	fmt.Print("Enter number 3 : ")
	fmt.Scan(&num3)

	if num1 >= num2 && num1 >= num3 {
		fmt.Print(num1, " is the greatest number")
	} else if num2 >= num1 && num2 >= num3 {
		fmt.Print(num2, " is the greatest number")
	} else if num3 >= num1 && num3 >= num2 {
		fmt.Print(num3, " is the greatest number")
	}

}
