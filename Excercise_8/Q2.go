package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input string
	var num1, num2 int

	for {
		fmt.Println("Enter the first number : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				num1 = integer
				break
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	for {
		fmt.Println("Enter the second number : (type exit to quit the program)")
		fmt.Scan(&input)
		if input == "exit" {
			os.Exit(0)
		} else {
			if integer, error := strconv.Atoi(input); error == nil {
				num2 = integer
				break
			}
			fmt.Println("Error In The Entered Number! Try Again...")
		}
	}

	fmt.Println("The Sum of", num1, "and", num2, "is equals to", Sum(num1, num2))

	fmt.Println("The Product of", num1, "and", num2, "is equals to", Product(num1, num2))

	fmt.Println("The Difference between", num1, "and", num2, "is equals to", Difference(num1, num2))

	fmt.Println("The Quotient of", num1, "by", num2, "is equals to", Quotient(num1, num2))
}

func Sum(num1 int, num2 int) int {
	return num1 + num2
}

func Difference(num1 int, num2 int) int {
	return num1 - num2
}

func Product(num1 int, num2 int) int {
	return num1 * num2
}

func Quotient(num1 int, num2 int) int {
	return num1 / num2
}
